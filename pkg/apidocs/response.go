package apidocs

import "net/http"

type APIError struct {
	Msg string `json:"msg"`
}

func Err(msg string) (int, *APIError) {
	return http.StatusInternalServerError, &APIError{
		Msg: msg,
	}
}

func Success(data interface{}) (int, interface{}) {
	return http.StatusOK, data
}
