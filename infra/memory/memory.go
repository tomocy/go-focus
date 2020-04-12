package memory

import "github.com/tomocy/focus"

type userRepo struct {
	users map[focus.UserID]*focus.User
}
