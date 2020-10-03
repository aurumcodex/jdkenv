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

	"github.com/logrusorgru/aurora/v3"
	toml "github.com/pelletier/go-toml"
)

// SetCorretto sets the Java environment to use the
func SetCorretto(dest string, version int, spinner bool, aur aurora.Aurora) (error, error, error) {
	jdks, eToml := toml.LoadFile("jdk_list.toml")
	if eToml != nil {
		// fmt.Fprintf(os.Stderr, "Unable to read JDK list file; error = %v", err)
		// os.Exit(1)
		return eToml, nil, nil
	}

	var url string
	switch version {
	case 8:
		url = BuildString(
			jdks.Get("corretto.base_url").(string),
			jdks.Get("corretto.8.file_url").(string),
		)
		// need to add in checks to see if the extracted directory exists first
		if eDL := Download(url, dest, spinner, aur); eDL != nil {
			return nil, eDL, nil
		}
		if eExtr := Extract(jdks.Get("corretto.8.filename").(string), dest, spinner, aur); eExtr != nil {
			return nil, nil, eExtr
		}

	case 11:
		url = BuildString(
			jdks.Get("corretto.base_url").(string),
			jdks.Get("corretto.11.file_url").(string),
		)
		// need to add in checks to see if the extracted directory exists first
		if eDL := Download(url, dest, spinner, aur); eDL != nil {
			return nil, eDL, nil
		}

	default:
		fmt.Fprintln(os.Stderr, "Invalid version passed")
		os.Exit(1)
	}

	// shadow the older error from loading the TOML file
	// err := util.Download(url, dest, spinner, aur)
	// if err != nil {
	// 	fmt.Fprintf(os.Stderr, "Download failed: %v\n", err)
	// 	os.Exit(1)
	// }
	// if eDL := Download(url, dest, spinner, aur); eDL != nil {
	// 	return nil, eDL, nil
	// }

	// shadow the older error from the downloading of tarball
	// err := util.Extract(jdks.Get("corretto."))

	return nil, nil, nil
}
