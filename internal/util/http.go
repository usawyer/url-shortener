package util

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
)

func InternalServerError(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
}

func BadRequest(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusBadRequest)
}

func JSON(w http.ResponseWriter, r *http.Request, statusCode int, obj any) {
	w.WriteHeader(statusCode)
	w.Header().Set("content-type", "application/json")
	jsonBytes, _ := json.Marshal(obj)
	w.Write(jsonBytes)
}

func BindJSON(w http.ResponseWriter, r *http.Request, object any) bool {
	if err := json.NewDecoder(r.Body).Decode(object); err != nil {
		log.Printf("ERROR: failed to decode request body err=%s\n", err.Error())

		var sErr *json.SyntaxError
		if errors.As(err, &sErr) {
			BadRequest(w, r)
			w.Header().Set("content-type", "text/plain")
			w.Write([]byte("invalid json"))
			return false
		}

		if errors.Is(err, io.EOF) {
			BadRequest(w, r)
			w.Header().Set("content-type", "text/plain")
			w.Write([]byte("empty request"))
			return false
		}

		var utErr *json.UnmarshalTypeError
		if errors.As(err, &utErr) {
			BadRequest(w, r)
			w.Header().Set("content-type", "text/plain")
			w.Write([]byte("incorrect request typing"))
			return false
		}

		InternalServerError(w, r)
		return false
	}

	return true
}
