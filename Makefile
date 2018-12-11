build:
	docker build -t docker.ci.pix4d.com/concourse-pool-test .
	docker push docker.ci.pix4d.com/concourse-pool-test

pipeline:
	fly -t prod-developers dp -p test-resource-pool --non-interactive
	fly -t prod-developers sp -c ci/pix4d_demo_pipeline.yml -p test-resource-pool --non-interactive
	fly -t prod-developers unpause-pipeline -p test-resource-pool

run:
	fly -t prod-developers tj -w -j test-resource-pool/dummy
