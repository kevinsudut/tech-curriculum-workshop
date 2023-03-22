package handler

import (
	"io/ioutil"
	"net/http"
	"strings"

	jsoniter "github.com/json-iterator/go"
	usercontroller "github.com/kevinsudut/tech-curriculum-workshops/app/controller/user"
)

func (h *Handler) Register(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		h.Response.Write(w, r, nil, err)
		return
	}

	var req usercontroller.RegisterRequest

	err = jsoniter.Unmarshal([]byte(data), &req)
	if err != nil {
		h.Response.Write(w, r, nil, err)
		return
	}

	req.Name = strings.TrimSpace(req.Name)
	req.Email = strings.TrimSpace(req.Email)

	resp, err := h.Server.UserController.Register(r.Context(), &req)
	if err != nil {
		h.Response.Write(w, r, nil, err)
		return
	}

	h.Response.Write(w, r, resp, nil)
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		h.Response.Write(w, r, nil, err)
		return
	}

	var req usercontroller.LoginRequest

	err = jsoniter.Unmarshal([]byte(data), &req)
	if err != nil {
		h.Response.Write(w, r, nil, err)
		return
	}

	req.Email = strings.TrimSpace(req.Email)

	resp, err := h.Server.UserController.Login(r.Context(), &req)
	if err != nil {
		h.Response.Write(w, r, nil, err)
		return
	}

	h.Session.Login(w, r, resp.User)

	h.Response.Write(w, r, resp, nil)
}

func (h *Handler) Logout(w http.ResponseWriter, r *http.Request) {
	h.Session.Logout(w, r)

	h.Response.Write(w, r, "Success", nil)
}
