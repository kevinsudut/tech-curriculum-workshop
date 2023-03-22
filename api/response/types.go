package response

import (
	"net/http"
)

type Response struct {
	Error string      `json:"error,omitempty"`
	Data  interface{} `json:"data,omitempty"`
}

type ResponseItf interface {
	Write(w http.ResponseWriter, r *http.Request, data interface{}, err error)
}
