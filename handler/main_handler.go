package handler

import (
	"encoding/json"
	"net/http"
	"vortex-engine/model"
)

// Handler for the /main route
func ProcessMainHandler(responseWriter http.ResponseWriter, request *http.Request) {
	responseWriter.Header().Set("Content-Type", "application/json")
	response := model.ResponseInfo{
		ResponseMessage: "Success",
		ResponseCode:    "SUCCESS",
	}
	json.NewEncoder(responseWriter).Encode(response)
}
