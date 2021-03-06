// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"sync"

	"github.com/cloudfoundry-incubator/bosh-backup-and-restore/ssh"
	sshcrypto "golang.org/x/crypto/ssh"
)

type FakeRemoteRunnerFactory struct {
	Stub        func(host, user, privateKey string, publicKeyCallback sshcrypto.HostKeyCallback, publicKeyAlgorithm []string, logger ssh.Logger) (ssh.RemoteRunner, error)
	mutex       sync.RWMutex
	argsForCall []struct {
		host               string
		user               string
		privateKey         string
		publicKeyCallback  sshcrypto.HostKeyCallback
		publicKeyAlgorithm []string
		logger             ssh.Logger
	}
	returns struct {
		result1 ssh.RemoteRunner
		result2 error
	}
	returnsOnCall map[int]struct {
		result1 ssh.RemoteRunner
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeRemoteRunnerFactory) Spy(host string, user string, privateKey string, publicKeyCallback sshcrypto.HostKeyCallback, publicKeyAlgorithm []string, logger ssh.Logger) (ssh.RemoteRunner, error) {
	var publicKeyAlgorithmCopy []string
	if publicKeyAlgorithm != nil {
		publicKeyAlgorithmCopy = make([]string, len(publicKeyAlgorithm))
		copy(publicKeyAlgorithmCopy, publicKeyAlgorithm)
	}
	fake.mutex.Lock()
	ret, specificReturn := fake.returnsOnCall[len(fake.argsForCall)]
	fake.argsForCall = append(fake.argsForCall, struct {
		host               string
		user               string
		privateKey         string
		publicKeyCallback  sshcrypto.HostKeyCallback
		publicKeyAlgorithm []string
		logger             ssh.Logger
	}{host, user, privateKey, publicKeyCallback, publicKeyAlgorithmCopy, logger})
	fake.recordInvocation("RemoteRunnerFactory", []interface{}{host, user, privateKey, publicKeyCallback, publicKeyAlgorithmCopy, logger})
	fake.mutex.Unlock()
	if fake.Stub != nil {
		return fake.Stub(host, user, privateKey, publicKeyCallback, publicKeyAlgorithm, logger)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.returns.result1, fake.returns.result2
}

func (fake *FakeRemoteRunnerFactory) CallCount() int {
	fake.mutex.RLock()
	defer fake.mutex.RUnlock()
	return len(fake.argsForCall)
}

func (fake *FakeRemoteRunnerFactory) ArgsForCall(i int) (string, string, string, sshcrypto.HostKeyCallback, []string, ssh.Logger) {
	fake.mutex.RLock()
	defer fake.mutex.RUnlock()
	return fake.argsForCall[i].host, fake.argsForCall[i].user, fake.argsForCall[i].privateKey, fake.argsForCall[i].publicKeyCallback, fake.argsForCall[i].publicKeyAlgorithm, fake.argsForCall[i].logger
}

func (fake *FakeRemoteRunnerFactory) Returns(result1 ssh.RemoteRunner, result2 error) {
	fake.Stub = nil
	fake.returns = struct {
		result1 ssh.RemoteRunner
		result2 error
	}{result1, result2}
}

func (fake *FakeRemoteRunnerFactory) ReturnsOnCall(i int, result1 ssh.RemoteRunner, result2 error) {
	fake.Stub = nil
	if fake.returnsOnCall == nil {
		fake.returnsOnCall = make(map[int]struct {
			result1 ssh.RemoteRunner
			result2 error
		})
	}
	fake.returnsOnCall[i] = struct {
		result1 ssh.RemoteRunner
		result2 error
	}{result1, result2}
}

func (fake *FakeRemoteRunnerFactory) Invocations() map[string][][]interface{} {
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

func (fake *FakeRemoteRunnerFactory) recordInvocation(key string, args []interface{}) {
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

var _ ssh.RemoteRunnerFactory = new(FakeRemoteRunnerFactory).Spy
