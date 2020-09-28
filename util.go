package main

import "strings"

// BuildString Builds a usable string from some inputs
func BuildString(inputs ...string) string {
	var sb strings.Builder

	for _, s := range inputs {
		sb.WriteString(s)
	}

	str := sb.String()
	return str
}
