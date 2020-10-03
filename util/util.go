/*
Package util contians various helping functions for jdkenv
(i.e. downloading and extracting compressed tarballs).

Copyright © 2020 Nathan Adams <aurumcodex@protonmail.com>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
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

// PrintErrorList prints out a list of all error codes that can be generated
// when running the program.
func PrintErrorList() {
	fmt.Print(`jdkenv :: error list

_Code_  _ErrName_  _Description_
	0    ErrNone    No errors returned; all functions executed correctly
	1    ErrVer     Invalid/Incompatible version was passed
    2    ErrConf    JDK list or configuration file(s) not found
    3    ErrDL      Downloading error
    4    ErrExtr    Archive extraction error

  124    ErrArch    Architecture not supported
  125    ErrOS      Operating system not supported
`)
}
