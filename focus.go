package focus

type User struct {
	id       UserID
	email    string
	password Password
}

type UserID string

type Password string
