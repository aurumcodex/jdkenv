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

// SetCorretto sets the Java environment to use the Amazon Corretto JDK.
func SetCorretto(dest string, version int, spin, color bool, aur aurora.Aurora) (error, error, error) {
	jdks, eToml := toml.LoadFile("jdk_list.toml")
	if eToml != nil {
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
		eDL := Download(Corretto, version, url, dest, spin, aur)
		if eDL != nil {
			return nil, eDL, nil
		}
		// need to add in checks to see if the extracted directory exists first
		// eExtr := Extract(jdks.Get("corretto.8.filename").(string), dest, spin, color, aur)
		// eExtr := Extract(".jdkenv/amazon_corretto_8.tar.gz", dest, spin, color, aur)
		eExtr := Extract(dest, ".jdkenv/", spin, color, aur)
		if eExtr != nil {
			return nil, nil, eExtr
		}

	case 11:
		url = BuildString(
			jdks.Get("corretto.base_url").(string),
			jdks.Get("corretto.11.file_url").(string),
		)
		// need to add in checks to see if the extracted directory exists first
		eDL := Download(Corretto, version, url, dest, spin, aur)
		if eDL != nil {
			return nil, eDL, nil
		}
		// need to add in checks to see if the extracted directory exists first
		// eExtr := Extract(jdks.Get("corretto.11.file_url").(string), dest, spin, color, aur)
		eExtr := Extract(".jdkenv/amazon_corretto_11.tar.gz", dest, spin, color, aur)
		if eExtr != nil {
			return nil, nil, eExtr
		}

	default:
		fmt.Fprintln(os.Stderr, "(e:1)ErrVer: Invalid version passed:", version)
		os.Exit(ErrVer)
	}

	return nil, nil, nil
}

// SetLiberica sets the Java environment to use the BellSoft Liberica JDK.
func SetLiberica(dest string, version int, spin, color bool, aur aurora.Aurora) (error, error, error) {
	jdks, eToml := toml.LoadFile("jdk_list.toml")
	if eToml != nil {
		return eToml, nil, nil
	}

	var url string
	switch version {
	case 8:
		url = BuildString(
			jdks.Get("liberica.base_url").(string),
			jdks.Get("liberica.8.ext_url").(string),
			jdks.Get("liberica.8.filename").(string),
		)
		// need to add in checks to see if the extracted directory exists first
		eDL := Download(Liberica, version, url, dest, spin, aur)
		if eDL != nil {
			return nil, eDL, nil
		}
		// need to add in checks to see if the extracted directory exists first
		eExtr := Extract(jdks.Get("liberica.8.filename").(string), dest, spin, color, aur)
		if eExtr != nil {
			return nil, nil, eExtr
		}

	case 11:
		url = BuildString(
			jdks.Get("liberica.base_url").(string),
			jdks.Get("liberica.11.ext_url").(string),
			jdks.Get("liberica.11.filename").(string),
		)
		// need to add in checks to see if the extracted directory exists first
		eDL := Download(Liberica, version, url, dest, spin, aur)
		if eDL != nil {
			return nil, eDL, nil
		}
		// need to add in checks to see if the extracted directory exists first
		eExtr := Extract(jdks.Get("liberica.11.filename").(string), dest, spin, color, aur)
		if eExtr != nil {
			return nil, nil, eExtr
		}

	case 15:
		url = BuildString(
			jdks.Get("liberica.base_url").(string),
			jdks.Get("liberica.15.ext_url").(string),
			jdks.Get("liberica.15.filename").(string),
		)
		// need to add in checks to see if the extracted directory exists first
		eDL := Download(Liberica, version, url, dest, spin, aur)
		if eDL != nil {
			return nil, eDL, nil
		}
		// need to add in checks to see if the extracted directory exists first
		eExtr := Extract(jdks.Get("liberica.15.filename").(string), dest, spin, color, aur)
		if eExtr != nil {
			return nil, nil, eExtr
		}

	default:
		fmt.Fprintln(os.Stderr, "(e:1)ErrVer: Invalid version passed:", version)
		os.Exit(ErrVer)
	}

	return nil, nil, nil
}

// SetOpenJDK sets the Java environment to use an AdoptOpenJDK implementation.
func SetOpenJDK(dest string, version int, openj9, spin, color bool, aur aurora.Aurora) (error, error, error) {
	jdks, eToml := toml.LoadFile("jdk_list.toml")
	if eToml != nil {
		return eToml, nil, nil
	}

	var url string
	switch version {
	case 8:
		if openj9 {
			url = BuildString(
				jdks.Get("openjdk.base_url").(string),
				jdks.Get("openjdk.8.ext_url").(string),
				jdks.Get("openjdk.release_path").(string),
				jdks.Get("openjdk.8.openj9.tag").(string),
				jdks.Get("openjdk.8.openj9.filename").(string),
			)
		} else {
			url = BuildString(
				jdks.Get("openjdk.base_url").(string),
				jdks.Get("openjdk.8.ext_url").(string),
				jdks.Get("openjdk.release_path").(string),
				jdks.Get("openjdk.8.tag").(string),
				jdks.Get("openjdk.8.filename").(string),
			)
		}
		// need to add in checks to see if the extracted directory exists first
		var eDL error
		if openj9 {
			eDL = Download(OpenJ9, version, url, dest, spin, aur)
		} else {
			eDL = Download(OpenJDK, version, url, dest, spin, aur)
		}
		if eDL != nil {
			return nil, eDL, nil
		}
		// need to add in checks to see if the extracted directory exists first
		var eExtr error
		if openj9 {
			eExtr = Extract(jdks.Get("openjdk.8.openj9.filename").(string), dest, spin, color, aur)
		} else {
			eExtr = Extract(jdks.Get("openjdk.8.filename").(string), dest, spin, color, aur)
		}
		if eExtr != nil {
			return nil, nil, eExtr
		}

	case 11:
		if openj9 {
			url = BuildString(
				jdks.Get("openjdk.base_url").(string),
				jdks.Get("openjdk.11.ext_url").(string),
				jdks.Get("openjdk.release_path").(string),
				jdks.Get("openjdk.11.openj9.tag").(string),
				jdks.Get("openjdk.11.openj9.filename").(string),
			)
		} else {
			url = BuildString(
				jdks.Get("openjdk.base_url").(string),
				jdks.Get("openjdk.11.ext_url").(string),
				jdks.Get("openjdk.release_path").(string),
				jdks.Get("openjdk.11.tag").(string),
				jdks.Get("openjdk.11.filename").(string),
			)
		}
		// need to add in checks to see if the extracted directory exists first
		eDL := Download(OpenJDK, version, url, dest, spin, aur)
		if eDL != nil {
			return nil, eDL, nil
		}
		// need to add in checks to see if the extracted directory exists first
		var eExtr error
		if openj9 {
			eExtr = Extract(jdks.Get("openjdk.11.openjdk.filename").(string), dest, spin, color, aur)
		} else {
			eExtr = Extract(jdks.Get("openjdk.11.filename").(string), dest, spin, color, aur)
		}
		if eExtr != nil {
			return nil, nil, eExtr
		}

	case 14:
		if openj9 {
			url = BuildString(
				jdks.Get("openjdk.base_url").(string),
				jdks.Get("openjdk.14.ext_url").(string),
				jdks.Get("openjdk.release_path").(string),
				jdks.Get("openjdk.14.openj9.tag").(string),
				jdks.Get("openjdk.14.openj9.filename").(string),
			)
		} else {
			url = BuildString(
				jdks.Get("openjdk.base_url").(string),
				jdks.Get("openjdk.14.ext_url").(string),
				jdks.Get("openjdk.release_path").(string),
				jdks.Get("openjdk.14.tag").(string),
				jdks.Get("openjdk.14.filename").(string),
			)
		}
		// need to add in checks to see if the extracted directory exists first
		eDL := Download(OpenJDK, version, url, dest, spin, aur)
		if eDL != nil {
			return nil, eDL, nil
		}
		// need to add in checks to see if the extracted directory exists first
		eExtr := Extract(jdks.Get("openjdk.14.filename").(string), dest, spin, color, aur)
		if eExtr != nil {
			return nil, nil, eExtr
		}

	case 15:
		url = BuildString(
			jdks.Get("openjdk.base_url").(string),
			jdks.Get("openjdk.15.ext_url").(string),
			jdks.Get("openjdk.release_path").(string),
			jdks.Get("openjdk.15.tag").(string),
			jdks.Get("openjdk.15.filename").(string),
		)
		// need to add in checks to see if the extracted directory exists first
		eDL := Download(OpenJDK, version, url, dest, spin, aur)
		if eDL != nil {
			return nil, eDL, nil
		}
		// need to add in checks to see if the extracted directory exists first
		eExtr := Extract(jdks.Get("openjdk.15.filename").(string), dest, spin, color, aur)
		if eExtr != nil {
			return nil, nil, eExtr
		}

	default:
		fmt.Fprintln(os.Stderr, "(e:1)ErrVer: Invalid version passed:", version)
		os.Exit(ErrVer)
	}

	return nil, nil, nil
}

// SetOracle sets the Java environment to use a reference implementation of OpenJDK built by Oracle.
func SetOracle(dest string, version int, spin, color bool, aur aurora.Aurora) (error, error, error) {
	jdks, eToml := toml.LoadFile("jdk_list.toml")
	if eToml != nil {
		return eToml, nil, nil
	}

	var url string
	switch version {
	case 8:
		url = BuildString(
			jdks.Get("oracle.base_url").(string),
			jdks.Get("oracle.8.ext_url").(string),
			jdks.Get("oracle.8.filename").(string),
		)
		// need to add in checks to see if the extracted directory exists first
		eDL := Download(Oracle, version, url, dest, spin, aur)
		if eDL != nil {
			return nil, eDL, nil
		}
		// need to add in checks to see if the extracted directory exists first
		eExtr := Extract(jdks.Get("oracle.8.filename").(string), dest, spin, color, aur)
		if eExtr != nil {
			return nil, nil, eExtr
		}

	case 11:
		url = BuildString(
			jdks.Get("oracle.base_url").(string),
			jdks.Get("oracle.11.ext_url").(string),
			jdks.Get("oracle.11.filename").(string),
		)
		// need to add in checks to see if the extracted directory exists first
		eDL := Download(Oracle, version, url, dest, spin, aur)
		if eDL != nil {
			return nil, eDL, nil
		}
		// need to add in checks to see if the extracted directory exists first
		eExtr := Extract(jdks.Get("oracle.11.filename").(string), dest, spin, color, aur)
		if eExtr != nil {
			return nil, nil, eExtr
		}

	case 15:
		url = BuildString(
			jdks.Get("oracle.base_url").(string),
			jdks.Get("oracle.15.ext_url").(string),
			jdks.Get("oracle.15.filename").(string),
		)
		// need to add in checks to see if the extracted directory exists first
		eDL := Download(Oracle, version, url, dest, spin, aur)
		if eDL != nil {
			return nil, eDL, nil
		}
		// need to add in checks to see if the extracted directory exists first
		eExtr := Extract(jdks.Get("oracle.15.filename").(string), dest, spin, color, aur)
		if eExtr != nil {
			return nil, nil, eExtr
		}

	default:
		fmt.Fprintln(os.Stderr, "(e:1)ErrVer: Invalid version passed:", version)
		os.Exit(ErrVer)
	}

	return nil, nil, nil
}
