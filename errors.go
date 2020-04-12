package focus

const (
	ErrNoSuchUser = err("no such user")
)

type err string

func (e err) Error() string {
	return string(e)
}
