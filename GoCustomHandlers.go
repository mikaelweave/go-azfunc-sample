package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

type InvokeResponse struct {
	Outputs     map[string]interface{}
	Logs        []string
	ReturnValue interface{}
}

func httpTriggerHandler(w http.ResponseWriter, r *http.Request) {
	headers := make(map[string]interface{})
	headers["Content-Type"] = "application/json"

	res := make(map[string]interface{})
	res["statusCode"] = http.StatusOK
	res["headers"] = headers
	res["body"] = "{\"value\": \"test return value\"}"

	outputs := make(map[string]interface{})
	outputs["res"] = res
	invokeResponse := InvokeResponse{outputs, nil, nil}

	js, err := json.Marshal(invokeResponse)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func main() {
	customHandlerPort, exists := os.LookupEnv("FUNCTIONS_CUSTOMHANDLER_PORT")
	if exists {
		fmt.Println("FUNCTIONS_CUSTOMHANDLER_PORT: " + customHandlerPort)
	}
	mux := http.NewServeMux()
	mux.HandleFunc("/HttpTriggerWithOutputs", httpTriggerHandler)
	fmt.Println("Go server Listening...on FUNCTIONS_CUSTOMHANDLER_PORT:", customHandlerPort)
	log.Fatal(http.ListenAndServe(":"+customHandlerPort, mux))
}
