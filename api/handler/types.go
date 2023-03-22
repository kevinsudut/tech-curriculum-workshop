package handler

import (
	"github.com/kevinsudut/tech-curriculum-workshops/api/response"
	"github.com/kevinsudut/tech-curriculum-workshops/lib/session"
	"github.com/kevinsudut/tech-curriculum-workshops/server"
)

type Handler struct {
	Server   *server.Server
	Session  session.SessionItf
	Response response.ResponseItf
}
