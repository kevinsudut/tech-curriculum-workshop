package session

import (
	"encoding/gob"
	"fmt"
	"net/http"

	"github.com/gorilla/sessions"
	usermodel "github.com/kevinsudut/tech-curriculum-workshops/app/model/user"
	"github.com/kevinsudut/tech-curriculum-workshops/config"
	"github.com/kevinsudut/tech-curriculum-workshops/lib/errors"
)

type Session struct {
	store *sessions.CookieStore
}

type SessionItf interface {
	CheckSession(w http.ResponseWriter, r *http.Request) (*usermodel.User, error)
	Login(w http.ResponseWriter, r *http.Request, user usermodel.User)
	Logout(w http.ResponseWriter, r *http.Request)
}

func Init() *Session {
	gob.Register(&usermodel.User{})

	return &Session{
		store: sessions.NewCookieStore([]byte(config.SESSION_ID)),
	}
}

func (s *Session) CheckSession(w http.ResponseWriter, r *http.Request) (*usermodel.User, error) {
	session, err := s.store.Get(r, config.SESSION_NAME)
	if err != nil {
		fmt.Println("ERR check session", err)
	}

	auth, ok := session.Values[config.SESSION_AUTHENTICATION].(*usermodel.User)
	if !ok || auth.ID <= 0 {
		return auth, errors.AddTrace(errors.ForbiddenRequest)
	}

	return auth, nil
}

func (s *Session) Login(w http.ResponseWriter, r *http.Request, user usermodel.User) {
	session, err := s.store.Get(r, config.SESSION_NAME)
	if err != nil {
		fmt.Println("ERR login", err)
	}
	s.store.MaxAge(config.SESSION_DURATION)

	session.Values[config.SESSION_AUTHENTICATION] = user
	err = session.Save(r, w)
	if err != nil {
		fmt.Println("ERR login", err)
	}
}

func (s *Session) Logout(w http.ResponseWriter, r *http.Request) {
	session, err := s.store.Get(r, config.SESSION_NAME)
	if err != nil {
		fmt.Println("ERR logout", err)
	}

	session.Values[config.SESSION_AUTHENTICATION] = usermodel.User{}
	session.Save(r, w)
}
