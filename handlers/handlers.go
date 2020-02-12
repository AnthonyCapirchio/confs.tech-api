package handlers

import (
	"encoding/json"
	"net/http"
)

func sendError(w http.ResponseWriter, code int, message string) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write([]byte("{\"error\": \"" + message + "\"}"))
}

func sendCollection(w http.ResponseWriter, collection interface{}) {
	w.Header().Add("Content-Type", "application/json")

	response, err := json.Marshal(collection)
	if err != nil {
		sendError(w, http.StatusInternalServerError, err.Error())
	} else {
		w.Write(response)
	}

}
