package middleware

import (
	"context"
	"net/http"

	"github.com/kevinsudut/tech-curriculum-workshops/config"
)

func (m *Middleware) MiddlewareAuth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		auth, err := m.Session.CheckSession(w, r)
		if err != nil {
			m.Response.Write(w, r, nil, err)
			return
		}

		ctx := context.WithValue(r.Context(), config.SESSION_AUTHENTICATION, auth)

		next.ServeHTTP(w, r.WithContext(ctx))
	}
}
