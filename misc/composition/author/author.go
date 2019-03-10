package author

import (
	"fmt"
)

// Author represent a blog author
type Author struct {
	FistName string
	LastName string
	Bio      string
}

// FullName returns author full name
func (a Author) FullName() string {
	return fmt.Sprintf("%s %s", a.FistName, a.LastName)
}
