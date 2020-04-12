package focus

import (
	"context"
	"fmt"
)

const (
	ErrNoSession = err("no session")
)

type SessionRepo interface {
	Pull(context.Context) (*Session, error)
	Push(context.Context, *Session) error
	Delete(context.Context) error
}

func NewSession(userID UserID) (*Session, error) {
	s := new(Session)

	if err := s.setUserID(userID); err != nil {
		return nil, err
	}

	return s, nil
}

type Session struct {
	userID UserID
}

func (s *Session) setUserID(id UserID) error {
	if id == "" {
		return fmt.Errorf("empty user id")
	}

	s.userID = id

	return nil
}
