package util

import (
	"fmt"
	"os"
	"runtime"
	"strings"
)

// const ()

// BuildString collects inputs passed to it and concatenates them into one larger string.
// Used for building URLs for use in downloading JDKs and other files
func BuildString(inputs ...string) string {
	var sb strings.Builder

	for _, s := range inputs {
		sb.WriteString(s)
	}

	return sb.String()
}

// CheckRuntime is a wrapper to check if the environment is Windows or not.
func CheckRuntime() {
	if runtime.GOOS == "windows" {
		fmt.Println("This application was not designed to be used in a Windows environment at this time.")
		os.Exit(125)
	}
}

// CheckValidJDK checks the given integer(s) and returns `true` if valid JDK is available,
// otherwise returns `false`.
func CheckValidJDK(jdk, version int) bool {
	switch jdk {
	case Corretto:
		switch version {
		case 8, 11:
			return true
		default:
			return false
		}

	case Liberica:
		switch version {
		case 8, 11, 15:
			return true
		default:
			return false
		}

	case OpenJDK:
		switch version {
		case 8, 11, 14, 15:
			return true
		default:
			return false
		}

	case Oracle:
		switch version {
		case 8, 11, 15:
			return true
		default:
			return false
		}

	default:
		return false
	}
}
