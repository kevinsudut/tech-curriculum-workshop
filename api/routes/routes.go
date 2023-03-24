package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kevinsudut/tech-curriculum-workshops/api/handler"
	"github.com/kevinsudut/tech-curriculum-workshops/api/middleware"
	"github.com/kevinsudut/tech-curriculum-workshops/lib/errors"
	"github.com/kevinsudut/tech-curriculum-workshops/lib/session"
	"github.com/kevinsudut/tech-curriculum-workshops/server"
)

func Init(router *mux.Router) error {
	session := session.Init()

	server, err := server.Init()
	if err != nil {
		return errors.AddTrace(err)
	}

	h := handler.Init(
		server,
		session,
	)

	middleware := middleware.Init(
		session,
	)

	router.HandleFunc("/", h.Index)

	userRoute := router.PathPrefix("/auth").Subrouter()
	userRoute.HandleFunc("/register", h.Register).Methods(http.MethodPost).Schemes("https")
	userRoute.HandleFunc("/login", h.Login).Methods(http.MethodPost).Schemes("https")
	userRoute.HandleFunc("/logout", h.Logout).Methods(http.MethodGet, http.MethodPost).Schemes("https")

	bookRoute := router.PathPrefix("/book").Subrouter()
	bookRoute.HandleFunc("/get-all", middleware.MiddlewareAuth(h.GetAllBook)).Methods(http.MethodGet).Schemes("https")
	bookRoute.HandleFunc("/search/{query}", middleware.MiddlewareAuth(h.SearchBook)).Methods(http.MethodGet).Schemes("https")
	bookRoute.HandleFunc("/{id}", middleware.MiddlewareAuth(h.GetBookByID)).Methods(http.MethodGet).Schemes("https")

	reviewRoute := router.PathPrefix("/review").Subrouter()
	reviewRoute.HandleFunc("/book/{id}", middleware.MiddlewareAuth(h.GetReviewByBook)).Methods(http.MethodGet).Schemes("https")
	reviewRoute.HandleFunc("/post", middleware.MiddlewareAuth(h.PostReview)).Methods(http.MethodPost).Schemes("https")
	reviewRoute.HandleFunc("/delete", middleware.MiddlewareAuth(h.DeleteReview)).Methods(http.MethodPost).Schemes("https")

	return nil
}
