package builtins

import (
	"fmt"
	"io"
	"os"
)

func Cat(args ...string) error {
	// if no files are specified, read from stdin
	if len(args) == 0 {
		_, err := io.Copy(os.Stdout, os.Stdin)
		return err
	}

	// iterate over files and output their contents
	for _, file := range args {
		f, err := os.Open(file)
		if err != nil {
			return fmt.Errorf("failed to open file: %v", err)
		}
		defer f.Close()

		_, err = io.Copy(os.Stdout, f)
		if err != nil {
			return fmt.Errorf("failed to read file: %v", err)
		}
	}

	return nil
}
