package util

import (
	"encoding/json"
	"net/http"
)

func JSON(w http.ResponseWriter, r *http.Request, statusCode int, obj any) {
	w.WriteHeader(statusCode)
	w.Header().Set("content-type", "application/json")
	jsonBytes, _ := json.Marshal(obj)
	w.Write(jsonBytes)
}
