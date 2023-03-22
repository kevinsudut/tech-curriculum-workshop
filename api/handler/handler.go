package handler

import (
	"net/http"
)

func (h *Handler) Index(w http.ResponseWriter, r *http.Request) {
	h.Response.Write(w, r, "200 OK", nil)
}
