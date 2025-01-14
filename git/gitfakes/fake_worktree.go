/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by counterfeiter. DO NOT EDIT.
package gitfakes

import (
	"sync"

	gita "github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"

	"sigs.k8s.io/release-sdk/git"
)

type FakeWorktree struct {
	AddStub        func(string) (plumbing.Hash, error)
	addMutex       sync.RWMutex
	addArgsForCall []struct {
		arg1 string
	}
	addReturns struct {
		result1 plumbing.Hash
		result2 error
	}
	addReturnsOnCall map[int]struct {
		result1 plumbing.Hash
		result2 error
	}
	CheckoutStub        func(*gita.CheckoutOptions) error
	checkoutMutex       sync.RWMutex
	checkoutArgsForCall []struct {
		arg1 *gita.CheckoutOptions
	}
	checkoutReturns struct {
		result1 error
	}
	checkoutReturnsOnCall map[int]struct {
		result1 error
	}
	CommitStub        func(string, *gita.CommitOptions) (plumbing.Hash, error)
	commitMutex       sync.RWMutex
	commitArgsForCall []struct {
		arg1 string
		arg2 *gita.CommitOptions
	}
	commitReturns struct {
		result1 plumbing.Hash
		result2 error
	}
	commitReturnsOnCall map[int]struct {
		result1 plumbing.Hash
		result2 error
	}
	StatusStub        func() (gita.Status, error)
	statusMutex       sync.RWMutex
	statusArgsForCall []struct {
	}
	statusReturns struct {
		result1 gita.Status
		result2 error
	}
	statusReturnsOnCall map[int]struct {
		result1 gita.Status
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeWorktree) Add(arg1 string) (plumbing.Hash, error) {
	fake.addMutex.Lock()
	ret, specificReturn := fake.addReturnsOnCall[len(fake.addArgsForCall)]
	fake.addArgsForCall = append(fake.addArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.AddStub
	fakeReturns := fake.addReturns
	fake.recordInvocation("Add", []interface{}{arg1})
	fake.addMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeWorktree) AddCallCount() int {
	fake.addMutex.RLock()
	defer fake.addMutex.RUnlock()
	return len(fake.addArgsForCall)
}

func (fake *FakeWorktree) AddCalls(stub func(string) (plumbing.Hash, error)) {
	fake.addMutex.Lock()
	defer fake.addMutex.Unlock()
	fake.AddStub = stub
}

func (fake *FakeWorktree) AddArgsForCall(i int) string {
	fake.addMutex.RLock()
	defer fake.addMutex.RUnlock()
	argsForCall := fake.addArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeWorktree) AddReturns(result1 plumbing.Hash, result2 error) {
	fake.addMutex.Lock()
	defer fake.addMutex.Unlock()
	fake.AddStub = nil
	fake.addReturns = struct {
		result1 plumbing.Hash
		result2 error
	}{result1, result2}
}

func (fake *FakeWorktree) AddReturnsOnCall(i int, result1 plumbing.Hash, result2 error) {
	fake.addMutex.Lock()
	defer fake.addMutex.Unlock()
	fake.AddStub = nil
	if fake.addReturnsOnCall == nil {
		fake.addReturnsOnCall = make(map[int]struct {
			result1 plumbing.Hash
			result2 error
		})
	}
	fake.addReturnsOnCall[i] = struct {
		result1 plumbing.Hash
		result2 error
	}{result1, result2}
}

func (fake *FakeWorktree) Checkout(arg1 *gita.CheckoutOptions) error {
	fake.checkoutMutex.Lock()
	ret, specificReturn := fake.checkoutReturnsOnCall[len(fake.checkoutArgsForCall)]
	fake.checkoutArgsForCall = append(fake.checkoutArgsForCall, struct {
		arg1 *gita.CheckoutOptions
	}{arg1})
	stub := fake.CheckoutStub
	fakeReturns := fake.checkoutReturns
	fake.recordInvocation("Checkout", []interface{}{arg1})
	fake.checkoutMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeWorktree) CheckoutCallCount() int {
	fake.checkoutMutex.RLock()
	defer fake.checkoutMutex.RUnlock()
	return len(fake.checkoutArgsForCall)
}

func (fake *FakeWorktree) CheckoutCalls(stub func(*gita.CheckoutOptions) error) {
	fake.checkoutMutex.Lock()
	defer fake.checkoutMutex.Unlock()
	fake.CheckoutStub = stub
}

func (fake *FakeWorktree) CheckoutArgsForCall(i int) *gita.CheckoutOptions {
	fake.checkoutMutex.RLock()
	defer fake.checkoutMutex.RUnlock()
	argsForCall := fake.checkoutArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeWorktree) CheckoutReturns(result1 error) {
	fake.checkoutMutex.Lock()
	defer fake.checkoutMutex.Unlock()
	fake.CheckoutStub = nil
	fake.checkoutReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeWorktree) CheckoutReturnsOnCall(i int, result1 error) {
	fake.checkoutMutex.Lock()
	defer fake.checkoutMutex.Unlock()
	fake.CheckoutStub = nil
	if fake.checkoutReturnsOnCall == nil {
		fake.checkoutReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.checkoutReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeWorktree) Commit(arg1 string, arg2 *gita.CommitOptions) (plumbing.Hash, error) {
	fake.commitMutex.Lock()
	ret, specificReturn := fake.commitReturnsOnCall[len(fake.commitArgsForCall)]
	fake.commitArgsForCall = append(fake.commitArgsForCall, struct {
		arg1 string
		arg2 *gita.CommitOptions
	}{arg1, arg2})
	stub := fake.CommitStub
	fakeReturns := fake.commitReturns
	fake.recordInvocation("Commit", []interface{}{arg1, arg2})
	fake.commitMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeWorktree) CommitCallCount() int {
	fake.commitMutex.RLock()
	defer fake.commitMutex.RUnlock()
	return len(fake.commitArgsForCall)
}

func (fake *FakeWorktree) CommitCalls(stub func(string, *gita.CommitOptions) (plumbing.Hash, error)) {
	fake.commitMutex.Lock()
	defer fake.commitMutex.Unlock()
	fake.CommitStub = stub
}

func (fake *FakeWorktree) CommitArgsForCall(i int) (string, *gita.CommitOptions) {
	fake.commitMutex.RLock()
	defer fake.commitMutex.RUnlock()
	argsForCall := fake.commitArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeWorktree) CommitReturns(result1 plumbing.Hash, result2 error) {
	fake.commitMutex.Lock()
	defer fake.commitMutex.Unlock()
	fake.CommitStub = nil
	fake.commitReturns = struct {
		result1 plumbing.Hash
		result2 error
	}{result1, result2}
}

func (fake *FakeWorktree) CommitReturnsOnCall(i int, result1 plumbing.Hash, result2 error) {
	fake.commitMutex.Lock()
	defer fake.commitMutex.Unlock()
	fake.CommitStub = nil
	if fake.commitReturnsOnCall == nil {
		fake.commitReturnsOnCall = make(map[int]struct {
			result1 plumbing.Hash
			result2 error
		})
	}
	fake.commitReturnsOnCall[i] = struct {
		result1 plumbing.Hash
		result2 error
	}{result1, result2}
}

func (fake *FakeWorktree) Status() (gita.Status, error) {
	fake.statusMutex.Lock()
	ret, specificReturn := fake.statusReturnsOnCall[len(fake.statusArgsForCall)]
	fake.statusArgsForCall = append(fake.statusArgsForCall, struct {
	}{})
	stub := fake.StatusStub
	fakeReturns := fake.statusReturns
	fake.recordInvocation("Status", []interface{}{})
	fake.statusMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeWorktree) StatusCallCount() int {
	fake.statusMutex.RLock()
	defer fake.statusMutex.RUnlock()
	return len(fake.statusArgsForCall)
}

func (fake *FakeWorktree) StatusCalls(stub func() (gita.Status, error)) {
	fake.statusMutex.Lock()
	defer fake.statusMutex.Unlock()
	fake.StatusStub = stub
}

func (fake *FakeWorktree) StatusReturns(result1 gita.Status, result2 error) {
	fake.statusMutex.Lock()
	defer fake.statusMutex.Unlock()
	fake.StatusStub = nil
	fake.statusReturns = struct {
		result1 gita.Status
		result2 error
	}{result1, result2}
}

func (fake *FakeWorktree) StatusReturnsOnCall(i int, result1 gita.Status, result2 error) {
	fake.statusMutex.Lock()
	defer fake.statusMutex.Unlock()
	fake.StatusStub = nil
	if fake.statusReturnsOnCall == nil {
		fake.statusReturnsOnCall = make(map[int]struct {
			result1 gita.Status
			result2 error
		})
	}
	fake.statusReturnsOnCall[i] = struct {
		result1 gita.Status
		result2 error
	}{result1, result2}
}

func (fake *FakeWorktree) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.addMutex.RLock()
	defer fake.addMutex.RUnlock()
	fake.checkoutMutex.RLock()
	defer fake.checkoutMutex.RUnlock()
	fake.commitMutex.RLock()
	defer fake.commitMutex.RUnlock()
	fake.statusMutex.RLock()
	defer fake.statusMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeWorktree) recordInvocation(key string, args []interface{}) {
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

var _ git.Worktree = new(FakeWorktree)
