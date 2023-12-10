package builtins

import (
	"errors"
	"fmt"
	"io"
)

var (
	ErrInvalidArgNum = errors.New("invalid argument count")
)

func Echo(w io.Writer, args ...string) error {
	switch len(args) {
	case 0: // change to home directory if available
		return fmt.Errorf("%w: nothing to echo", ErrInvalidArgNum)
	default:
		_, err := fmt.Fprintln(w, args)
		return err
	}
}