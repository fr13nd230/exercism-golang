package twofer

import (
	"fmt"
	"strings"
)

func ShareWith(name string) string {
	if len(strings.TrimSpace(name)) == 0 {
		return "One for you, one for me."
	}
	return fmt.Sprintf("One for %s, one for me.	", name)
}
