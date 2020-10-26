package playground

import (
	"azure-playground-generator/pkg/errors"
	"context"
	"net/http"
)

// DeletePlayground deletes a playground with a given name
func DeletePlayground(ctx context.Context, name string) (interface{}, error) {

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

	response, err := groupsClient.Delete(ctx, name)
	if err != nil {
		return nil, err
	}

	return response, nil
}
