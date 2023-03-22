package handler

import (
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	bookcontroller "github.com/kevinsudut/tech-curriculum-workshops/app/controller/book"
	"github.com/kevinsudut/tech-curriculum-workshops/util"
)

func (h *Handler) GetAllBook(w http.ResponseWriter, r *http.Request) {
	resp, err := h.Server.BookController.GetAllBook(r.Context(), &bookcontroller.GetAllBookRequest{})
	if err != nil {
		h.Response.Write(w, r, nil, err)
		return
	}

	h.Response.Write(w, r, resp, nil)
}

func (h *Handler) GetBookByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	req := bookcontroller.GetBookByIDRequest{
		ID: util.Atoi64(vars["id"]),
	}

	resp, err := h.Server.BookController.GetBookByID(r.Context(), &req)
	if err != nil {
		h.Response.Write(w, r, nil, err)
		return
	}

	h.Response.Write(w, r, resp, nil)
}

func (h *Handler) SearchBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	req := bookcontroller.SearchBookRequest{
		Query: strings.TrimSpace(vars["query"]),
	}

	resp, err := h.Server.BookController.SearchBook(r.Context(), &req)
	if err != nil {
		h.Response.Write(w, r, nil, err)
		return
	}

	h.Response.Write(w, r, resp, nil)
}
