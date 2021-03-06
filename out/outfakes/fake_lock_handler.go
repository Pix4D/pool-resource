// Code generated by counterfeiter. DO NOT EDIT.
package outfakes

import (
	sync "sync"

	out "github.com/concourse/pool-resource/out"
)

type FakeLockHandler struct {
	AddLockStub        func(string, []byte, bool) (string, error)
	addLockMutex       sync.RWMutex
	addLockArgsForCall []struct {
		arg1 string
		arg2 []byte
		arg3 bool
	}
	addLockReturns struct {
		result1 string
		result2 error
	}
	addLockReturnsOnCall map[int]struct {
		result1 string
		result2 error
	}
	BroadcastLockPoolStub        func() ([]byte, error)
	broadcastLockPoolMutex       sync.RWMutex
	broadcastLockPoolArgsForCall []struct {
	}
	broadcastLockPoolReturns struct {
		result1 []byte
		result2 error
	}
	broadcastLockPoolReturnsOnCall map[int]struct {
		result1 []byte
		result2 error
	}
	ClaimLockStub        func(string) (string, error)
	claimLockMutex       sync.RWMutex
	claimLockArgsForCall []struct {
		arg1 string
	}
	claimLockReturns struct {
		result1 string
		result2 error
	}
	claimLockReturnsOnCall map[int]struct {
		result1 string
		result2 error
	}
	GetStatsStub        func() (*out.Stats, error)
	getStatsMutex       sync.RWMutex
	getStatsArgsForCall []struct {
	}
	getStatsReturns struct {
		result1 *out.Stats
		result2 error
	}
	getStatsReturnsOnCall map[int]struct {
		result1 *out.Stats
		result2 error
	}
	GrabAvailableLockStub        func() (string, string, error)
	grabAvailableLockMutex       sync.RWMutex
	grabAvailableLockArgsForCall []struct {
	}
	grabAvailableLockReturns struct {
		result1 string
		result2 string
		result3 error
	}
	grabAvailableLockReturnsOnCall map[int]struct {
		result1 string
		result2 string
		result3 error
	}
	RemoveLockStub        func(string) (string, error)
	removeLockMutex       sync.RWMutex
	removeLockArgsForCall []struct {
		arg1 string
	}
	removeLockReturns struct {
		result1 string
		result2 error
	}
	removeLockReturnsOnCall map[int]struct {
		result1 string
		result2 error
	}
	ResetLockStub        func() error
	resetLockMutex       sync.RWMutex
	resetLockArgsForCall []struct {
	}
	resetLockReturns struct {
		result1 error
	}
	resetLockReturnsOnCall map[int]struct {
		result1 error
	}
	SetupStub        func() error
	setupMutex       sync.RWMutex
	setupArgsForCall []struct {
	}
	setupReturns struct {
		result1 error
	}
	setupReturnsOnCall map[int]struct {
		result1 error
	}
	UnclaimLockStub        func(string) (string, error)
	unclaimLockMutex       sync.RWMutex
	unclaimLockArgsForCall []struct {
		arg1 string
	}
	unclaimLockReturns struct {
		result1 string
		result2 error
	}
	unclaimLockReturnsOnCall map[int]struct {
		result1 string
		result2 error
	}
	UpdateLockStub        func(string, []byte) (string, error)
	updateLockMutex       sync.RWMutex
	updateLockArgsForCall []struct {
		arg1 string
		arg2 []byte
	}
	updateLockReturns struct {
		result1 string
		result2 error
	}
	updateLockReturnsOnCall map[int]struct {
		result1 string
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeLockHandler) AddLock(arg1 string, arg2 []byte, arg3 bool) (string, error) {
	var arg2Copy []byte
	if arg2 != nil {
		arg2Copy = make([]byte, len(arg2))
		copy(arg2Copy, arg2)
	}
	fake.addLockMutex.Lock()
	ret, specificReturn := fake.addLockReturnsOnCall[len(fake.addLockArgsForCall)]
	fake.addLockArgsForCall = append(fake.addLockArgsForCall, struct {
		arg1 string
		arg2 []byte
		arg3 bool
	}{arg1, arg2Copy, arg3})
	fake.recordInvocation("AddLock", []interface{}{arg1, arg2Copy, arg3})
	fake.addLockMutex.Unlock()
	if fake.AddLockStub != nil {
		return fake.AddLockStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.addLockReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeLockHandler) AddLockCallCount() int {
	fake.addLockMutex.RLock()
	defer fake.addLockMutex.RUnlock()
	return len(fake.addLockArgsForCall)
}

func (fake *FakeLockHandler) AddLockCalls(stub func(string, []byte, bool) (string, error)) {
	fake.addLockMutex.Lock()
	defer fake.addLockMutex.Unlock()
	fake.AddLockStub = stub
}

func (fake *FakeLockHandler) AddLockArgsForCall(i int) (string, []byte, bool) {
	fake.addLockMutex.RLock()
	defer fake.addLockMutex.RUnlock()
	argsForCall := fake.addLockArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeLockHandler) AddLockReturns(result1 string, result2 error) {
	fake.addLockMutex.Lock()
	defer fake.addLockMutex.Unlock()
	fake.AddLockStub = nil
	fake.addLockReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeLockHandler) AddLockReturnsOnCall(i int, result1 string, result2 error) {
	fake.addLockMutex.Lock()
	defer fake.addLockMutex.Unlock()
	fake.AddLockStub = nil
	if fake.addLockReturnsOnCall == nil {
		fake.addLockReturnsOnCall = make(map[int]struct {
			result1 string
			result2 error
		})
	}
	fake.addLockReturnsOnCall[i] = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeLockHandler) BroadcastLockPool() ([]byte, error) {
	fake.broadcastLockPoolMutex.Lock()
	ret, specificReturn := fake.broadcastLockPoolReturnsOnCall[len(fake.broadcastLockPoolArgsForCall)]
	fake.broadcastLockPoolArgsForCall = append(fake.broadcastLockPoolArgsForCall, struct {
	}{})
	fake.recordInvocation("BroadcastLockPool", []interface{}{})
	fake.broadcastLockPoolMutex.Unlock()
	if fake.BroadcastLockPoolStub != nil {
		return fake.BroadcastLockPoolStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.broadcastLockPoolReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeLockHandler) BroadcastLockPoolCallCount() int {
	fake.broadcastLockPoolMutex.RLock()
	defer fake.broadcastLockPoolMutex.RUnlock()
	return len(fake.broadcastLockPoolArgsForCall)
}

func (fake *FakeLockHandler) BroadcastLockPoolCalls(stub func() ([]byte, error)) {
	fake.broadcastLockPoolMutex.Lock()
	defer fake.broadcastLockPoolMutex.Unlock()
	fake.BroadcastLockPoolStub = stub
}

func (fake *FakeLockHandler) BroadcastLockPoolReturns(result1 []byte, result2 error) {
	fake.broadcastLockPoolMutex.Lock()
	defer fake.broadcastLockPoolMutex.Unlock()
	fake.BroadcastLockPoolStub = nil
	fake.broadcastLockPoolReturns = struct {
		result1 []byte
		result2 error
	}{result1, result2}
}

func (fake *FakeLockHandler) BroadcastLockPoolReturnsOnCall(i int, result1 []byte, result2 error) {
	fake.broadcastLockPoolMutex.Lock()
	defer fake.broadcastLockPoolMutex.Unlock()
	fake.BroadcastLockPoolStub = nil
	if fake.broadcastLockPoolReturnsOnCall == nil {
		fake.broadcastLockPoolReturnsOnCall = make(map[int]struct {
			result1 []byte
			result2 error
		})
	}
	fake.broadcastLockPoolReturnsOnCall[i] = struct {
		result1 []byte
		result2 error
	}{result1, result2}
}

func (fake *FakeLockHandler) ClaimLock(arg1 string) (string, error) {
	fake.claimLockMutex.Lock()
	ret, specificReturn := fake.claimLockReturnsOnCall[len(fake.claimLockArgsForCall)]
	fake.claimLockArgsForCall = append(fake.claimLockArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("ClaimLock", []interface{}{arg1})
	fake.claimLockMutex.Unlock()
	if fake.ClaimLockStub != nil {
		return fake.ClaimLockStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.claimLockReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeLockHandler) ClaimLockCallCount() int {
	fake.claimLockMutex.RLock()
	defer fake.claimLockMutex.RUnlock()
	return len(fake.claimLockArgsForCall)
}

func (fake *FakeLockHandler) ClaimLockCalls(stub func(string) (string, error)) {
	fake.claimLockMutex.Lock()
	defer fake.claimLockMutex.Unlock()
	fake.ClaimLockStub = stub
}

func (fake *FakeLockHandler) ClaimLockArgsForCall(i int) string {
	fake.claimLockMutex.RLock()
	defer fake.claimLockMutex.RUnlock()
	argsForCall := fake.claimLockArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeLockHandler) ClaimLockReturns(result1 string, result2 error) {
	fake.claimLockMutex.Lock()
	defer fake.claimLockMutex.Unlock()
	fake.ClaimLockStub = nil
	fake.claimLockReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeLockHandler) ClaimLockReturnsOnCall(i int, result1 string, result2 error) {
	fake.claimLockMutex.Lock()
	defer fake.claimLockMutex.Unlock()
	fake.ClaimLockStub = nil
	if fake.claimLockReturnsOnCall == nil {
		fake.claimLockReturnsOnCall = make(map[int]struct {
			result1 string
			result2 error
		})
	}
	fake.claimLockReturnsOnCall[i] = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeLockHandler) GetStats() (*out.Stats, error) {
	fake.getStatsMutex.Lock()
	ret, specificReturn := fake.getStatsReturnsOnCall[len(fake.getStatsArgsForCall)]
	fake.getStatsArgsForCall = append(fake.getStatsArgsForCall, struct {
	}{})
	fake.recordInvocation("GetStats", []interface{}{})
	fake.getStatsMutex.Unlock()
	if fake.GetStatsStub != nil {
		return fake.GetStatsStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.getStatsReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeLockHandler) GetStatsCallCount() int {
	fake.getStatsMutex.RLock()
	defer fake.getStatsMutex.RUnlock()
	return len(fake.getStatsArgsForCall)
}

func (fake *FakeLockHandler) GetStatsCalls(stub func() (*out.Stats, error)) {
	fake.getStatsMutex.Lock()
	defer fake.getStatsMutex.Unlock()
	fake.GetStatsStub = stub
}

func (fake *FakeLockHandler) GetStatsReturns(result1 *out.Stats, result2 error) {
	fake.getStatsMutex.Lock()
	defer fake.getStatsMutex.Unlock()
	fake.GetStatsStub = nil
	fake.getStatsReturns = struct {
		result1 *out.Stats
		result2 error
	}{result1, result2}
}

func (fake *FakeLockHandler) GetStatsReturnsOnCall(i int, result1 *out.Stats, result2 error) {
	fake.getStatsMutex.Lock()
	defer fake.getStatsMutex.Unlock()
	fake.GetStatsStub = nil
	if fake.getStatsReturnsOnCall == nil {
		fake.getStatsReturnsOnCall = make(map[int]struct {
			result1 *out.Stats
			result2 error
		})
	}
	fake.getStatsReturnsOnCall[i] = struct {
		result1 *out.Stats
		result2 error
	}{result1, result2}
}

func (fake *FakeLockHandler) GrabAvailableLock() (string, string, error) {
	fake.grabAvailableLockMutex.Lock()
	ret, specificReturn := fake.grabAvailableLockReturnsOnCall[len(fake.grabAvailableLockArgsForCall)]
	fake.grabAvailableLockArgsForCall = append(fake.grabAvailableLockArgsForCall, struct {
	}{})
	fake.recordInvocation("GrabAvailableLock", []interface{}{})
	fake.grabAvailableLockMutex.Unlock()
	if fake.GrabAvailableLockStub != nil {
		return fake.GrabAvailableLockStub()
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	fakeReturns := fake.grabAvailableLockReturns
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *FakeLockHandler) GrabAvailableLockCallCount() int {
	fake.grabAvailableLockMutex.RLock()
	defer fake.grabAvailableLockMutex.RUnlock()
	return len(fake.grabAvailableLockArgsForCall)
}

func (fake *FakeLockHandler) GrabAvailableLockCalls(stub func() (string, string, error)) {
	fake.grabAvailableLockMutex.Lock()
	defer fake.grabAvailableLockMutex.Unlock()
	fake.GrabAvailableLockStub = stub
}

func (fake *FakeLockHandler) GrabAvailableLockReturns(result1 string, result2 string, result3 error) {
	fake.grabAvailableLockMutex.Lock()
	defer fake.grabAvailableLockMutex.Unlock()
	fake.GrabAvailableLockStub = nil
	fake.grabAvailableLockReturns = struct {
		result1 string
		result2 string
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeLockHandler) GrabAvailableLockReturnsOnCall(i int, result1 string, result2 string, result3 error) {
	fake.grabAvailableLockMutex.Lock()
	defer fake.grabAvailableLockMutex.Unlock()
	fake.GrabAvailableLockStub = nil
	if fake.grabAvailableLockReturnsOnCall == nil {
		fake.grabAvailableLockReturnsOnCall = make(map[int]struct {
			result1 string
			result2 string
			result3 error
		})
	}
	fake.grabAvailableLockReturnsOnCall[i] = struct {
		result1 string
		result2 string
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeLockHandler) RemoveLock(arg1 string) (string, error) {
	fake.removeLockMutex.Lock()
	ret, specificReturn := fake.removeLockReturnsOnCall[len(fake.removeLockArgsForCall)]
	fake.removeLockArgsForCall = append(fake.removeLockArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("RemoveLock", []interface{}{arg1})
	fake.removeLockMutex.Unlock()
	if fake.RemoveLockStub != nil {
		return fake.RemoveLockStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.removeLockReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeLockHandler) RemoveLockCallCount() int {
	fake.removeLockMutex.RLock()
	defer fake.removeLockMutex.RUnlock()
	return len(fake.removeLockArgsForCall)
}

func (fake *FakeLockHandler) RemoveLockCalls(stub func(string) (string, error)) {
	fake.removeLockMutex.Lock()
	defer fake.removeLockMutex.Unlock()
	fake.RemoveLockStub = stub
}

func (fake *FakeLockHandler) RemoveLockArgsForCall(i int) string {
	fake.removeLockMutex.RLock()
	defer fake.removeLockMutex.RUnlock()
	argsForCall := fake.removeLockArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeLockHandler) RemoveLockReturns(result1 string, result2 error) {
	fake.removeLockMutex.Lock()
	defer fake.removeLockMutex.Unlock()
	fake.RemoveLockStub = nil
	fake.removeLockReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeLockHandler) RemoveLockReturnsOnCall(i int, result1 string, result2 error) {
	fake.removeLockMutex.Lock()
	defer fake.removeLockMutex.Unlock()
	fake.RemoveLockStub = nil
	if fake.removeLockReturnsOnCall == nil {
		fake.removeLockReturnsOnCall = make(map[int]struct {
			result1 string
			result2 error
		})
	}
	fake.removeLockReturnsOnCall[i] = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeLockHandler) ResetLock() error {
	fake.resetLockMutex.Lock()
	ret, specificReturn := fake.resetLockReturnsOnCall[len(fake.resetLockArgsForCall)]
	fake.resetLockArgsForCall = append(fake.resetLockArgsForCall, struct {
	}{})
	fake.recordInvocation("ResetLock", []interface{}{})
	fake.resetLockMutex.Unlock()
	if fake.ResetLockStub != nil {
		return fake.ResetLockStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.resetLockReturns
	return fakeReturns.result1
}

func (fake *FakeLockHandler) ResetLockCallCount() int {
	fake.resetLockMutex.RLock()
	defer fake.resetLockMutex.RUnlock()
	return len(fake.resetLockArgsForCall)
}

func (fake *FakeLockHandler) ResetLockCalls(stub func() error) {
	fake.resetLockMutex.Lock()
	defer fake.resetLockMutex.Unlock()
	fake.ResetLockStub = stub
}

func (fake *FakeLockHandler) ResetLockReturns(result1 error) {
	fake.resetLockMutex.Lock()
	defer fake.resetLockMutex.Unlock()
	fake.ResetLockStub = nil
	fake.resetLockReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeLockHandler) ResetLockReturnsOnCall(i int, result1 error) {
	fake.resetLockMutex.Lock()
	defer fake.resetLockMutex.Unlock()
	fake.ResetLockStub = nil
	if fake.resetLockReturnsOnCall == nil {
		fake.resetLockReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.resetLockReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeLockHandler) Setup() error {
	fake.setupMutex.Lock()
	ret, specificReturn := fake.setupReturnsOnCall[len(fake.setupArgsForCall)]
	fake.setupArgsForCall = append(fake.setupArgsForCall, struct {
	}{})
	fake.recordInvocation("Setup", []interface{}{})
	fake.setupMutex.Unlock()
	if fake.SetupStub != nil {
		return fake.SetupStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.setupReturns
	return fakeReturns.result1
}

func (fake *FakeLockHandler) SetupCallCount() int {
	fake.setupMutex.RLock()
	defer fake.setupMutex.RUnlock()
	return len(fake.setupArgsForCall)
}

func (fake *FakeLockHandler) SetupCalls(stub func() error) {
	fake.setupMutex.Lock()
	defer fake.setupMutex.Unlock()
	fake.SetupStub = stub
}

func (fake *FakeLockHandler) SetupReturns(result1 error) {
	fake.setupMutex.Lock()
	defer fake.setupMutex.Unlock()
	fake.SetupStub = nil
	fake.setupReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeLockHandler) SetupReturnsOnCall(i int, result1 error) {
	fake.setupMutex.Lock()
	defer fake.setupMutex.Unlock()
	fake.SetupStub = nil
	if fake.setupReturnsOnCall == nil {
		fake.setupReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.setupReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeLockHandler) UnclaimLock(arg1 string) (string, error) {
	fake.unclaimLockMutex.Lock()
	ret, specificReturn := fake.unclaimLockReturnsOnCall[len(fake.unclaimLockArgsForCall)]
	fake.unclaimLockArgsForCall = append(fake.unclaimLockArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("UnclaimLock", []interface{}{arg1})
	fake.unclaimLockMutex.Unlock()
	if fake.UnclaimLockStub != nil {
		return fake.UnclaimLockStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.unclaimLockReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeLockHandler) UnclaimLockCallCount() int {
	fake.unclaimLockMutex.RLock()
	defer fake.unclaimLockMutex.RUnlock()
	return len(fake.unclaimLockArgsForCall)
}

func (fake *FakeLockHandler) UnclaimLockCalls(stub func(string) (string, error)) {
	fake.unclaimLockMutex.Lock()
	defer fake.unclaimLockMutex.Unlock()
	fake.UnclaimLockStub = stub
}

func (fake *FakeLockHandler) UnclaimLockArgsForCall(i int) string {
	fake.unclaimLockMutex.RLock()
	defer fake.unclaimLockMutex.RUnlock()
	argsForCall := fake.unclaimLockArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeLockHandler) UnclaimLockReturns(result1 string, result2 error) {
	fake.unclaimLockMutex.Lock()
	defer fake.unclaimLockMutex.Unlock()
	fake.UnclaimLockStub = nil
	fake.unclaimLockReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeLockHandler) UnclaimLockReturnsOnCall(i int, result1 string, result2 error) {
	fake.unclaimLockMutex.Lock()
	defer fake.unclaimLockMutex.Unlock()
	fake.UnclaimLockStub = nil
	if fake.unclaimLockReturnsOnCall == nil {
		fake.unclaimLockReturnsOnCall = make(map[int]struct {
			result1 string
			result2 error
		})
	}
	fake.unclaimLockReturnsOnCall[i] = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeLockHandler) UpdateLock(arg1 string, arg2 []byte) (string, error) {
	var arg2Copy []byte
	if arg2 != nil {
		arg2Copy = make([]byte, len(arg2))
		copy(arg2Copy, arg2)
	}
	fake.updateLockMutex.Lock()
	ret, specificReturn := fake.updateLockReturnsOnCall[len(fake.updateLockArgsForCall)]
	fake.updateLockArgsForCall = append(fake.updateLockArgsForCall, struct {
		arg1 string
		arg2 []byte
	}{arg1, arg2Copy})
	fake.recordInvocation("UpdateLock", []interface{}{arg1, arg2Copy})
	fake.updateLockMutex.Unlock()
	if fake.UpdateLockStub != nil {
		return fake.UpdateLockStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.updateLockReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeLockHandler) UpdateLockCallCount() int {
	fake.updateLockMutex.RLock()
	defer fake.updateLockMutex.RUnlock()
	return len(fake.updateLockArgsForCall)
}

func (fake *FakeLockHandler) UpdateLockCalls(stub func(string, []byte) (string, error)) {
	fake.updateLockMutex.Lock()
	defer fake.updateLockMutex.Unlock()
	fake.UpdateLockStub = stub
}

func (fake *FakeLockHandler) UpdateLockArgsForCall(i int) (string, []byte) {
	fake.updateLockMutex.RLock()
	defer fake.updateLockMutex.RUnlock()
	argsForCall := fake.updateLockArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeLockHandler) UpdateLockReturns(result1 string, result2 error) {
	fake.updateLockMutex.Lock()
	defer fake.updateLockMutex.Unlock()
	fake.UpdateLockStub = nil
	fake.updateLockReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeLockHandler) UpdateLockReturnsOnCall(i int, result1 string, result2 error) {
	fake.updateLockMutex.Lock()
	defer fake.updateLockMutex.Unlock()
	fake.UpdateLockStub = nil
	if fake.updateLockReturnsOnCall == nil {
		fake.updateLockReturnsOnCall = make(map[int]struct {
			result1 string
			result2 error
		})
	}
	fake.updateLockReturnsOnCall[i] = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeLockHandler) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.addLockMutex.RLock()
	defer fake.addLockMutex.RUnlock()
	fake.broadcastLockPoolMutex.RLock()
	defer fake.broadcastLockPoolMutex.RUnlock()
	fake.claimLockMutex.RLock()
	defer fake.claimLockMutex.RUnlock()
	fake.getStatsMutex.RLock()
	defer fake.getStatsMutex.RUnlock()
	fake.grabAvailableLockMutex.RLock()
	defer fake.grabAvailableLockMutex.RUnlock()
	fake.removeLockMutex.RLock()
	defer fake.removeLockMutex.RUnlock()
	fake.resetLockMutex.RLock()
	defer fake.resetLockMutex.RUnlock()
	fake.setupMutex.RLock()
	defer fake.setupMutex.RUnlock()
	fake.unclaimLockMutex.RLock()
	defer fake.unclaimLockMutex.RUnlock()
	fake.updateLockMutex.RLock()
	defer fake.updateLockMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeLockHandler) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ out.LockHandler = new(FakeLockHandler)
