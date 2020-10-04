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
	"time"

	spin "github.com/briandowns/spinner"
	"github.com/logrusorgru/aurora/v3"
	arc "github.com/mholt/archiver/v3"
)

// Extract wraps `archiver.Unarchive()` for simplicity, with added spinner effect.
// Takes path parameter and extracts to dest parameter.
func Extract(path, dest string, spinner, color bool, aur aurora.Aurora) error {
	if spinner {
		s := spin.New(charSet, 100*time.Millisecond)
		s.Prefix = "  "
		s.Suffix = fmt.Sprintf("%v %v", aur.Bold(aur.Cyan(" Extracting:")), path)
		s.FinalMSG = fmt.Sprintf("   %v %v\n", aur.Bold(aur.Green(" Extracted")), path)
		if color {
			s.Color("bold", "yellow")
		}

		s.Start()
		err := arc.Unarchive(path, dest)
		s.Stop()

		return err
	}

	fmt.Printf("Extracting: %v", path)
	err := arc.Unarchive(path, dest)

	return err
}
