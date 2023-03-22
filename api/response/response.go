package response

import (
	"net/http"
)

func (resp *Response) Write(w http.ResponseWriter, r *http.Request, data interface{}, err error) {
	w.Header().Set(CONTENT_TYPE, JSON_CONTENT_TYPE)
	w.WriteHeader(http.StatusOK)
	w.Write(buildResponse(data, err))
}
