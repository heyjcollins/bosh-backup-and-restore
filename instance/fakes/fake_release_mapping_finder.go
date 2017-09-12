// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"sync"

	"github.com/cloudfoundry-incubator/bosh-backup-and-restore/instance"
)

type FakeReleaseMappingFinder struct {
	Stub        func(manifest string) (instance.ReleaseMapping, error)
	mutex       sync.RWMutex
	argsForCall []struct {
		manifest string
	}
	returns struct {
		result1 instance.ReleaseMapping
		result2 error
	}
	returnsOnCall map[int]struct {
		result1 instance.ReleaseMapping
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeReleaseMappingFinder) Spy(manifest string) (instance.ReleaseMapping, error) {
	fake.mutex.Lock()
	ret, specificReturn := fake.returnsOnCall[len(fake.argsForCall)]
	fake.argsForCall = append(fake.argsForCall, struct {
		manifest string
	}{manifest})
	fake.recordInvocation("ReleaseMappingFinder", []interface{}{manifest})
	fake.mutex.Unlock()
	if fake.Stub != nil {
		return fake.Stub(manifest)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.returns.result1, fake.returns.result2
}

func (fake *FakeReleaseMappingFinder) CallCount() int {
	fake.mutex.RLock()
	defer fake.mutex.RUnlock()
	return len(fake.argsForCall)
}

func (fake *FakeReleaseMappingFinder) ArgsForCall(i int) string {
	fake.mutex.RLock()
	defer fake.mutex.RUnlock()
	return fake.argsForCall[i].manifest
}

func (fake *FakeReleaseMappingFinder) Returns(result1 instance.ReleaseMapping, result2 error) {
	fake.Stub = nil
	fake.returns = struct {
		result1 instance.ReleaseMapping
		result2 error
	}{result1, result2}
}

func (fake *FakeReleaseMappingFinder) ReturnsOnCall(i int, result1 instance.ReleaseMapping, result2 error) {
	fake.Stub = nil
	if fake.returnsOnCall == nil {
		fake.returnsOnCall = make(map[int]struct {
			result1 instance.ReleaseMapping
			result2 error
		})
	}
	fake.returnsOnCall[i] = struct {
		result1 instance.ReleaseMapping
		result2 error
	}{result1, result2}
}

func (fake *FakeReleaseMappingFinder) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.mutex.RLock()
	defer fake.mutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeReleaseMappingFinder) recordInvocation(key string, args []interface{}) {
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

var _ instance.ReleaseMappingFinder = new(FakeReleaseMappingFinder).Spy