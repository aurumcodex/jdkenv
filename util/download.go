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
	"strings"
	"time"

	"github.com/cavaliercoder/grab"
	"github.com/logrusorgru/aurora/v3"
)

var charSet = []string{"⠁", "⠉", "⠙", "⠸", "⢰", "⣠", "⣄", "⡆", "⠇", "⠃", "⠁"}

// Download gets a file from a given string and stores it.
func Download(url, dest string, spinner bool, aur aurora.Aurora) error {
	// consider adding parameters for aurora state and target destination
	client := grab.NewClient()
	req, _ := grab.NewRequest(dest, url)

	urlStr := strings.SplitN(req.URL().String(), "/", -1)
	fmt.Print(aur.Bold(aur.Cyan("  Downloading ")))

	fmt.Println(urlStr)

	response := client.Do(req)

	fmt.Print(aur.Bold(aur.Cyan("      Status: ")))
	fmt.Println(response.HTTPResponse.Status)

	var t *time.Ticker
	if spinner {
		t = time.NewTicker(100 * time.Millisecond)
		defer t.Stop()
	}

	frame := 0
Loop:
	for {
		select {
		case <-t.C:
			if spinner {
				fmt.Print("\033[2K") // print a "clear-line" to prevent trailing parentheses
				fmt.Printf("  %v", aur.Bold(aur.Yellow(charSet[frame])))

				fmt.Printf(" transferred [%.2f / %.2f MiB] @ %.2f KiB/s (%.2f%%)\r",
					float64(response.BytesComplete())/1024.0/1024.0,
					float64(response.Size)/1024.0/1024.0,
					float64(response.BytesPerSecond())/1024.0,
					100*response.Progress(),
				)

				if frame < len(charSet)-1 {
					frame++
				} else if frame == len(charSet)-1 {
					frame = 0
				}
			} else {
				fmt.Printf(" transferred [%.2f / %.2f MiB] @ %.2f KiB/s (%.2f%%)\n",
					float64(response.BytesComplete())/1024.0/1024.0,
					float64(response.Size)/1024.0/1024.0,
					float64(response.BytesPerSecond())/1024.0,
					100*response.Progress(),
				)
			}

		case <-response.Done:
			break Loop
		}
	}

	if response.Err() == nil {
		fmt.Print("\033[2K")
		fmt.Printf("%v %v\n", aur.Bold(aur.Green("  Downloaded ")), response.Filename)
	}

	return response.Err()
}
