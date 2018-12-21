package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"time"

	"github.com/concourse/pool-resource/out"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/push"
)

func main() {

	rand.Seed(time.Now().UnixNano())

	if len(os.Args) < 2 {
		println("usage: " + os.Args[0] + " <source>")
		os.Exit(1)
	}

	sourceDir := os.Args[1]

	var request out.OutRequest
	err := json.NewDecoder(os.Stdin).Decode(&request)
	if err != nil {
		fatal("reading request", err)
	}
	defer os.Stdin.Close()

	errorMessages := request.Validate()
	if len(errorMessages) > 0 {
		for _, errorMessage := range errorMessages {
			println(errorMessage)
		}
		os.Exit(1)
	}

	if request.Source.RetryDelay == 0 {
		request.Source.RetryDelay = 10 * time.Second
	}

	var (
		executionTime     prometheus.Histogram
		executionTimer    *prometheus.Timer
		numAvailableLocks prometheus.Gauge
		numClaimedLocks   prometheus.Gauge

		stats *out.Stats
	)

	lockPool := out.NewLockPool(request.Source, os.Stderr)

	if request.Source.PrometheusPushGateway != "" {
		fmt.Fprintf(os.Stderr, "Setting up prometheus client, pushgateway=%s\n",
			request.Source.PrometheusPushGateway)
		executionTime = prometheus.NewHistogram(prometheus.HistogramOpts{
			Name: "resource_pool_execution_time",
			Help: "Execution time of the current operation",
		})
		executionTimer = prometheus.NewTimer(executionTime)
		numAvailableLocks = prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "resource_pool_num_available_locks",
			Help: "The number of unclaimed locks in the current pool",
		})
		numClaimedLocks = prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "resource_pool_num_claimed_locks",
			Help: "The number of claimed locks in the current pool",
		})
	} else {
		fmt.Fprintln(os.Stderr, "No prometheus pushgateway set. Will not emit any metrics.")
	}

	var (
		lock          string
		version       out.Version
		operationName string
	)

	// TODO bad design: should never support multiple operations in a single run
	// ... or alternatively, these should be passed as a list instead of a dict
	if request.Params.Acquire {
		operationName = "acquire_lock"
		lock, version, err = lockPool.AcquireLock()
		if err != nil {
			fatal("acquiring lock", err)
		}
	} else if request.Params.Release != "" {
		operationName = "release_lock"
		poolName := filepath.Join(sourceDir, request.Params.Release)
		lock, version, err = lockPool.ReleaseLock(poolName)
		if err != nil {
			fatal("releasing lock", err)
		}
	} else if request.Params.Add != "" {
		operationName = "add_unclaimed_lock"
		lockPath := filepath.Join(sourceDir, request.Params.Add)
		lock, version, err = lockPool.AddUnclaimedLock(lockPath)
		if err != nil {
			fatal("adding lock", err)
		}
	} else if request.Params.AddClaimed != "" {
		operationName = "add_claimed_lock"
		lockPath := filepath.Join(sourceDir, request.Params.AddClaimed)
		lock, version, err = lockPool.AddClaimedLock(lockPath)
		if err != nil {
			fatal("adding pre-claimed lock", err)
		}
	} else if request.Params.Remove != "" {
		operationName = "remove_lock"
		removePath := filepath.Join(sourceDir, request.Params.Remove)
		lock, version, err = lockPool.RemoveLock(removePath)
		if err != nil {
			fatal("removing lock", err)
		}
	} else if request.Params.Claim != "" {
		operationName = "claim_lock"
		lock = request.Params.Claim
		version, err = lockPool.ClaimLock(lock)
		if err != nil {
			fatal("claiming lock", err)
		}
	} else if request.Params.Update != "" {
		operationName = "update_lock"
		lockPath := filepath.Join(sourceDir, request.Params.Update)
		lock, version, err = lockPool.UpdateLock(lockPath)
		if err != nil {
			fatal("updating lock", err)
		}
	}

	err = json.NewEncoder(os.Stdout).Encode(out.OutResponse{
		Version: version,
		Metadata: []out.MetadataPair{
			{Name: "lock_name", Value: lock},
			{Name: "pool_name", Value: request.Source.Pool},
		},
	})

	if err != nil {
		fatal("encoding output", err)
	}

	if request.Source.PrometheusPushGateway != "" {
		executionTimer.ObserveDuration()

		err = lockPool.LockHandler.Setup()
		if err != nil {
			fatal("resetting up the git repository", err)
		}
		stats, err = lockPool.LockHandler.GetStats()
		if err != nil {
			fatal("getting lock handler stats", err)
		}
		numAvailableLocks.Set(float64(stats.Unclaimed))
		numClaimedLocks.Set(float64(stats.Claimed))

		fmt.Fprintln(os.Stderr, "Pushing prometheus metrics...")
		err = push.New(request.Source.PrometheusPushGateway, "concourse_resource_pool").
			Collector(executionTime).
			Grouping("pool", request.Source.Pool).
			Grouping("branch", request.Source.Branch). // git branch of the locks repo. Useful to filter out test runs
			Grouping("operation", operationName).
			Push()
		if err != nil {
			fatal("pushing metrics", err)
		}

		err = push.New(request.Source.PrometheusPushGateway, "concourse_resource_pool").
			Collector(numAvailableLocks).
			Collector(numClaimedLocks).
			Grouping("pool", request.Source.Pool).
			Grouping("branch", request.Source.Branch). // git branch of the locks repo. Useful to filter out test runs
			Push()
		if err != nil {
			fatal("pushing metrics", err)
		}
	}
}

func fatal(doing string, err error) {
	println("error " + doing + ": " + err.Error())
	os.Exit(1)
}
