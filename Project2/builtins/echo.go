package builtins

import (
	"fmt"
	"io"
)

func Echo(w io.Writer, args ...string) error {
	s := args[0]
	_, err := fmt.Fprintf(w, "%s\n", s)

	return err
}
