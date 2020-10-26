package config

// Source:
// https://github.com/Azure-Samples/azure-sdk-for-go-samples/blob/4e95ecd68b1c0969dc8e7b6bf3a3b2f7e5ecdc76/internal/config/env.go

import (
	"fmt"
	"strconv"

	"github.com/gobuffalo/envy"
)

// ParseEnvironment loads a sibling `.env` file then looks through all environment
// variables to set global configuration.
func ParseEnvironment() error {
	envy.Load()

	var err error

	// these must be provided by environment
	//functionHTTPWorkerPort
	functionHTTPWorkerPortString, err := envy.MustGet("FUNCTIONS_HTTPWORKER_PORT")
	if err != nil {
		return fmt.Errorf("Expected env vars not provided: %s", err)
	}
	functionHTTPWorkerPort, err = strconv.Atoi(functionHTTPWorkerPortString)
	if err != nil {
		return fmt.Errorf("Expected env vars must be an integer: %s", err)
	}

	// clientID
	clientID, err = envy.MustGet("AZURE_CLIENT_ID")
	if err != nil {
		return fmt.Errorf("Expected env vars not provided: %s", err)
	}

	// clientSecret
	clientSecret, err = envy.MustGet("AZURE_CLIENT_SECRET")
	if err != nil {
		return fmt.Errorf("Expected env vars not provided: %s", err)
	}

	// tenantID (AAD)
	tenantID, err = envy.MustGet("AZURE_TENANT_ID")
	if err != nil {
		return fmt.Errorf("Expected env vars not provided: %s", err)
	}

	// subscriptionID (ARM)
	subscriptionID, err = envy.MustGet("AZURE_SUBSCRIPTION_ID")
	if err != nil {
		return fmt.Errorf("Expected env vars not provided: %s", err)
	}

	return nil
}
