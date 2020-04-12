package usecase

import (
	"context"
	"fmt"

	"github.com/tomocy/focus"
)

type registerUser struct {
	userRepo focus.UserRepo
	sessRepo focus.SessionRepo
}

func (u *registerUser) Do(email, pass string) (*focus.User, error) {
	ctx := context.TODO()

	id, err := u.userRepo.NextID(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to generate user id: %w", err)
	}
	hashed, err := focus.HashPassword(pass)
	if err != nil {
		return nil, fmt.Errorf("failed to hash password: %w", err)
	}
	user, err := focus.NewUser(id, email, hashed)
	if err != nil {
		return nil, err
	}

	if err := u.userRepo.Save(ctx, user); err != nil {
		return nil, fmt.Errorf("failed to save user: %s", err)
	}

	sess, err := focus.NewSession(user.ID())
	if err != nil {
		return nil, fmt.Errorf("failed to generate session: %w", err)
	}

	if err := u.sessRepo.Push(ctx, sess); err != nil {
		return nil, fmt.Errorf("failed to push session: %w", err)
	}

	return user, nil
}

type authenticateUser struct{}
