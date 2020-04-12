package memory

import (
	"context"

	"github.com/tomocy/focus"
	"github.com/tomocy/focus/infra/rand"
)

type userRepo struct {
	users map[focus.UserID]*focus.User
}

func (r userRepo) NextID(context.Context) (focus.UserID, error) {
	return focus.UserID(rand.GenerateString(30)), nil
}
