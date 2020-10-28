package config

// Source:
// https://github.com/Azure-Samples/azure-sdk-for-go-samples/blob/3d51ac8a1a5097b8881a8cf29888d4a44f7205f5/internal/config/config.go

var (
	functionHTTPWorkerPort int = 8082
	subscriptionID         string
	userAgent              string
)

// FunctionHTTPWorkerPort is the port the go server will run on and Azure Function will send requests to
func FunctionHTTPWorkerPort() int {
	return functionHTTPWorkerPort
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
