package focus

import "golang.org/x/crypto/bcrypt"

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
