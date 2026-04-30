package handler

import (
	"encoding/json"
	"net/http"
	"vortex-engine/external/quote"
	"vortex-engine/logger"
	"vortex-engine/model"
)

// Handler for the /main route
func ProcessMainHandler(responseWriter http.ResponseWriter, request *http.Request) {
	logger.Info.Println("Main handler invoked.")
	logger.Info.Println("Origin: " + request.Header.Get("Origin"))

	responseWriter.Header().Set("Content-Type", "application/json")

	quote := quote.GetZenQuote()

	response := model.ResponseInfo{
		ResponseMessage: quote.Text,
		ResponseCode:    "SUCCESS",
	}

	json.NewEncoder(responseWriter).Encode(response)
}
