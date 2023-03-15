package builtins

import (
	"fmt"
	"os"
)

// MakeDirectory: Creates one or more directories with the given names
func MakeDirectory(args ...string) error {
	switch len(args) {
	case 0:
		// Returns error if no arguments were given
		return fmt.Errorf("%w: expected at least one argument (directory)", ErrInvalidArgCount)
	default:
		// Creates each directory
		for _, dir := range args {
			// Creates the directory with directory name and permissions
			err := os.MkdirAll(dir, 0755)
			if err != nil {
				return err // Returns error if the directory couldn't be created
			}
			// Success message
			fmt.Printf("Created directory: %s\n", dir)
		}
		return nil
	}
}
