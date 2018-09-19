package actionerror

import "fmt"

// BuildpacksStackMismatchError is returned when one or more of the requested buildpacks
// specified during a push does not exist for the application's stack.
type BuildpacksStackMismatchError struct {
	Stack string
}

// Error method to display the error message.
func (e BuildpacksStackMismatchError) Error() string {
	return fmt.Sprintf("Buildpacks not found for stack %s", e.Stack)
}
