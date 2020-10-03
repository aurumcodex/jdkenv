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

const (
	// ErrNone is the error code that means that nothing bad happened when executing.
	ErrNone int = 0
	// ErrVer is an error code for an invalid JDK/Java version passed into program.
	ErrVer int = 1
	// ErrConf is an error code for a configuration file not found.
	ErrConf int = 2
	// ErrDL is an error code for a downloading failure.
	ErrDL int = 3
	// ErrExtr is an error code for an extraction failure.
	ErrExtr int = 4

	// ErrArch is an error code for incompatible CPU architecture.
	// * Currently unused
	ErrArch int = 124
	// ErrOS is an error code for incompatible OS.
	ErrOS int = 125
)
