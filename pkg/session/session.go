package session

import (
	"net/http"
	"time"

	"github.com/golangcollege/sessions"
)

type Session struct {
	*sessions.Session
}

func New(secret *string) *Session {
	session := &Session{sessions.New([]byte(*secret))}
	session.setupSession()
	return session
}

func (s *Session) setupSession() {
	s.Lifetime = 5 * time.Minute
	s.Secure = true
	s.SameSite = http.SameSiteStrictMode
}
