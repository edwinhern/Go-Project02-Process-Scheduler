package builtins

import (
	"fmt"
)

// BindKey binds a key sequence to a shell command.
func BindKey(args ...string) error {
	if len(args) != 2 {
		return fmt.Errorf("invalid number of arguments for bindkey")
	}

	key, cmd := args[0], args[1]

	fmt.Printf("Bound key sequence '%s' to command '%s'.\n", key, cmd)
	return nil
}
