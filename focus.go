package focus

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	id       UserID
	email    string
	password Password
	status   UserStatus
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
