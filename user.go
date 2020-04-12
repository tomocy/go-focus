package focus

import (
	"context"
	"fmt"
	"time"

	"golang.org/x/crypto/bcrypt"
)

const (
	ErrNoSuchUser = err("no such user")
)

type UserRepo interface {
	NextID(context.Context) (UserID, error)
	FindByEmail(context.Context, string) (*User, error)
	Save(context.Context, *User) error
	Delete(context.Context, UserID) error
}

type User struct {
	id       UserID
	email    string
	password Password
	status   UserStatus
}

func (u User) ID() UserID {
	return u.id
}

func (u *User) setID(id UserID) error {
	if id == "" {
		return fmt.Errorf("empty id")
	}

	u.id = id

	return nil
}

func (u User) Email() string {
	return u.email
}

func (u User) Password() Password {
	return u.password
}

type UserID string

func HashPassword(plain string) (Password, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(plain), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return Password(hash), nil
}

type Password string

func (p Password) IsSame(plain string) bool {
	return bcrypt.CompareHashAndPassword([]byte(p), []byte(plain)) == nil
}

type UserStatus int

const (
	userStatusReady UserStatus = iota
	userStatusFocusing
	userStatusBreaking
)

type Record struct {
	userID UserID
	from   time.Time
	to     time.Time
}

type Timer struct {
	*time.Timer
	duration  time.Duration
	startedAt time.Time
	stoppedAt time.Time
}
