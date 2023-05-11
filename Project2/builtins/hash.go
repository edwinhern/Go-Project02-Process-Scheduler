package builtins

import (
	"crypto/sha256"
	"fmt"
	"io/ioutil"
	"os"
)

// Hash: Calculates the SHA256 hash of a file
func Hash(args ...string) error {
	if len(args) == 0 {
		// Returns error if no arguments were given
		return fmt.Errorf("%w: expected at least one argument (file)", ErrInvalidArgCount)
	}

	// Calculates the SHA256 hash for each file given in arguments
	for _, file := range args {
		f, err := os.Open(file)
		if err != nil {
			return err // Returns error if the file couldn't be opened
		}
		defer f.Close()

		// Read file content to calculate hash
		content, err := ioutil.ReadAll(f)
		if err != nil {
			return err // Returns error if the file content couldn't be read
		}

		// Calculate hash
		h := sha256.Sum256(content)

		// Print hash
		fmt.Printf("SHA256 hash of %s:\n%x\n", file, h)
	}

	return nil
}
