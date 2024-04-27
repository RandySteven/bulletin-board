package utils

import (
	"encoding/json"
	"log"
	"net/http"
	"task_mission/entities/dtos/responses"
)

func ContentType(w http.ResponseWriter, contentType string) {
	w.Header().Set("Content-Type", contentType)
}

func BindJSON(req *http.Request, request any) error {
	return json.NewDecoder(req.Body).Decode(&request)
}

func BindForm(req *http.Request, request any) error {
	return req.ParseForm()
}

func ResponseHandler(w http.ResponseWriter, responseCode int, message string, dataKey *string, responseData any, err error) {
	w.WriteHeader(responseCode)
	responseMap := make(map[string]any)
	if dataKey != nil && responseData != nil {
		responseMap[*dataKey] = responseData
	}
	response := responses.NewResponse(message, responseMap, err)
	log.Print(response)
	err = json.NewEncoder(w).Encode(&response)
	if err != nil {
		log.Fatal(err)
		return
	}
}
