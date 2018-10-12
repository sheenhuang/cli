package cfnetworkingaction

import "code.cloudfoundry.org/cli/actor/v7action"

//go:generate counterfeiter . V3Actor
type V3Actor interface {
	GetApplicationByNameAndSpace(appName string, spaceGUID string) (v7action.Application, v7action.Warnings, error)
	GetApplicationsBySpace(spaceGUID string) ([]v7action.Application, v7action.Warnings, error)
}
