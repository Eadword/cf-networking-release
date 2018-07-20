// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"policy-server/db"
	"policy-server/store"
	"sync"
)

type EgressPolicyStore struct {
	AllWithTxStub        func(tx db.Transaction) ([]store.EgressPolicy, error)
	allWithTxMutex       sync.RWMutex
	allWithTxArgsForCall []struct {
		tx db.Transaction
	}
	allWithTxReturns struct {
		result1 []store.EgressPolicy
		result2 error
	}
	allWithTxReturnsOnCall map[int]struct {
		result1 []store.EgressPolicy
		result2 error
	}
	ByGuidsWithTxStub        func(tx db.Transaction, ids []string) ([]store.EgressPolicy, error)
	byGuidsWithTxMutex       sync.RWMutex
	byGuidsWithTxArgsForCall []struct {
		tx  db.Transaction
		ids []string
	}
	byGuidsWithTxReturns struct {
		result1 []store.EgressPolicy
		result2 error
	}
	byGuidsWithTxReturnsOnCall map[int]struct {
		result1 []store.EgressPolicy
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *EgressPolicyStore) AllWithTx(tx db.Transaction) ([]store.EgressPolicy, error) {
	fake.allWithTxMutex.Lock()
	ret, specificReturn := fake.allWithTxReturnsOnCall[len(fake.allWithTxArgsForCall)]
	fake.allWithTxArgsForCall = append(fake.allWithTxArgsForCall, struct {
		tx db.Transaction
	}{tx})
	fake.recordInvocation("AllWithTx", []interface{}{tx})
	fake.allWithTxMutex.Unlock()
	if fake.AllWithTxStub != nil {
		return fake.AllWithTxStub(tx)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.allWithTxReturns.result1, fake.allWithTxReturns.result2
}

func (fake *EgressPolicyStore) AllWithTxCallCount() int {
	fake.allWithTxMutex.RLock()
	defer fake.allWithTxMutex.RUnlock()
	return len(fake.allWithTxArgsForCall)
}

func (fake *EgressPolicyStore) AllWithTxArgsForCall(i int) db.Transaction {
	fake.allWithTxMutex.RLock()
	defer fake.allWithTxMutex.RUnlock()
	return fake.allWithTxArgsForCall[i].tx
}

func (fake *EgressPolicyStore) AllWithTxReturns(result1 []store.EgressPolicy, result2 error) {
	fake.AllWithTxStub = nil
	fake.allWithTxReturns = struct {
		result1 []store.EgressPolicy
		result2 error
	}{result1, result2}
}

func (fake *EgressPolicyStore) AllWithTxReturnsOnCall(i int, result1 []store.EgressPolicy, result2 error) {
	fake.AllWithTxStub = nil
	if fake.allWithTxReturnsOnCall == nil {
		fake.allWithTxReturnsOnCall = make(map[int]struct {
			result1 []store.EgressPolicy
			result2 error
		})
	}
	fake.allWithTxReturnsOnCall[i] = struct {
		result1 []store.EgressPolicy
		result2 error
	}{result1, result2}
}

func (fake *EgressPolicyStore) ByGuidsWithTx(tx db.Transaction, ids []string) ([]store.EgressPolicy, error) {
	var idsCopy []string
	if ids != nil {
		idsCopy = make([]string, len(ids))
		copy(idsCopy, ids)
	}
	fake.byGuidsWithTxMutex.Lock()
	ret, specificReturn := fake.byGuidsWithTxReturnsOnCall[len(fake.byGuidsWithTxArgsForCall)]
	fake.byGuidsWithTxArgsForCall = append(fake.byGuidsWithTxArgsForCall, struct {
		tx  db.Transaction
		ids []string
	}{tx, idsCopy})
	fake.recordInvocation("ByGuidsWithTx", []interface{}{tx, idsCopy})
	fake.byGuidsWithTxMutex.Unlock()
	if fake.ByGuidsWithTxStub != nil {
		return fake.ByGuidsWithTxStub(tx, ids)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.byGuidsWithTxReturns.result1, fake.byGuidsWithTxReturns.result2
}

func (fake *EgressPolicyStore) ByGuidsWithTxCallCount() int {
	fake.byGuidsWithTxMutex.RLock()
	defer fake.byGuidsWithTxMutex.RUnlock()
	return len(fake.byGuidsWithTxArgsForCall)
}

func (fake *EgressPolicyStore) ByGuidsWithTxArgsForCall(i int) (db.Transaction, []string) {
	fake.byGuidsWithTxMutex.RLock()
	defer fake.byGuidsWithTxMutex.RUnlock()
	return fake.byGuidsWithTxArgsForCall[i].tx, fake.byGuidsWithTxArgsForCall[i].ids
}

func (fake *EgressPolicyStore) ByGuidsWithTxReturns(result1 []store.EgressPolicy, result2 error) {
	fake.ByGuidsWithTxStub = nil
	fake.byGuidsWithTxReturns = struct {
		result1 []store.EgressPolicy
		result2 error
	}{result1, result2}
}

func (fake *EgressPolicyStore) ByGuidsWithTxReturnsOnCall(i int, result1 []store.EgressPolicy, result2 error) {
	fake.ByGuidsWithTxStub = nil
	if fake.byGuidsWithTxReturnsOnCall == nil {
		fake.byGuidsWithTxReturnsOnCall = make(map[int]struct {
			result1 []store.EgressPolicy
			result2 error
		})
	}
	fake.byGuidsWithTxReturnsOnCall[i] = struct {
		result1 []store.EgressPolicy
		result2 error
	}{result1, result2}
}

func (fake *EgressPolicyStore) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.allWithTxMutex.RLock()
	defer fake.allWithTxMutex.RUnlock()
	fake.byGuidsWithTxMutex.RLock()
	defer fake.byGuidsWithTxMutex.RUnlock()
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
