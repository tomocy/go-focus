package focus

import "context"

type SessionRepo interface {
	Pull(context.Context) (*Session, error)
}

type Session struct {
	userID UserID
}
