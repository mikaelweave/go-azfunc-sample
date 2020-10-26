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

// PlaygroundListHandler returns playgrounds from a Gin server request
func PlaygroundListHandler(w http.ResponseWriter, r *http.Request) {

	playgrounds, err := playground.ListPlaygrounds(r.Context())
	if err != nil {
		WriteHTTPErrorResponse(w, err)
		return
	}

	WriteHTTPResponse(w, http.StatusOK, playgrounds)
}

// PlaygroundCreateHandler creates a playground from a Gin server request
func PlaygroundCreateHandler(w http.ResponseWriter, r *http.Request) {

	// Decode request object
	req, err := requestDecoder(r)
	if err != nil {
		WriteHTTPErrorResponse(w, err)
		return
	}

	// Decode Body
	var playgroundReq playground.Playground
	err = json.Unmarshal([]byte(req.Body), &playgroundReq)
	if err != nil {
		WriteHTTPErrorResponse(w, errors.NewBadRequest("invalid input"))
		return
	}

	// Test inputs
	if playgroundReq.Name == nil || playgroundReq.Location == nil || playgroundReq.OwnerID == nil {
		WriteHTTPErrorResponse(w, errors.NewBadRequest("Please provide a name, location, and ownerId for this request"))
		return
	}

	// Create playground
	resp, err := playground.CreatePlayground(r.Context(), *playgroundReq.Name, *playgroundReq.Location, *playgroundReq.OwnerID)
	if err != nil {
		WriteHTTPErrorResponse(w, err)
		return
	}

	WriteHTTPResponse(w, http.StatusCreated, resp)
}

// PlaygroundGetHandler gets a specific playground (given the name) from a Gin server request
func PlaygroundGetHandler(w http.ResponseWriter, r *http.Request) {
	// Decode request object
	req, err := requestDecoder(r)
	if err != nil {
		WriteHTTPErrorResponse(w, err)
		return
	}

	// Get name parameter
	name, ok := req.Params["name"]
	if !ok {
		WriteHTTPErrorResponse(w, errors.NewBadRequest("Get Playground requires the name URL parameter"))
		return
	}

	// Get and return playground
	playgroundRtn, err := playground.GetPlayground(r.Context(), name)
	if err != nil {
		WriteHTTPErrorResponse(w, err)
		return
	}

	WriteHTTPResponse(w, http.StatusOK, playgroundRtn)
}

// PlaygroundDeleteHandler deletes a specific playground (given the name) from a Gin server request
func PlaygroundDeleteHandler(w http.ResponseWriter, r *http.Request) {

	// Decode request object
	req, err := requestDecoder(r)
	if err != nil {
		WriteHTTPErrorResponse(w, err)
		return
	}

	// Get name parameter
	name, ok := req.Params["name"]
	if !ok {
		WriteHTTPErrorResponse(w, errors.NewBadRequest("Delete Playground requires the name URL parameter"))
		return
	}

	// Delete Playground
	playgroundRtn, err := playground.DeletePlayground(r.Context(), name)
	if err != nil {
		WriteHTTPErrorResponse(w, err)
		return
	}

	WriteHTTPResponse(w, http.StatusAccepted, playgroundRtn)
}
