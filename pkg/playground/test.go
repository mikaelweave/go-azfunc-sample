package playground

import (
	"context"
)

// Test returns a Test string
func Test(ctx context.Context, reqBody interface{}) (interface{}, error) {
	// Build Response
	resp := make(map[string]interface{})
	resp["requestBody"] = reqBody
	resp["value"] = "test return value"

	return resp, nil
}
