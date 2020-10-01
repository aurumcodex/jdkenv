package util

import "strings"

const (
	openJDK      int = (iota + 1)
	corretto     int = (iota + 1)
	liberica     int = (iota + 1)
	adoptOpenJDK int = (iota + 1)

	v8  int = 8
	v11 int = 11
	v14 int = 14
	v15 int = 15
)

// BuildString collects inputs passed to it and concatenates them into one larger string.
// Used for building URLs for use in downloading JDKs and other files
func BuildString(inputs ...string) string {
	var sb strings.Builder

	for _, s := range inputs {
		sb.WriteString(s)
	}

	return sb.String()
}
