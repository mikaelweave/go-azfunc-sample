package playground

import (
	"azure-playground-generator/internal/config"

	"github.com/Azure/azure-sdk-for-go/profiles/latest/authorization/mgmt/authorization"
	"github.com/Azure/azure-sdk-for-go/profiles/latest/resources/mgmt/resources"
	"github.com/Azure/go-autorest/autorest/azure/auth"
)

func getResourceGroupsClient() (resources.GroupsClient, error) {
	var groupsClient resources.GroupsClient
	var err error

	groupsClient = resources.NewGroupsClient(config.SubscriptionID())
	groupsClient.AddToUserAgent(config.UserAgent())
	authorizer, err := auth.NewAuthorizerFromEnvironment()

	if err == nil {
		groupsClient.Authorizer = authorizer
		return groupsClient, nil
	}

	return groupsClient, err
}

func getRoleAssignmentsClient() (authorization.RoleAssignmentsClient, error) {
	var roleClient authorization.RoleAssignmentsClient
	var err error

	roleClient = authorization.NewRoleAssignmentsClient(config.SubscriptionID())
	roleClient.AddToUserAgent(config.UserAgent())
	authorizer, err := auth.NewAuthorizerFromEnvironment()

	if err == nil {
		roleClient.Authorizer = authorizer
		return roleClient, nil
	}

	return roleClient, err
}
