package focus

import "context"

const (
	ErrNoSession = err("no session")
)

type SessionRepo interface {
	Pull(context.Context) (*Session, error)
	Push(context.Context, *Session) error
	Delete(context.Context) error
}

type Session struct {
	userID UserID
}
