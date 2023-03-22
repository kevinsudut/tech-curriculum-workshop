package middleware

import (
	"github.com/kevinsudut/tech-curriculum-workshops/api/response"
	"github.com/kevinsudut/tech-curriculum-workshops/lib/session"
)

type Middleware struct {
	Session  session.SessionItf
	Response response.ResponseItf
}
