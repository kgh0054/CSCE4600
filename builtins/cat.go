package builtins

import (
	"fmt"
	"io"
	"os"
)

func OutputFile(w io.Writer, args ...string) error {
	switch len(args) {
	case 0: // change to home directory if available
		return fmt.Errorf("%w: no file specified", ErrInvalidArgCount)
	case 1:
		f, err := os.Open(args[0])
		if err != nil {
			return fmt.Errorf("%w: file not found", ErrInvalidArgCount)
		}

		defer f.Close()
    	buf := make([]byte, 1024)
    	for {
			n, err := f.Read(buf)
			if err == io.EOF {
				break
			}
			if err != nil {
				fmt.Println(err)
				continue
			}
			if n > 0 {
				fmt.Println(string(buf[:n]))
			}
		}

		return err
	default:
		_, err := fmt.Fprintln(w, args)
		return err
	}
}