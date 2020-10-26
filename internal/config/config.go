package config

import (
	"fmt"

	"github.com/Azure/go-autorest/autorest/azure"
)

// Source:
// https://github.com/Azure-Samples/azure-sdk-for-go-samples/blob/3d51ac8a1a5097b8881a8cf29888d4a44f7205f5/internal/config/config.go

var (
	functionHTTPWorkerPort int = 8082
	clientID               string
	clientSecret           string
	tenantID               string
	subscriptionID         string
	cloudName              string = "AzurePublicCloud"
	userAgent              string
	environment            *azure.Environment
)

// FunctionHTTPWorkerPort is the port the go server will run on and Azure Function will send requests to
func FunctionHTTPWorkerPort() int {
	return functionHTTPWorkerPort
}

// ClientID is the OAuth client ID.
func ClientID() string {
	return clientID
}

// ClientSecret is the OAuth client secret.
func ClientSecret() string {
	return clientSecret
}

// TenantID is the AAD tenant to which this client belongs.
func TenantID() string {
	return tenantID
}

// SubscriptionID is a target subscription for Azure resources.
func SubscriptionID() string {
	return subscriptionID
}

// UserAgent specifies a string to append to the agent identifier.
func UserAgent() string {
	if len(userAgent) > 0 {
		return userAgent
	}
	return "playground-manager"
}

// Environment returns an `azure.Environment{...}` for the current cloud.
func Environment() *azure.Environment {
	if environment != nil {
		return environment
	}
	env, err := azure.EnvironmentFromName(cloudName)
	if err != nil {
		// TODO: move to initialization of var
		panic(fmt.Sprintf(
			"invalid cloud name '%s' specified, cannot continue\n", cloudName))
	}
	environment = &env
	return environment
}
