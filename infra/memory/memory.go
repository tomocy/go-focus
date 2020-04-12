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

func (r userRepo) FindByEmail(_ context.Context, email string) (*focus.User, error) {
	for _, u := range r.users {
		if u.Email() == email {
			return u, nil
		}
	}

	return nil, focus.ErrNoSuchUser
}
