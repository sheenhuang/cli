// Code generated by counterfeiter. DO NOT EDIT.
package v2v3actionfakes

import (
	"sync"

	"code.cloudfoundry.org/cli/actor/v2action"
	"code.cloudfoundry.org/cli/actor/v2v3action"
	"code.cloudfoundry.org/cli/util/manifest"
)

type FakeV2Actor struct {
	CreateApplicationManifestByNameAndSpaceStub        func(string, string) (manifest.Application, v2action.Warnings, error)
	createApplicationManifestByNameAndSpaceMutex       sync.RWMutex
	createApplicationManifestByNameAndSpaceArgsForCall []struct {
		arg1 string
		arg2 string
	}
	createApplicationManifestByNameAndSpaceReturns struct {
		result1 manifest.Application
		result2 v2action.Warnings
		result3 error
	}
	createApplicationManifestByNameAndSpaceReturnsOnCall map[int]struct {
		result1 manifest.Application
		result2 v2action.Warnings
		result3 error
	}
	GetApplicationInstancesWithStatsByApplicationStub        func(guid string) ([]v2action.ApplicationInstanceWithStats, v2action.Warnings, error)
	getApplicationInstancesWithStatsByApplicationMutex       sync.RWMutex
	getApplicationInstancesWithStatsByApplicationArgsForCall []struct {
		guid string
	}
	getApplicationInstancesWithStatsByApplicationReturns struct {
		result1 []v2action.ApplicationInstanceWithStats
		result2 v2action.Warnings
		result3 error
	}
	getApplicationInstancesWithStatsByApplicationReturnsOnCall map[int]struct {
		result1 []v2action.ApplicationInstanceWithStats
		result2 v2action.Warnings
		result3 error
	}
	GetApplicationRoutesStub        func(appGUID string) (v2action.Routes, v2action.Warnings, error)
	getApplicationRoutesMutex       sync.RWMutex
	getApplicationRoutesArgsForCall []struct {
		appGUID string
	}
	getApplicationRoutesReturns struct {
		result1 v2action.Routes
		result2 v2action.Warnings
		result3 error
	}
	getApplicationRoutesReturnsOnCall map[int]struct {
		result1 v2action.Routes
		result2 v2action.Warnings
		result3 error
	}
	GetFeatureFlagsStub        func() ([]v2action.FeatureFlag, v2action.Warnings, error)
	getFeatureFlagsMutex       sync.RWMutex
	getFeatureFlagsArgsForCall []struct{}
	getFeatureFlagsReturns     struct {
		result1 []v2action.FeatureFlag
		result2 v2action.Warnings
		result3 error
	}
	getFeatureFlagsReturnsOnCall map[int]struct {
		result1 []v2action.FeatureFlag
		result2 v2action.Warnings
		result3 error
	}
	GetServiceStub        func(serviceGUID string) (v2action.Service, v2action.Warnings, error)
	getServiceMutex       sync.RWMutex
	getServiceArgsForCall []struct {
		serviceGUID string
	}
	getServiceReturns struct {
		result1 v2action.Service
		result2 v2action.Warnings
		result3 error
	}
	getServiceReturnsOnCall map[int]struct {
		result1 v2action.Service
		result2 v2action.Warnings
		result3 error
	}
	GetServiceInstanceByNameAndSpaceStub        func(serviceInstanceName string, spaceGUID string) (v2action.ServiceInstance, v2action.Warnings, error)
	getServiceInstanceByNameAndSpaceMutex       sync.RWMutex
	getServiceInstanceByNameAndSpaceArgsForCall []struct {
		serviceInstanceName string
		spaceGUID           string
	}
	getServiceInstanceByNameAndSpaceReturns struct {
		result1 v2action.ServiceInstance
		result2 v2action.Warnings
		result3 error
	}
	getServiceInstanceByNameAndSpaceReturnsOnCall map[int]struct {
		result1 v2action.ServiceInstance
		result2 v2action.Warnings
		result3 error
	}
	GetServiceInstanceSharedTosByServiceInstanceStub        func(serviceInstanceGUID string) ([]v2action.ServiceInstanceSharedTo, v2action.Warnings, error)
	getServiceInstanceSharedTosByServiceInstanceMutex       sync.RWMutex
	getServiceInstanceSharedTosByServiceInstanceArgsForCall []struct {
		serviceInstanceGUID string
	}
	getServiceInstanceSharedTosByServiceInstanceReturns struct {
		result1 []v2action.ServiceInstanceSharedTo
		result2 v2action.Warnings
		result3 error
	}
	getServiceInstanceSharedTosByServiceInstanceReturnsOnCall map[int]struct {
		result1 []v2action.ServiceInstanceSharedTo
		result2 v2action.Warnings
		result3 error
	}
	GetSpaceByOrganizationAndNameStub        func(orgGUID string, spaceName string) (v2action.Space, v2action.Warnings, error)
	getSpaceByOrganizationAndNameMutex       sync.RWMutex
	getSpaceByOrganizationAndNameArgsForCall []struct {
		orgGUID   string
		spaceName string
	}
	getSpaceByOrganizationAndNameReturns struct {
		result1 v2action.Space
		result2 v2action.Warnings
		result3 error
	}
	getSpaceByOrganizationAndNameReturnsOnCall map[int]struct {
		result1 v2action.Space
		result2 v2action.Warnings
		result3 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeV2Actor) CreateApplicationManifestByNameAndSpace(arg1 string, arg2 string) (manifest.Application, v2action.Warnings, error) {
	fake.createApplicationManifestByNameAndSpaceMutex.Lock()
	ret, specificReturn := fake.createApplicationManifestByNameAndSpaceReturnsOnCall[len(fake.createApplicationManifestByNameAndSpaceArgsForCall)]
	fake.createApplicationManifestByNameAndSpaceArgsForCall = append(fake.createApplicationManifestByNameAndSpaceArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("CreateApplicationManifestByNameAndSpace", []interface{}{arg1, arg2})
	fake.createApplicationManifestByNameAndSpaceMutex.Unlock()
	if fake.CreateApplicationManifestByNameAndSpaceStub != nil {
		return fake.CreateApplicationManifestByNameAndSpaceStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	return fake.createApplicationManifestByNameAndSpaceReturns.result1, fake.createApplicationManifestByNameAndSpaceReturns.result2, fake.createApplicationManifestByNameAndSpaceReturns.result3
}

func (fake *FakeV2Actor) CreateApplicationManifestByNameAndSpaceCallCount() int {
	fake.createApplicationManifestByNameAndSpaceMutex.RLock()
	defer fake.createApplicationManifestByNameAndSpaceMutex.RUnlock()
	return len(fake.createApplicationManifestByNameAndSpaceArgsForCall)
}

func (fake *FakeV2Actor) CreateApplicationManifestByNameAndSpaceArgsForCall(i int) (string, string) {
	fake.createApplicationManifestByNameAndSpaceMutex.RLock()
	defer fake.createApplicationManifestByNameAndSpaceMutex.RUnlock()
	return fake.createApplicationManifestByNameAndSpaceArgsForCall[i].arg1, fake.createApplicationManifestByNameAndSpaceArgsForCall[i].arg2
}

func (fake *FakeV2Actor) CreateApplicationManifestByNameAndSpaceReturns(result1 manifest.Application, result2 v2action.Warnings, result3 error) {
	fake.CreateApplicationManifestByNameAndSpaceStub = nil
	fake.createApplicationManifestByNameAndSpaceReturns = struct {
		result1 manifest.Application
		result2 v2action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeV2Actor) CreateApplicationManifestByNameAndSpaceReturnsOnCall(i int, result1 manifest.Application, result2 v2action.Warnings, result3 error) {
	fake.CreateApplicationManifestByNameAndSpaceStub = nil
	if fake.createApplicationManifestByNameAndSpaceReturnsOnCall == nil {
		fake.createApplicationManifestByNameAndSpaceReturnsOnCall = make(map[int]struct {
			result1 manifest.Application
			result2 v2action.Warnings
			result3 error
		})
	}
	fake.createApplicationManifestByNameAndSpaceReturnsOnCall[i] = struct {
		result1 manifest.Application
		result2 v2action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeV2Actor) GetApplicationInstancesWithStatsByApplication(guid string) ([]v2action.ApplicationInstanceWithStats, v2action.Warnings, error) {
	fake.getApplicationInstancesWithStatsByApplicationMutex.Lock()
	ret, specificReturn := fake.getApplicationInstancesWithStatsByApplicationReturnsOnCall[len(fake.getApplicationInstancesWithStatsByApplicationArgsForCall)]
	fake.getApplicationInstancesWithStatsByApplicationArgsForCall = append(fake.getApplicationInstancesWithStatsByApplicationArgsForCall, struct {
		guid string
	}{guid})
	fake.recordInvocation("GetApplicationInstancesWithStatsByApplication", []interface{}{guid})
	fake.getApplicationInstancesWithStatsByApplicationMutex.Unlock()
	if fake.GetApplicationInstancesWithStatsByApplicationStub != nil {
		return fake.GetApplicationInstancesWithStatsByApplicationStub(guid)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	return fake.getApplicationInstancesWithStatsByApplicationReturns.result1, fake.getApplicationInstancesWithStatsByApplicationReturns.result2, fake.getApplicationInstancesWithStatsByApplicationReturns.result3
}

func (fake *FakeV2Actor) GetApplicationInstancesWithStatsByApplicationCallCount() int {
	fake.getApplicationInstancesWithStatsByApplicationMutex.RLock()
	defer fake.getApplicationInstancesWithStatsByApplicationMutex.RUnlock()
	return len(fake.getApplicationInstancesWithStatsByApplicationArgsForCall)
}

func (fake *FakeV2Actor) GetApplicationInstancesWithStatsByApplicationArgsForCall(i int) string {
	fake.getApplicationInstancesWithStatsByApplicationMutex.RLock()
	defer fake.getApplicationInstancesWithStatsByApplicationMutex.RUnlock()
	return fake.getApplicationInstancesWithStatsByApplicationArgsForCall[i].guid
}

func (fake *FakeV2Actor) GetApplicationInstancesWithStatsByApplicationReturns(result1 []v2action.ApplicationInstanceWithStats, result2 v2action.Warnings, result3 error) {
	fake.GetApplicationInstancesWithStatsByApplicationStub = nil
	fake.getApplicationInstancesWithStatsByApplicationReturns = struct {
		result1 []v2action.ApplicationInstanceWithStats
		result2 v2action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeV2Actor) GetApplicationInstancesWithStatsByApplicationReturnsOnCall(i int, result1 []v2action.ApplicationInstanceWithStats, result2 v2action.Warnings, result3 error) {
	fake.GetApplicationInstancesWithStatsByApplicationStub = nil
	if fake.getApplicationInstancesWithStatsByApplicationReturnsOnCall == nil {
		fake.getApplicationInstancesWithStatsByApplicationReturnsOnCall = make(map[int]struct {
			result1 []v2action.ApplicationInstanceWithStats
			result2 v2action.Warnings
			result3 error
		})
	}
	fake.getApplicationInstancesWithStatsByApplicationReturnsOnCall[i] = struct {
		result1 []v2action.ApplicationInstanceWithStats
		result2 v2action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeV2Actor) GetApplicationRoutes(appGUID string) (v2action.Routes, v2action.Warnings, error) {
	fake.getApplicationRoutesMutex.Lock()
	ret, specificReturn := fake.getApplicationRoutesReturnsOnCall[len(fake.getApplicationRoutesArgsForCall)]
	fake.getApplicationRoutesArgsForCall = append(fake.getApplicationRoutesArgsForCall, struct {
		appGUID string
	}{appGUID})
	fake.recordInvocation("GetApplicationRoutes", []interface{}{appGUID})
	fake.getApplicationRoutesMutex.Unlock()
	if fake.GetApplicationRoutesStub != nil {
		return fake.GetApplicationRoutesStub(appGUID)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	return fake.getApplicationRoutesReturns.result1, fake.getApplicationRoutesReturns.result2, fake.getApplicationRoutesReturns.result3
}

func (fake *FakeV2Actor) GetApplicationRoutesCallCount() int {
	fake.getApplicationRoutesMutex.RLock()
	defer fake.getApplicationRoutesMutex.RUnlock()
	return len(fake.getApplicationRoutesArgsForCall)
}

func (fake *FakeV2Actor) GetApplicationRoutesArgsForCall(i int) string {
	fake.getApplicationRoutesMutex.RLock()
	defer fake.getApplicationRoutesMutex.RUnlock()
	return fake.getApplicationRoutesArgsForCall[i].appGUID
}

func (fake *FakeV2Actor) GetApplicationRoutesReturns(result1 v2action.Routes, result2 v2action.Warnings, result3 error) {
	fake.GetApplicationRoutesStub = nil
	fake.getApplicationRoutesReturns = struct {
		result1 v2action.Routes
		result2 v2action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeV2Actor) GetApplicationRoutesReturnsOnCall(i int, result1 v2action.Routes, result2 v2action.Warnings, result3 error) {
	fake.GetApplicationRoutesStub = nil
	if fake.getApplicationRoutesReturnsOnCall == nil {
		fake.getApplicationRoutesReturnsOnCall = make(map[int]struct {
			result1 v2action.Routes
			result2 v2action.Warnings
			result3 error
		})
	}
	fake.getApplicationRoutesReturnsOnCall[i] = struct {
		result1 v2action.Routes
		result2 v2action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeV2Actor) GetFeatureFlags() ([]v2action.FeatureFlag, v2action.Warnings, error) {
	fake.getFeatureFlagsMutex.Lock()
	ret, specificReturn := fake.getFeatureFlagsReturnsOnCall[len(fake.getFeatureFlagsArgsForCall)]
	fake.getFeatureFlagsArgsForCall = append(fake.getFeatureFlagsArgsForCall, struct{}{})
	fake.recordInvocation("GetFeatureFlags", []interface{}{})
	fake.getFeatureFlagsMutex.Unlock()
	if fake.GetFeatureFlagsStub != nil {
		return fake.GetFeatureFlagsStub()
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	return fake.getFeatureFlagsReturns.result1, fake.getFeatureFlagsReturns.result2, fake.getFeatureFlagsReturns.result3
}

func (fake *FakeV2Actor) GetFeatureFlagsCallCount() int {
	fake.getFeatureFlagsMutex.RLock()
	defer fake.getFeatureFlagsMutex.RUnlock()
	return len(fake.getFeatureFlagsArgsForCall)
}

func (fake *FakeV2Actor) GetFeatureFlagsReturns(result1 []v2action.FeatureFlag, result2 v2action.Warnings, result3 error) {
	fake.GetFeatureFlagsStub = nil
	fake.getFeatureFlagsReturns = struct {
		result1 []v2action.FeatureFlag
		result2 v2action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeV2Actor) GetFeatureFlagsReturnsOnCall(i int, result1 []v2action.FeatureFlag, result2 v2action.Warnings, result3 error) {
	fake.GetFeatureFlagsStub = nil
	if fake.getFeatureFlagsReturnsOnCall == nil {
		fake.getFeatureFlagsReturnsOnCall = make(map[int]struct {
			result1 []v2action.FeatureFlag
			result2 v2action.Warnings
			result3 error
		})
	}
	fake.getFeatureFlagsReturnsOnCall[i] = struct {
		result1 []v2action.FeatureFlag
		result2 v2action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeV2Actor) GetService(serviceGUID string) (v2action.Service, v2action.Warnings, error) {
	fake.getServiceMutex.Lock()
	ret, specificReturn := fake.getServiceReturnsOnCall[len(fake.getServiceArgsForCall)]
	fake.getServiceArgsForCall = append(fake.getServiceArgsForCall, struct {
		serviceGUID string
	}{serviceGUID})
	fake.recordInvocation("GetService", []interface{}{serviceGUID})
	fake.getServiceMutex.Unlock()
	if fake.GetServiceStub != nil {
		return fake.GetServiceStub(serviceGUID)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	return fake.getServiceReturns.result1, fake.getServiceReturns.result2, fake.getServiceReturns.result3
}

func (fake *FakeV2Actor) GetServiceCallCount() int {
	fake.getServiceMutex.RLock()
	defer fake.getServiceMutex.RUnlock()
	return len(fake.getServiceArgsForCall)
}

func (fake *FakeV2Actor) GetServiceArgsForCall(i int) string {
	fake.getServiceMutex.RLock()
	defer fake.getServiceMutex.RUnlock()
	return fake.getServiceArgsForCall[i].serviceGUID
}

func (fake *FakeV2Actor) GetServiceReturns(result1 v2action.Service, result2 v2action.Warnings, result3 error) {
	fake.GetServiceStub = nil
	fake.getServiceReturns = struct {
		result1 v2action.Service
		result2 v2action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeV2Actor) GetServiceReturnsOnCall(i int, result1 v2action.Service, result2 v2action.Warnings, result3 error) {
	fake.GetServiceStub = nil
	if fake.getServiceReturnsOnCall == nil {
		fake.getServiceReturnsOnCall = make(map[int]struct {
			result1 v2action.Service
			result2 v2action.Warnings
			result3 error
		})
	}
	fake.getServiceReturnsOnCall[i] = struct {
		result1 v2action.Service
		result2 v2action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeV2Actor) GetServiceInstanceByNameAndSpace(serviceInstanceName string, spaceGUID string) (v2action.ServiceInstance, v2action.Warnings, error) {
	fake.getServiceInstanceByNameAndSpaceMutex.Lock()
	ret, specificReturn := fake.getServiceInstanceByNameAndSpaceReturnsOnCall[len(fake.getServiceInstanceByNameAndSpaceArgsForCall)]
	fake.getServiceInstanceByNameAndSpaceArgsForCall = append(fake.getServiceInstanceByNameAndSpaceArgsForCall, struct {
		serviceInstanceName string
		spaceGUID           string
	}{serviceInstanceName, spaceGUID})
	fake.recordInvocation("GetServiceInstanceByNameAndSpace", []interface{}{serviceInstanceName, spaceGUID})
	fake.getServiceInstanceByNameAndSpaceMutex.Unlock()
	if fake.GetServiceInstanceByNameAndSpaceStub != nil {
		return fake.GetServiceInstanceByNameAndSpaceStub(serviceInstanceName, spaceGUID)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	return fake.getServiceInstanceByNameAndSpaceReturns.result1, fake.getServiceInstanceByNameAndSpaceReturns.result2, fake.getServiceInstanceByNameAndSpaceReturns.result3
}

func (fake *FakeV2Actor) GetServiceInstanceByNameAndSpaceCallCount() int {
	fake.getServiceInstanceByNameAndSpaceMutex.RLock()
	defer fake.getServiceInstanceByNameAndSpaceMutex.RUnlock()
	return len(fake.getServiceInstanceByNameAndSpaceArgsForCall)
}

func (fake *FakeV2Actor) GetServiceInstanceByNameAndSpaceArgsForCall(i int) (string, string) {
	fake.getServiceInstanceByNameAndSpaceMutex.RLock()
	defer fake.getServiceInstanceByNameAndSpaceMutex.RUnlock()
	return fake.getServiceInstanceByNameAndSpaceArgsForCall[i].serviceInstanceName, fake.getServiceInstanceByNameAndSpaceArgsForCall[i].spaceGUID
}

func (fake *FakeV2Actor) GetServiceInstanceByNameAndSpaceReturns(result1 v2action.ServiceInstance, result2 v2action.Warnings, result3 error) {
	fake.GetServiceInstanceByNameAndSpaceStub = nil
	fake.getServiceInstanceByNameAndSpaceReturns = struct {
		result1 v2action.ServiceInstance
		result2 v2action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeV2Actor) GetServiceInstanceByNameAndSpaceReturnsOnCall(i int, result1 v2action.ServiceInstance, result2 v2action.Warnings, result3 error) {
	fake.GetServiceInstanceByNameAndSpaceStub = nil
	if fake.getServiceInstanceByNameAndSpaceReturnsOnCall == nil {
		fake.getServiceInstanceByNameAndSpaceReturnsOnCall = make(map[int]struct {
			result1 v2action.ServiceInstance
			result2 v2action.Warnings
			result3 error
		})
	}
	fake.getServiceInstanceByNameAndSpaceReturnsOnCall[i] = struct {
		result1 v2action.ServiceInstance
		result2 v2action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeV2Actor) GetServiceInstanceSharedTosByServiceInstance(serviceInstanceGUID string) ([]v2action.ServiceInstanceSharedTo, v2action.Warnings, error) {
	fake.getServiceInstanceSharedTosByServiceInstanceMutex.Lock()
	ret, specificReturn := fake.getServiceInstanceSharedTosByServiceInstanceReturnsOnCall[len(fake.getServiceInstanceSharedTosByServiceInstanceArgsForCall)]
	fake.getServiceInstanceSharedTosByServiceInstanceArgsForCall = append(fake.getServiceInstanceSharedTosByServiceInstanceArgsForCall, struct {
		serviceInstanceGUID string
	}{serviceInstanceGUID})
	fake.recordInvocation("GetServiceInstanceSharedTosByServiceInstance", []interface{}{serviceInstanceGUID})
	fake.getServiceInstanceSharedTosByServiceInstanceMutex.Unlock()
	if fake.GetServiceInstanceSharedTosByServiceInstanceStub != nil {
		return fake.GetServiceInstanceSharedTosByServiceInstanceStub(serviceInstanceGUID)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	return fake.getServiceInstanceSharedTosByServiceInstanceReturns.result1, fake.getServiceInstanceSharedTosByServiceInstanceReturns.result2, fake.getServiceInstanceSharedTosByServiceInstanceReturns.result3
}

func (fake *FakeV2Actor) GetServiceInstanceSharedTosByServiceInstanceCallCount() int {
	fake.getServiceInstanceSharedTosByServiceInstanceMutex.RLock()
	defer fake.getServiceInstanceSharedTosByServiceInstanceMutex.RUnlock()
	return len(fake.getServiceInstanceSharedTosByServiceInstanceArgsForCall)
}

func (fake *FakeV2Actor) GetServiceInstanceSharedTosByServiceInstanceArgsForCall(i int) string {
	fake.getServiceInstanceSharedTosByServiceInstanceMutex.RLock()
	defer fake.getServiceInstanceSharedTosByServiceInstanceMutex.RUnlock()
	return fake.getServiceInstanceSharedTosByServiceInstanceArgsForCall[i].serviceInstanceGUID
}

func (fake *FakeV2Actor) GetServiceInstanceSharedTosByServiceInstanceReturns(result1 []v2action.ServiceInstanceSharedTo, result2 v2action.Warnings, result3 error) {
	fake.GetServiceInstanceSharedTosByServiceInstanceStub = nil
	fake.getServiceInstanceSharedTosByServiceInstanceReturns = struct {
		result1 []v2action.ServiceInstanceSharedTo
		result2 v2action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeV2Actor) GetServiceInstanceSharedTosByServiceInstanceReturnsOnCall(i int, result1 []v2action.ServiceInstanceSharedTo, result2 v2action.Warnings, result3 error) {
	fake.GetServiceInstanceSharedTosByServiceInstanceStub = nil
	if fake.getServiceInstanceSharedTosByServiceInstanceReturnsOnCall == nil {
		fake.getServiceInstanceSharedTosByServiceInstanceReturnsOnCall = make(map[int]struct {
			result1 []v2action.ServiceInstanceSharedTo
			result2 v2action.Warnings
			result3 error
		})
	}
	fake.getServiceInstanceSharedTosByServiceInstanceReturnsOnCall[i] = struct {
		result1 []v2action.ServiceInstanceSharedTo
		result2 v2action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeV2Actor) GetSpaceByOrganizationAndName(orgGUID string, spaceName string) (v2action.Space, v2action.Warnings, error) {
	fake.getSpaceByOrganizationAndNameMutex.Lock()
	ret, specificReturn := fake.getSpaceByOrganizationAndNameReturnsOnCall[len(fake.getSpaceByOrganizationAndNameArgsForCall)]
	fake.getSpaceByOrganizationAndNameArgsForCall = append(fake.getSpaceByOrganizationAndNameArgsForCall, struct {
		orgGUID   string
		spaceName string
	}{orgGUID, spaceName})
	fake.recordInvocation("GetSpaceByOrganizationAndName", []interface{}{orgGUID, spaceName})
	fake.getSpaceByOrganizationAndNameMutex.Unlock()
	if fake.GetSpaceByOrganizationAndNameStub != nil {
		return fake.GetSpaceByOrganizationAndNameStub(orgGUID, spaceName)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	return fake.getSpaceByOrganizationAndNameReturns.result1, fake.getSpaceByOrganizationAndNameReturns.result2, fake.getSpaceByOrganizationAndNameReturns.result3
}

func (fake *FakeV2Actor) GetSpaceByOrganizationAndNameCallCount() int {
	fake.getSpaceByOrganizationAndNameMutex.RLock()
	defer fake.getSpaceByOrganizationAndNameMutex.RUnlock()
	return len(fake.getSpaceByOrganizationAndNameArgsForCall)
}

func (fake *FakeV2Actor) GetSpaceByOrganizationAndNameArgsForCall(i int) (string, string) {
	fake.getSpaceByOrganizationAndNameMutex.RLock()
	defer fake.getSpaceByOrganizationAndNameMutex.RUnlock()
	return fake.getSpaceByOrganizationAndNameArgsForCall[i].orgGUID, fake.getSpaceByOrganizationAndNameArgsForCall[i].spaceName
}

func (fake *FakeV2Actor) GetSpaceByOrganizationAndNameReturns(result1 v2action.Space, result2 v2action.Warnings, result3 error) {
	fake.GetSpaceByOrganizationAndNameStub = nil
	fake.getSpaceByOrganizationAndNameReturns = struct {
		result1 v2action.Space
		result2 v2action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeV2Actor) GetSpaceByOrganizationAndNameReturnsOnCall(i int, result1 v2action.Space, result2 v2action.Warnings, result3 error) {
	fake.GetSpaceByOrganizationAndNameStub = nil
	if fake.getSpaceByOrganizationAndNameReturnsOnCall == nil {
		fake.getSpaceByOrganizationAndNameReturnsOnCall = make(map[int]struct {
			result1 v2action.Space
			result2 v2action.Warnings
			result3 error
		})
	}
	fake.getSpaceByOrganizationAndNameReturnsOnCall[i] = struct {
		result1 v2action.Space
		result2 v2action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeV2Actor) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.createApplicationManifestByNameAndSpaceMutex.RLock()
	defer fake.createApplicationManifestByNameAndSpaceMutex.RUnlock()
	fake.getApplicationInstancesWithStatsByApplicationMutex.RLock()
	defer fake.getApplicationInstancesWithStatsByApplicationMutex.RUnlock()
	fake.getApplicationRoutesMutex.RLock()
	defer fake.getApplicationRoutesMutex.RUnlock()
	fake.getFeatureFlagsMutex.RLock()
	defer fake.getFeatureFlagsMutex.RUnlock()
	fake.getServiceMutex.RLock()
	defer fake.getServiceMutex.RUnlock()
	fake.getServiceInstanceByNameAndSpaceMutex.RLock()
	defer fake.getServiceInstanceByNameAndSpaceMutex.RUnlock()
	fake.getServiceInstanceSharedTosByServiceInstanceMutex.RLock()
	defer fake.getServiceInstanceSharedTosByServiceInstanceMutex.RUnlock()
	fake.getSpaceByOrganizationAndNameMutex.RLock()
	defer fake.getSpaceByOrganizationAndNameMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeV2Actor) recordInvocation(key string, args []interface{}) {
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

var _ v2v3action.V2Actor = new(FakeV2Actor)
