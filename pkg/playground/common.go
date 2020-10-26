package playground

import (
	"azure-playground-generator/internal/auth"
	"azure-playground-generator/internal/config"

	"github.com/Azure/azure-sdk-for-go/profiles/latest/authorization/mgmt/authorization"
	"github.com/Azure/azure-sdk-for-go/profiles/latest/resources/mgmt/resources"
)

func getResourceGroupsClient() (resources.GroupsClient, error) {
	var groupsClient resources.GroupsClient
	var err error

	groupsClient = resources.NewGroupsClient(config.SubscriptionID())
	groupsClient.AddToUserAgent(config.UserAgent())
	authorizer, err := auth.GetResourceManagementAuthorizer()

	if err == nil {
		groupsClient.Authorizer = authorizer
		return groupsClient, nil
	}

	return groupsClient, err
}

func getRoleAssignmentsClient() (authorization.RoleAssignmentsClient, error) {
	roleClient := authorization.NewRoleAssignmentsClient(config.SubscriptionID())

	a, _ := auth.GetResourceManagementAuthorizer()
	roleClient.Authorizer = a
	roleClient.AddToUserAgent(config.UserAgent())
	return roleClient, nil
}
