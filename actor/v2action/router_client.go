package v2action

import "code.cloudfoundry.org/cli/api/router"

//go:generate counterfeiter . RouterClient

// RouterClient is a Router API client.
type RouterClient interface {
	GetRouterGroups() []router.RouterGroup

	API() string
	APIVersion() string
	RoutingEndpoint() string
	TokenEndpoint() string
}
