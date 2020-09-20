// Package config manages loading configuration from environment
package config

// Source:
// https://github.com/Azure-Samples/azure-sdk-for-go-samples/blob/3d51ac8a1a5097b8881a8cf29888d4a44f7205f5/internal/config/config.go

var (
	functionHTTPWorkerPort int = 8082
)

// FunctionHTTPWorkerPort is the port the go server will run on and Azure Function will send requests to
func FunctionHTTPWorkerPort() int {
	return functionHTTPWorkerPort
}
