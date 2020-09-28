package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	// color "github.com/bclicn/color"

	grab "github.com/cavaliercoder/grab"
	// aur "github.com/logrusorgru/aurora/v3"
)

// rip immutable arrays
var charSet []string = []string{"⠁", "⠉", "⠙", "⠸", "⢰", "⣠", "⣄", "⡆", "⠇", "⠃", "⠁"}

// Download Downloads a file given a URL
func Download(url string) { // add in args for the file dir path and the files' name(s)
	client := grab.NewClient()
	req, _ := grab.NewRequest(".", url)

	urlStr := strings.SplitN(req.URL().String(), "/", -1)

	fmt.Print(au.Bold(au.Cyan("  Downloading ")))
	fmt.Println(urlStr[len(urlStr)-1])

	response := client.Do(req)

	// fmt.Printf(color.BCyan("      Status: "))
	fmt.Print(au.Bold(au.Cyan("      Status: ")))
	fmt.Printf("%v\n", response.HTTPResponse.Status)

	t := time.NewTicker(100 * time.Millisecond)
	defer t.Stop()

	frame := 0 // set to zero for start, then when it hits the array length, drop it back down to 0
Loop:
	for {
		select {
		case <-t.C:
			fmt.Print("\033[2K") // print the clear-line escape code so that the trailing parenthesis doesn't show up
			// fmt.Printf(color.BYellow("  %v"), charSet[frame])
			fmt.Printf("  %v", au.Bold(au.Yellow(charSet[frame])))
			// fmt.Printf(aur.Bold(aur.Yellow("  %v"), charSet[frame]))
			// fmt.Print(fmt.Sprintf(aur.Bold("  %v").Yellow()), charSet[frame])

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

		case <-response.Done:
			break Loop
		}
	}

	if err := response.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Download failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Print("\033[2K")
	fmt.Print(au.Bold(au.Green("   Downloaded ")))
	fmt.Printf("%v\n", response.Filename)
}

// Extract ...
func Extract(file string) {
	// TODO: add in extracting functionality

}
