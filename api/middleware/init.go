package middleware

import (
	"github.com/kevinsudut/tech-curriculum-workshops/api/response"
	"github.com/kevinsudut/tech-curriculum-workshops/lib/session"
)

func Init(
	session session.SessionItf,
) *Middleware {
	return &Middleware{
		Session:  session,
		Response: &response.Response{},
	}
}
