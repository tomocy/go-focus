package memory

import (
	"context"

	"github.com/tomocy/focus"
)

func NewSessionRepo() *sessionRepo {
	return new(sessionRepo)
}

type sessionRepo struct {
	sess *focus.Session
}

func (r sessionRepo) Pull(context.Context) (*focus.Session, error) {
	if r.sess == nil {
		return nil, focus.ErrNoSession
	}

	return r.sess, nil
}

func (r *sessionRepo) Push(_ context.Context, s *focus.Session) error {
	r.sess = s

	return nil
}

func (r *sessionRepo) Delete(context.Context) error {
	r.sess = nil

	return nil
}
