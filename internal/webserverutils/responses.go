package webserverutils

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type ErrorResponseBody struct {
	Error string `json:"error"`
}

func RespondWithJsonError(w http.ResponseWriter, message string, code int) {
	w.Header().Add("Content-Type", "applciation/json")
	w.WriteHeader(code)
	respBody := ErrorResponseBody{
		Error: message,
	}
	respBytes, _ := json.Marshal(respBody)
	w.Write(respBytes)
}

type RequestError struct {
	Message string
}

func (re RequestError) Error() string {
	return fmt.Sprintf("Invalid Request Body: %s", re.Message)
}

func NewRequestError(message string) error {
	return RequestError{message}
}
