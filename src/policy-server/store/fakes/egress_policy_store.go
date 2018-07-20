// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"policy-server/db"
	"policy-server/store"
	"sync"
)

type EgressPolicyStore struct {
	CreateWithTxStub        func(db.Transaction, []store.EgressPolicy) error
	createWithTxMutex       sync.RWMutex
	createWithTxArgsForCall []struct {
		arg1 db.Transaction
		arg2 []store.EgressPolicy
	}
	createWithTxReturns struct {
		result1 error
	}
	createWithTxReturnsOnCall map[int]struct {
		result1 error
	}
	DeleteWithTxStub        func(db.Transaction, []store.EgressPolicy) error
	deleteWithTxMutex       sync.RWMutex
	deleteWithTxArgsForCall []struct {
		arg1 db.Transaction
		arg2 []store.EgressPolicy
	}
	deleteWithTxReturns struct {
		result1 error
	}
	deleteWithTxReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *EgressPolicyStore) CreateWithTx(arg1 db.Transaction, arg2 []store.EgressPolicy) error {
	var arg2Copy []store.EgressPolicy
	if arg2 != nil {
		arg2Copy = make([]store.EgressPolicy, len(arg2))
		copy(arg2Copy, arg2)
	}
	fake.createWithTxMutex.Lock()
	ret, specificReturn := fake.createWithTxReturnsOnCall[len(fake.createWithTxArgsForCall)]
	fake.createWithTxArgsForCall = append(fake.createWithTxArgsForCall, struct {
		arg1 db.Transaction
		arg2 []store.EgressPolicy
	}{arg1, arg2Copy})
	fake.recordInvocation("CreateWithTx", []interface{}{arg1, arg2Copy})
	fake.createWithTxMutex.Unlock()
	if fake.CreateWithTxStub != nil {
		return fake.CreateWithTxStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.createWithTxReturns.result1
}

func (fake *EgressPolicyStore) CreateWithTxCallCount() int {
	fake.createWithTxMutex.RLock()
	defer fake.createWithTxMutex.RUnlock()
	return len(fake.createWithTxArgsForCall)
}

func (fake *EgressPolicyStore) CreateWithTxArgsForCall(i int) (db.Transaction, []store.EgressPolicy) {
	fake.createWithTxMutex.RLock()
	defer fake.createWithTxMutex.RUnlock()
	return fake.createWithTxArgsForCall[i].arg1, fake.createWithTxArgsForCall[i].arg2
}

func (fake *EgressPolicyStore) CreateWithTxReturns(result1 error) {
	fake.CreateWithTxStub = nil
	fake.createWithTxReturns = struct {
		result1 error
	}{result1}
}

func (fake *EgressPolicyStore) CreateWithTxReturnsOnCall(i int, result1 error) {
	fake.CreateWithTxStub = nil
	if fake.createWithTxReturnsOnCall == nil {
		fake.createWithTxReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.createWithTxReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *EgressPolicyStore) DeleteWithTx(arg1 db.Transaction, arg2 []store.EgressPolicy) error {
	var arg2Copy []store.EgressPolicy
	if arg2 != nil {
		arg2Copy = make([]store.EgressPolicy, len(arg2))
		copy(arg2Copy, arg2)
	}
	fake.deleteWithTxMutex.Lock()
	ret, specificReturn := fake.deleteWithTxReturnsOnCall[len(fake.deleteWithTxArgsForCall)]
	fake.deleteWithTxArgsForCall = append(fake.deleteWithTxArgsForCall, struct {
		arg1 db.Transaction
		arg2 []store.EgressPolicy
	}{arg1, arg2Copy})
	fake.recordInvocation("DeleteWithTx", []interface{}{arg1, arg2Copy})
	fake.deleteWithTxMutex.Unlock()
	if fake.DeleteWithTxStub != nil {
		return fake.DeleteWithTxStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.deleteWithTxReturns.result1
}

func (fake *EgressPolicyStore) DeleteWithTxCallCount() int {
	fake.deleteWithTxMutex.RLock()
	defer fake.deleteWithTxMutex.RUnlock()
	return len(fake.deleteWithTxArgsForCall)
}

func (fake *EgressPolicyStore) DeleteWithTxArgsForCall(i int) (db.Transaction, []store.EgressPolicy) {
	fake.deleteWithTxMutex.RLock()
	defer fake.deleteWithTxMutex.RUnlock()
	return fake.deleteWithTxArgsForCall[i].arg1, fake.deleteWithTxArgsForCall[i].arg2
}

func (fake *EgressPolicyStore) DeleteWithTxReturns(result1 error) {
	fake.DeleteWithTxStub = nil
	fake.deleteWithTxReturns = struct {
		result1 error
	}{result1}
}

func (fake *EgressPolicyStore) DeleteWithTxReturnsOnCall(i int, result1 error) {
	fake.DeleteWithTxStub = nil
	if fake.deleteWithTxReturnsOnCall == nil {
		fake.deleteWithTxReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.deleteWithTxReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *EgressPolicyStore) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.createWithTxMutex.RLock()
	defer fake.createWithTxMutex.RUnlock()
	fake.deleteWithTxMutex.RLock()
	defer fake.deleteWithTxMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *EgressPolicyStore) recordInvocation(key string, args []interface{}) {
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
