package api

import (
	"encoding/json"
	"net/http"

	"azure-playground-generator/pkg/errors"
)

// Source:
// https://github.com/Azure-Samples/functions-custom-handlers/blob/3bd2a534130992af6f8af6608a3dc6007fd31161/go/GoCustomHandlers.go
// https://github.com/Optum/dce/blob/790404bd0b336994d627add7d3a176d90d4d6156/pkg/api/error.go

// InvokeResponse is an Azure Function construct showing how the function expects a return from the Go server
type InvokeResponse struct {
	Outputs     map[string]interface{}
	Logs        []string
	ReturnValue interface{}
}

// HTTPBindingResponse is a output of InvokeResponse used to pass data back to a HTTP Output for an Azure Function
type HTTPBindingResponse struct {
	statusCode int
	body       interface{}
	headers    map[string]string
}

// WriteHTTPResponse takes data and writes it in a standard way to a http.ResponseWriter
func WriteHTTPResponse(w http.ResponseWriter, status int, body interface{}) {
	outputs := make(map[string]interface{})
	headers := make(map[string]string)
	headers["System"] = "Azure Playground Generator"
	headers["Content-Type"] = "application/json"

	// Serialize body
	jsonData, err := json.Marshal(body)
	if err != nil {
		WriteHTTPErrorResponse(w, errors.NewInternalServer("error serializing body", err))
		return
	}

	// Create response object
	res := make(map[string]interface{})
	res["statusCode"] = status
	res["body"] = string(jsonData)
	res["headers"] = headers
	outputs["res"] = res
	invokeResponse := InvokeResponse{Outputs: outputs}
	// Serialize response object
	j, err := json.Marshal(invokeResponse)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(j)
}

// WriteHTTPErrorResponse takes an error and writes it in a standard way to a http.ResponseWriter
func WriteHTTPErrorResponse(w http.ResponseWriter, err error) {
	// If custom error object
	switch t := err.(type) {
	case errors.HTTPCode:
		WriteHTTPResponse(w, t.HTTPCode(), err.Error())
		return
	}

	// If standard error
	WriteHTTPResponse(
		w,
		http.StatusInternalServerError,
		err.Error(),
	)
}
