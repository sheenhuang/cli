package v2v3action

import (
	"code.cloudfoundry.org/cli/actor/v7action"
)

//go:generate counterfeiter . V3Actor

type V3Actor interface {
	ManifestV3Actor
	GetApplicationSummaryByNameAndSpace(appName string, spaceGUID string, withObfuscatedValues bool) (v7action.ApplicationSummary, v7action.Warnings, error)
	GetOrganizationByName(orgName string) (v7action.Organization, v7action.Warnings, error)
	ShareServiceInstanceToSpaces(serviceInstanceGUID string, spaceGUIDs []string) (v7action.RelationshipList, v7action.Warnings, error)
	UnshareServiceInstanceByServiceInstanceAndSpace(serviceInstanceGUID string, spaceGUID string) (v7action.Warnings, error)

	CloudControllerAPIVersion() string
}
