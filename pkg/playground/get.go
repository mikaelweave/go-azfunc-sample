package playground

import (
	"azure-playground-generator/pkg/errors"
	"context"
	"net/http"
)

// GetPlayground returns a playground with a given name
func GetPlayground(ctx context.Context, name string) (*Playground, error) {

	groupsClient, err := getResourceGroupsClient()
	if err != nil {
		return nil, err
	}

	// See if playground exists (throw error if not)
	existResp, err := groupsClient.CheckExistence(ctx, name)
	if err != nil {
		return nil, err
	}
	if existResp.StatusCode == http.StatusNotFound {
		return nil, errors.NewNotFound(name)
	}

	group, err := groupsClient.Get(ctx, name)
	if err != nil {
		return nil, err
	}

	playground := groupToPlayground(group)

	return &playground, nil
}
