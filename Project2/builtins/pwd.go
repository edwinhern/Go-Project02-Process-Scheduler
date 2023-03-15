package builtins

import (
	"fmt"
	"io"
	"os"
)

// CurrentDirectory: Writes the current working directory
func CurrentDirectory(w io.Writer, args ...string) error {
	// Fetches the current working directory
	wd, err := os.Getwd()
	if err != nil {
		return err // Returns error if the directory couldn't be fetched
	}
	// Writes the current working directory to the writer
	_, err = fmt.Fprintf(w, "%s\n", wd)
	// Returns any error that occurred while writing to the writer
	return err
}
