// This file was generated by counterfeiter
package repfakes

import (
	"net/http"
	"sync"
	"time"

	"code.cloudfoundry.org/bbs/models"
	"code.cloudfoundry.org/lager"
	"code.cloudfoundry.org/rep"
)

type FakeClient struct {
	StateStub        func(logger lager.Logger) (rep.CellState, error)
	stateMutex       sync.RWMutex
	stateArgsForCall []struct {
		logger lager.Logger
	}
	stateReturns struct {
		result1 rep.CellState
		result2 error
	}
	stateReturnsOnCall map[int]struct {
		result1 rep.CellState
		result2 error
	}
	PerformStub        func(logger lager.Logger, work rep.Work) (rep.Work, error)
	performMutex       sync.RWMutex
	performArgsForCall []struct {
		logger lager.Logger
		work   rep.Work
	}
	performReturns struct {
		result1 rep.Work
		result2 error
	}
	performReturnsOnCall map[int]struct {
		result1 rep.Work
		result2 error
	}
	StopLRPInstanceStub        func(logger lager.Logger, key models.ActualLRPKey, instanceKey models.ActualLRPInstanceKey) error
	stopLRPInstanceMutex       sync.RWMutex
	stopLRPInstanceArgsForCall []struct {
		logger      lager.Logger
		key         models.ActualLRPKey
		instanceKey models.ActualLRPInstanceKey
	}
	stopLRPInstanceReturns struct {
		result1 error
	}
	stopLRPInstanceReturnsOnCall map[int]struct {
		result1 error
	}
	CancelTaskStub        func(logger lager.Logger, taskGuid string) error
	cancelTaskMutex       sync.RWMutex
	cancelTaskArgsForCall []struct {
		logger   lager.Logger
		taskGuid string
	}
	cancelTaskReturns struct {
		result1 error
	}
	cancelTaskReturnsOnCall map[int]struct {
		result1 error
	}
	SetStateClientStub        func(stateClient *http.Client)
	setStateClientMutex       sync.RWMutex
	setStateClientArgsForCall []struct {
		stateClient *http.Client
	}
	StateClientTimeoutStub        func() time.Duration
	stateClientTimeoutMutex       sync.RWMutex
	stateClientTimeoutArgsForCall []struct{}
	stateClientTimeoutReturns     struct {
		result1 time.Duration
	}
	stateClientTimeoutReturnsOnCall map[int]struct {
		result1 time.Duration
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeClient) State(logger lager.Logger) (rep.CellState, error) {
	fake.stateMutex.Lock()
	ret, specificReturn := fake.stateReturnsOnCall[len(fake.stateArgsForCall)]
	fake.stateArgsForCall = append(fake.stateArgsForCall, struct {
		logger lager.Logger
	}{logger})
	fake.recordInvocation("State", []interface{}{logger})
	fake.stateMutex.Unlock()
	if fake.StateStub != nil {
		return fake.StateStub(logger)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.stateReturns.result1, fake.stateReturns.result2
}

func (fake *FakeClient) StateCallCount() int {
	fake.stateMutex.RLock()
	defer fake.stateMutex.RUnlock()
	return len(fake.stateArgsForCall)
}

func (fake *FakeClient) StateArgsForCall(i int) lager.Logger {
	fake.stateMutex.RLock()
	defer fake.stateMutex.RUnlock()
	return fake.stateArgsForCall[i].logger
}

func (fake *FakeClient) StateReturns(result1 rep.CellState, result2 error) {
	fake.StateStub = nil
	fake.stateReturns = struct {
		result1 rep.CellState
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) StateReturnsOnCall(i int, result1 rep.CellState, result2 error) {
	fake.StateStub = nil
	if fake.stateReturnsOnCall == nil {
		fake.stateReturnsOnCall = make(map[int]struct {
			result1 rep.CellState
			result2 error
		})
	}
	fake.stateReturnsOnCall[i] = struct {
		result1 rep.CellState
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) Perform(logger lager.Logger, work rep.Work) (rep.Work, error) {
	fake.performMutex.Lock()
	ret, specificReturn := fake.performReturnsOnCall[len(fake.performArgsForCall)]
	fake.performArgsForCall = append(fake.performArgsForCall, struct {
		logger lager.Logger
		work   rep.Work
	}{logger, work})
	fake.recordInvocation("Perform", []interface{}{logger, work})
	fake.performMutex.Unlock()
	if fake.PerformStub != nil {
		return fake.PerformStub(logger, work)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.performReturns.result1, fake.performReturns.result2
}

func (fake *FakeClient) PerformCallCount() int {
	fake.performMutex.RLock()
	defer fake.performMutex.RUnlock()
	return len(fake.performArgsForCall)
}

func (fake *FakeClient) PerformArgsForCall(i int) (lager.Logger, rep.Work) {
	fake.performMutex.RLock()
	defer fake.performMutex.RUnlock()
	return fake.performArgsForCall[i].logger, fake.performArgsForCall[i].work
}

func (fake *FakeClient) PerformReturns(result1 rep.Work, result2 error) {
	fake.PerformStub = nil
	fake.performReturns = struct {
		result1 rep.Work
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) PerformReturnsOnCall(i int, result1 rep.Work, result2 error) {
	fake.PerformStub = nil
	if fake.performReturnsOnCall == nil {
		fake.performReturnsOnCall = make(map[int]struct {
			result1 rep.Work
			result2 error
		})
	}
	fake.performReturnsOnCall[i] = struct {
		result1 rep.Work
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) StopLRPInstance(logger lager.Logger, key models.ActualLRPKey, instanceKey models.ActualLRPInstanceKey) error {
	fake.stopLRPInstanceMutex.Lock()
	ret, specificReturn := fake.stopLRPInstanceReturnsOnCall[len(fake.stopLRPInstanceArgsForCall)]
	fake.stopLRPInstanceArgsForCall = append(fake.stopLRPInstanceArgsForCall, struct {
		logger      lager.Logger
		key         models.ActualLRPKey
		instanceKey models.ActualLRPInstanceKey
	}{logger, key, instanceKey})
	fake.recordInvocation("StopLRPInstance", []interface{}{logger, key, instanceKey})
	fake.stopLRPInstanceMutex.Unlock()
	if fake.StopLRPInstanceStub != nil {
		return fake.StopLRPInstanceStub(logger, key, instanceKey)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.stopLRPInstanceReturns.result1
}

func (fake *FakeClient) StopLRPInstanceCallCount() int {
	fake.stopLRPInstanceMutex.RLock()
	defer fake.stopLRPInstanceMutex.RUnlock()
	return len(fake.stopLRPInstanceArgsForCall)
}

func (fake *FakeClient) StopLRPInstanceArgsForCall(i int) (lager.Logger, models.ActualLRPKey, models.ActualLRPInstanceKey) {
	fake.stopLRPInstanceMutex.RLock()
	defer fake.stopLRPInstanceMutex.RUnlock()
	return fake.stopLRPInstanceArgsForCall[i].logger, fake.stopLRPInstanceArgsForCall[i].key, fake.stopLRPInstanceArgsForCall[i].instanceKey
}

func (fake *FakeClient) StopLRPInstanceReturns(result1 error) {
	fake.StopLRPInstanceStub = nil
	fake.stopLRPInstanceReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeClient) StopLRPInstanceReturnsOnCall(i int, result1 error) {
	fake.StopLRPInstanceStub = nil
	if fake.stopLRPInstanceReturnsOnCall == nil {
		fake.stopLRPInstanceReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.stopLRPInstanceReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeClient) CancelTask(logger lager.Logger, taskGuid string) error {
	fake.cancelTaskMutex.Lock()
	ret, specificReturn := fake.cancelTaskReturnsOnCall[len(fake.cancelTaskArgsForCall)]
	fake.cancelTaskArgsForCall = append(fake.cancelTaskArgsForCall, struct {
		logger   lager.Logger
		taskGuid string
	}{logger, taskGuid})
	fake.recordInvocation("CancelTask", []interface{}{logger, taskGuid})
	fake.cancelTaskMutex.Unlock()
	if fake.CancelTaskStub != nil {
		return fake.CancelTaskStub(logger, taskGuid)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.cancelTaskReturns.result1
}

func (fake *FakeClient) CancelTaskCallCount() int {
	fake.cancelTaskMutex.RLock()
	defer fake.cancelTaskMutex.RUnlock()
	return len(fake.cancelTaskArgsForCall)
}

func (fake *FakeClient) CancelTaskArgsForCall(i int) (lager.Logger, string) {
	fake.cancelTaskMutex.RLock()
	defer fake.cancelTaskMutex.RUnlock()
	return fake.cancelTaskArgsForCall[i].logger, fake.cancelTaskArgsForCall[i].taskGuid
}

func (fake *FakeClient) CancelTaskReturns(result1 error) {
	fake.CancelTaskStub = nil
	fake.cancelTaskReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeClient) CancelTaskReturnsOnCall(i int, result1 error) {
	fake.CancelTaskStub = nil
	if fake.cancelTaskReturnsOnCall == nil {
		fake.cancelTaskReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.cancelTaskReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeClient) SetStateClient(stateClient *http.Client) {
	fake.setStateClientMutex.Lock()
	fake.setStateClientArgsForCall = append(fake.setStateClientArgsForCall, struct {
		stateClient *http.Client
	}{stateClient})
	fake.recordInvocation("SetStateClient", []interface{}{stateClient})
	fake.setStateClientMutex.Unlock()
	if fake.SetStateClientStub != nil {
		fake.SetStateClientStub(stateClient)
	}
}

func (fake *FakeClient) SetStateClientCallCount() int {
	fake.setStateClientMutex.RLock()
	defer fake.setStateClientMutex.RUnlock()
	return len(fake.setStateClientArgsForCall)
}

func (fake *FakeClient) SetStateClientArgsForCall(i int) *http.Client {
	fake.setStateClientMutex.RLock()
	defer fake.setStateClientMutex.RUnlock()
	return fake.setStateClientArgsForCall[i].stateClient
}

func (fake *FakeClient) StateClientTimeout() time.Duration {
	fake.stateClientTimeoutMutex.Lock()
	ret, specificReturn := fake.stateClientTimeoutReturnsOnCall[len(fake.stateClientTimeoutArgsForCall)]
	fake.stateClientTimeoutArgsForCall = append(fake.stateClientTimeoutArgsForCall, struct{}{})
	fake.recordInvocation("StateClientTimeout", []interface{}{})
	fake.stateClientTimeoutMutex.Unlock()
	if fake.StateClientTimeoutStub != nil {
		return fake.StateClientTimeoutStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.stateClientTimeoutReturns.result1
}

func (fake *FakeClient) StateClientTimeoutCallCount() int {
	fake.stateClientTimeoutMutex.RLock()
	defer fake.stateClientTimeoutMutex.RUnlock()
	return len(fake.stateClientTimeoutArgsForCall)
}

func (fake *FakeClient) StateClientTimeoutReturns(result1 time.Duration) {
	fake.StateClientTimeoutStub = nil
	fake.stateClientTimeoutReturns = struct {
		result1 time.Duration
	}{result1}
}

func (fake *FakeClient) StateClientTimeoutReturnsOnCall(i int, result1 time.Duration) {
	fake.StateClientTimeoutStub = nil
	if fake.stateClientTimeoutReturnsOnCall == nil {
		fake.stateClientTimeoutReturnsOnCall = make(map[int]struct {
			result1 time.Duration
		})
	}
	fake.stateClientTimeoutReturnsOnCall[i] = struct {
		result1 time.Duration
	}{result1}
}

func (fake *FakeClient) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.stateMutex.RLock()
	defer fake.stateMutex.RUnlock()
	fake.performMutex.RLock()
	defer fake.performMutex.RUnlock()
	fake.stopLRPInstanceMutex.RLock()
	defer fake.stopLRPInstanceMutex.RUnlock()
	fake.cancelTaskMutex.RLock()
	defer fake.cancelTaskMutex.RUnlock()
	fake.setStateClientMutex.RLock()
	defer fake.setStateClientMutex.RUnlock()
	fake.stateClientTimeoutMutex.RLock()
	defer fake.stateClientTimeoutMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeClient) recordInvocation(key string, args []interface{}) {
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

var _ rep.Client = new(FakeClient)
