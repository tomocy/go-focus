package usecase

import (
	"context"
	"fmt"
	"testing"

	"github.com/tomocy/focus"
	"github.com/tomocy/focus/infra/memory"
)

func TestRegisterUser(t *testing.T) {
	userRepo, sessRepo := memory.NewUserRepo(), memory.NewSessionRepo()

	u := registerUser{
		userRepo: userRepo,
		sessRepo: sessRepo,
	}

	email, pass := "email", "pass"
	user, err := u.Do(email, pass)
	if err != nil {
		t.Errorf("should have registered user: %s", err)
		return
	}

	if err := assertUser(
		user,
		checkIfIDOfUserIsFilled(), checkIfEmailOfUserIsSame(email), checkIfPasswordOfUserIsCorrect(pass),
	); err != nil {
		t.Errorf("should have returned the reigstered user: %s", err)
		return
	}

	saved, err := userRepo.FindByEmail(context.Background(), email)
	if err != nil {
		t.Errorf("should have saved the registered user: %s", err)
		return
	}
	if err := assertUser(
		saved,
		checkIfIDOfUserIsFilled(), checkIfEmailOfUserIsSame(email), checkIfPasswordOfUserIsCorrect(pass),
	); err != nil {
		t.Errorf("should have returned the reigstered user: %s", err)
		return
	}

	if _, err := sessRepo.Pull(context.Background()); err != nil {
		t.Errorf("should have push session of the registerd user: %s", err)
		return
	}
}

func TestAuthenticateUser(t *testing.T) {
	userRepo, sessRepo := memory.NewUserRepo(), memory.NewSessionRepo()

	email, pass := "email", "pass"

	regiUser := registerUser{
		userRepo: userRepo,
		sessRepo: sessRepo,
	}
	regiUser.Do(email, pass)

	sessRepo.Delete(context.Background())

	u := authenticateUser{
		userRepo: userRepo,
		sessRepo: sessRepo,
	}
	user, err := u.Do(email, pass)
	if err != nil {
		t.Errorf("should have authenticated user: %s", err)
		return
	}

	sess, err := sessRepo.Pull(context.Background())
	if err != nil {
		t.Errorf("should have pushed session of the authenticated user: %s", err)
		return
	}

	if err := assertUser(
		user,
		checkIfIDOfUserIsSame(sess.UserID()),
	); err != nil {
		t.Errorf("should have returned the authenticaed user: %s", err)
		return
	}
}

func assertUser(u *focus.User, ops ...assertUserOption) error {
	for _, o := range ops {
		if err := o(u); err != nil {
			return err
		}
	}

	return nil
}

type assertUserOption func(u *focus.User) error

func checkIfIDOfUserIsFilled() assertUserOption {
	return func(u *focus.User) error {
		if u.ID() == "" {
			return fmt.Errorf("empty id")
		}

		return nil
	}
}

func checkIfIDOfUserIsSame(id focus.UserID) assertUserOption {
	return func(u *focus.User) error {
		if u.ID() != id {
			return reportUnexpected("id", u.ID(), id)
		}

		return nil
	}
}

func checkIfEmailOfUserIsSame(email string) assertUserOption {
	return func(u *focus.User) error {
		if u.Email() != email {
			return reportUnexpected("email", u.Email(), email)
		}

		return nil
	}
}

func checkIfPasswordOfUserIsCorrect(plain string) assertUserOption {
	return func(u *focus.User) error {
		if !u.Password().IsSame(plain) {
			return fmt.Errorf("incorrect password")
		}

		return nil
	}
}

func reportUnexpected(name string, actual, expected interface{}) error {
	return fmt.Errorf("unexpected %s: got %v, expected %v", name, actual, expected)
}
