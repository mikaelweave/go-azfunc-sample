package api

import (
	"encoding/json"
	"net/http"

	"azure-playground-generator/pkg/errors"
	"azure-playground-generator/pkg/playground"
)

// InvokeRequest is a struct that represents an Azure Function input
type InvokeRequest struct {
	Data     map[string]interface{}
	Metadata map[string]interface{}
}

// InvokeHTTPRequest is a struct that represents an Azure Function input with HTTP input
type InvokeHTTPRequest struct {
	Data     map[string]HTTPInput
	Metadata map[string]interface{}
}

// HTTPInput is a struct that represents the data of a HTTP binding for Azure Functions
type HTTPInput struct {
	Body    string
	Headers map[string][]string
	//Identities string
	Method string
	Params map[string]string
	Query  interface{}
	URL    string
}

func requestDecoder(r *http.Request) (*HTTPInput, error) {

	// Decode Azure Function Request
	var invokeReq InvokeHTTPRequest
	d := json.NewDecoder(r.Body)
	decodeErr := d.Decode(&invokeReq)
	if decodeErr != nil {
		return nil, errors.NewBadRequest("Invalid format from function")
	}

	// Pull out request
	req, ok := invokeReq.Data["req"]
	if !ok {
		return nil, errors.NewBadRequest("Function input req not found or improperly constructed")
	}

	return &req, nil
}

// HTTPTriggerHandler returns test data from server request
func HTTPTriggerHandler(w http.ResponseWriter, r *http.Request) {

	// Decode request object
	req, err := requestDecoder(r)
	if err != nil {
		WriteHTTPErrorResponse(w, err)
		return
	}

	// Get data from playground package
	test, err := playground.Test(r.Context(), req.Body)
	if err != nil {
		WriteHTTPErrorResponse(w, err)
		return
	}

	WriteHTTPResponse(w, http.StatusOK, test)
}
