package memory

import (
	"context"

	"github.com/tomocy/focus"
	"github.com/tomocy/focus/infra/rand"
)

func NewUserRepo() *userRepo {
	return &userRepo{
		users: make(map[focus.UserID]*focus.User),
	}
}

type userRepo struct {
	users map[focus.UserID]*focus.User
}

func (r userRepo) NextID(context.Context) (focus.UserID, error) {
	return focus.UserID(rand.GenerateString(30)), nil
}

func (r userRepo) Find(_ context.Context, id focus.UserID) (*focus.User, error) {
	u, ok := r.users[id]
	if !ok {
		return nil, focus.ErrNoSuchUser
	}

	return u, nil
}

func (r userRepo) FindByEmail(_ context.Context, email string) (*focus.User, error) {
	for _, u := range r.users {
		if u.Email() == email {
			return u, nil
		}
	}

	return nil, focus.ErrNoSuchUser
}

func (r *userRepo) Save(_ context.Context, u *focus.User) error {
	r.users[u.ID()] = u

	return nil
}

func (r *userRepo) Delete(_ context.Context, id focus.UserID) error {
	delete(r.users, id)

	return nil
}
