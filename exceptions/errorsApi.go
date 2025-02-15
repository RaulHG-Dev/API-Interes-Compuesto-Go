package exceptions

import (
	"encoding/json"
	"net/http"
)

type error struct {
	Success   bool   `json:"success"`
	CodeError uint16 `json:"code"`
	Message   string `json:"message"`
}

func NotFound(w http.ResponseWriter, r *http.Request) {
	code := http.StatusNotFound
	notFoundError := error{
		CodeError: uint16(code),
		Message:   http.StatusText(code),
	}

	jsonData, _ := json.Marshal(notFoundError)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(jsonData)
}

func NotAllowedHandler(w http.ResponseWriter, r *http.Request) {
	code := http.StatusMethodNotAllowed
	notFoundError := error{
		CodeError: uint16(code),
		Message:   http.StatusText(code),
	}

	jsonData, _ := json.Marshal(notFoundError)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(jsonData)
}
