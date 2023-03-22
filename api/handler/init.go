package handler

import (
	"github.com/kevinsudut/tech-curriculum-workshops/api/response"
	"github.com/kevinsudut/tech-curriculum-workshops/lib/session"
	"github.com/kevinsudut/tech-curriculum-workshops/server"
)

func Init(
	server *server.Server,
	session session.SessionItf,
) *Handler {
	return &Handler{
		Server:   server,
		Session:  session,
		Response: &response.Response{},
	}
}
