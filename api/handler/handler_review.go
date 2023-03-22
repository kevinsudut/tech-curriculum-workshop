package handler

import (
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	jsoniter "github.com/json-iterator/go"
	reviewcontroller "github.com/kevinsudut/tech-curriculum-workshops/app/controller/review"
	"github.com/kevinsudut/tech-curriculum-workshops/util"
)

func (h *Handler) GetReviewByBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	req := reviewcontroller.GetReviewByBookRequest{
		BookID: util.Atoi64(vars["id"]),
	}

	resp, err := h.Server.ReviewController.GetReviewByBook(r.Context(), &req)
	if err != nil {
		h.Response.Write(w, r, nil, err)
		return
	}

	h.Response.Write(w, r, resp, nil)
}

func (h *Handler) PostReview(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		h.Response.Write(w, r, nil, err)
		return
	}

	var req reviewcontroller.PostReviewRequest

	err = jsoniter.Unmarshal([]byte(data), &req)
	if err != nil {
		h.Response.Write(w, r, nil, err)
		return
	}

	req.Content = strings.TrimSpace(req.Content)

	resp, err := h.Server.ReviewController.PostReview(r.Context(), &req)
	if err != nil {
		h.Response.Write(w, r, nil, err)
		return
	}

	h.Response.Write(w, r, resp, nil)
}

func (h *Handler) DeleteReview(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		h.Response.Write(w, r, nil, err)
		return
	}

	var req reviewcontroller.DeleteReviewRequest

	err = jsoniter.Unmarshal([]byte(data), &req)
	if err != nil {
		h.Response.Write(w, r, nil, err)
		return
	}

	resp, err := h.Server.ReviewController.DeleteReview(r.Context(), &req)
	if err != nil {
		h.Response.Write(w, r, nil, err)
		return
	}

	h.Response.Write(w, r, resp, nil)
}
