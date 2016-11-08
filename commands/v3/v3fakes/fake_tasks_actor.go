// This file was generated by counterfeiter
package v3fakes

import (
	"sync"

	"code.cloudfoundry.org/cli/actors/v3actions"
	"code.cloudfoundry.org/cli/commands/v3"
)

type FakeTasksActor struct {
	GetApplicationByNameAndSpaceStub        func(appName string, spaceGUID string) (v3actions.Application, v3actions.Warnings, error)
	getApplicationByNameAndSpaceMutex       sync.RWMutex
	getApplicationByNameAndSpaceArgsForCall []struct {
		appName   string
		spaceGUID string
	}
	getApplicationByNameAndSpaceReturns struct {
		result1 v3actions.Application
		result2 v3actions.Warnings
		result3 error
	}
	GetApplicationTasksStub        func(appGUID string) ([]v3actions.Task, v3actions.Warnings, error)
	getApplicationTasksMutex       sync.RWMutex
	getApplicationTasksArgsForCall []struct {
		appGUID string
	}
	getApplicationTasksReturns struct {
		result1 []v3actions.Task
		result2 v3actions.Warnings
		result3 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeTasksActor) GetApplicationByNameAndSpace(appName string, spaceGUID string) (v3actions.Application, v3actions.Warnings, error) {
	fake.getApplicationByNameAndSpaceMutex.Lock()
	fake.getApplicationByNameAndSpaceArgsForCall = append(fake.getApplicationByNameAndSpaceArgsForCall, struct {
		appName   string
		spaceGUID string
	}{appName, spaceGUID})
	fake.recordInvocation("GetApplicationByNameAndSpace", []interface{}{appName, spaceGUID})
	fake.getApplicationByNameAndSpaceMutex.Unlock()
	if fake.GetApplicationByNameAndSpaceStub != nil {
		return fake.GetApplicationByNameAndSpaceStub(appName, spaceGUID)
	} else {
		return fake.getApplicationByNameAndSpaceReturns.result1, fake.getApplicationByNameAndSpaceReturns.result2, fake.getApplicationByNameAndSpaceReturns.result3
	}
}

func (fake *FakeTasksActor) GetApplicationByNameAndSpaceCallCount() int {
	fake.getApplicationByNameAndSpaceMutex.RLock()
	defer fake.getApplicationByNameAndSpaceMutex.RUnlock()
	return len(fake.getApplicationByNameAndSpaceArgsForCall)
}

func (fake *FakeTasksActor) GetApplicationByNameAndSpaceArgsForCall(i int) (string, string) {
	fake.getApplicationByNameAndSpaceMutex.RLock()
	defer fake.getApplicationByNameAndSpaceMutex.RUnlock()
	return fake.getApplicationByNameAndSpaceArgsForCall[i].appName, fake.getApplicationByNameAndSpaceArgsForCall[i].spaceGUID
}

func (fake *FakeTasksActor) GetApplicationByNameAndSpaceReturns(result1 v3actions.Application, result2 v3actions.Warnings, result3 error) {
	fake.GetApplicationByNameAndSpaceStub = nil
	fake.getApplicationByNameAndSpaceReturns = struct {
		result1 v3actions.Application
		result2 v3actions.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeTasksActor) GetApplicationTasks(appGUID string) ([]v3actions.Task, v3actions.Warnings, error) {
	fake.getApplicationTasksMutex.Lock()
	fake.getApplicationTasksArgsForCall = append(fake.getApplicationTasksArgsForCall, struct {
		appGUID string
	}{appGUID})
	fake.recordInvocation("GetApplicationTasks", []interface{}{appGUID})
	fake.getApplicationTasksMutex.Unlock()
	if fake.GetApplicationTasksStub != nil {
		return fake.GetApplicationTasksStub(appGUID)
	} else {
		return fake.getApplicationTasksReturns.result1, fake.getApplicationTasksReturns.result2, fake.getApplicationTasksReturns.result3
	}
}

func (fake *FakeTasksActor) GetApplicationTasksCallCount() int {
	fake.getApplicationTasksMutex.RLock()
	defer fake.getApplicationTasksMutex.RUnlock()
	return len(fake.getApplicationTasksArgsForCall)
}

func (fake *FakeTasksActor) GetApplicationTasksArgsForCall(i int) string {
	fake.getApplicationTasksMutex.RLock()
	defer fake.getApplicationTasksMutex.RUnlock()
	return fake.getApplicationTasksArgsForCall[i].appGUID
}

func (fake *FakeTasksActor) GetApplicationTasksReturns(result1 []v3actions.Task, result2 v3actions.Warnings, result3 error) {
	fake.GetApplicationTasksStub = nil
	fake.getApplicationTasksReturns = struct {
		result1 []v3actions.Task
		result2 v3actions.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeTasksActor) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getApplicationByNameAndSpaceMutex.RLock()
	defer fake.getApplicationByNameAndSpaceMutex.RUnlock()
	fake.getApplicationTasksMutex.RLock()
	defer fake.getApplicationTasksMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeTasksActor) recordInvocation(key string, args []interface{}) {
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

var _ v3.TasksActor = new(FakeTasksActor)
