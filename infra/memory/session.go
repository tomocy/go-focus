package memory

import (
	"context"

	"github.com/tomocy/focus"
)

type sessionRepo struct {
	sess *focus.Session
}

func (r sessionRepo) Pull(context.Context) (*focus.Session, error) {
	if r.sess == nil {
		return nil, focus.ErrNoSession
	}

	return r.sess, nil
}
