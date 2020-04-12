package usecase

import "github.com/tomocy/focus"

type registerUser struct {
	userRepo focus.UserRepo
	sessRepo focus.SessionRepo
}
