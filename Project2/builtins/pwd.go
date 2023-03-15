package builtins

import (
	"fmt"
	"io"
	"os"
)

func CurrentDirectory(w io.Writer, args ...string) error {
	wd, err := os.Getwd()
	if err != nil {
		return err
	}
	_, err = fmt.Fprintf(w, "%s\n", wd)
	return err
}
