package usecase

import (
	"fmt"
)

func reportUnexpected(name string, actual, expected interface{}) error {
	return fmt.Errorf("unexpected %s: got %v, expected %v", name, actual, expected)
}
