package playground

import (
	"azure-playground-generator/internal/config"
	"azure-playground-generator/pkg/errors"
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/azure-sdk-for-go/profiles/latest/authorization/mgmt/authorization"
	"github.com/Azure/azure-sdk-for-go/profiles/latest/resources/mgmt/resources"
	"github.com/Azure/go-autorest/autorest/to"
	uuid "github.com/satori/go.uuid"
)

// CreatePlayground creates a specific resource group with playground attributes
func CreatePlayground(ctx context.Context, name string, location string, ownerPrincipalID string) (*Playground, error) {

	// Get Client
	groupsClient, err := getResourceGroupsClient()
	if err != nil {
		return nil, err
	}

	// See if group already exists (throw error if so)
	existResp, err := groupsClient.CheckExistence(ctx, name)
	if err != nil {
		return nil, err
	}
	if existResp.StatusCode != http.StatusNotFound {
		return nil, errors.NewAlreadyExists(name)
	}

	var parameters resources.Group
	parameters.Location = to.StringPtr(location)
	parameters.Tags = make(map[string]*string)
	parameters.Tags["System"] = to.StringPtr("Playground")
	parameters.Tags["OwnerId"] = to.StringPtr(ownerPrincipalID)

	group, err := groupsClient.CreateOrUpdate(ctx, name, parameters)
	if err != nil {
		return nil, err
	}

	// Assign Contributor Role
	err = assignPlaygroundRole(ctx, *group.ID, ownerPrincipalID, "b24988ac-6180-42a0-ab88-20f7382dd24c")
	if err != nil {
		return nil, err
	}
	// Assign User Access Administrator Role
	err = assignPlaygroundRole(ctx, *group.ID, ownerPrincipalID, "18d7d88d-d35e-4fb5-a5c3-7773c20a72d9")
	if err != nil {
		return nil, err
	}

	playground := groupToPlayground(group)
	return &playground, nil
}

func assignPlaygroundRole(ctx context.Context, scope string, ownerPrincipalID string, roleID string) (err error) {

	roleAssignmentClient, err := getRoleAssignmentsClient()
	if err != nil {
		return err
	}

	// Assign Contributor Role
	roleDefID := fmt.Sprintf("/subscriptions/%s/providers/Microsoft.Authorization/roleDefinitions/%s", config.SubscriptionID(), roleID)
	_, err = roleAssignmentClient.Create(
		ctx,
		scope,
		uuid.NewV1().String(),
		authorization.RoleAssignmentCreateParameters{
			Properties: &authorization.RoleAssignmentProperties{
				PrincipalID:      to.StringPtr(ownerPrincipalID),
				RoleDefinitionID: to.StringPtr(roleDefID),
			},
		},
	)

	if err != nil {
		return err
	}

	return nil
}
