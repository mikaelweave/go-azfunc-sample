package auth

import (
	"azure-playground-generator/internal/config"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/adal"
)

// Source:
// https://github.com/Azure-Samples/azure-sdk-for-go-samples/blob/13286d0626c4115990db2d86631e954c566df82e/authorization/authorization.go

var (
	armAuthorizer autorest.Authorizer
)

// GetResourceManagementAuthorizer gets an OAuthTokenAuthorizer for Azure Resource Manager
func GetResourceManagementAuthorizer() (autorest.Authorizer, error) {
	if armAuthorizer != nil {
		return armAuthorizer, nil
	}

	var a autorest.Authorizer
	var err error

	a, err = getAuthorizerForResource(config.Environment().ResourceManagerEndpoint)

	if err == nil {
		// cache
		armAuthorizer = a
	} else {
		// clear cache
		armAuthorizer = nil
	}
	return armAuthorizer, err
}

func getAuthorizerForResource(resource string) (autorest.Authorizer, error) {
	var a autorest.Authorizer
	var err error

	oauthConfig, err := adal.NewOAuthConfig(
		config.Environment().ActiveDirectoryEndpoint, config.TenantID())
	if err != nil {
		return nil, err
	}

	token, err := adal.NewServicePrincipalToken(
		*oauthConfig, config.ClientID(), config.ClientSecret(), resource)
	if err != nil {
		return nil, err
	}
	a = autorest.NewBearerAuthorizer(token)

	return a, err
}
