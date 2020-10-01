package util

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/cavaliercoder/grab"
	aur "github.com/logrusorgru/aurora/v3"
)

var charSet = []string{"⠁", "⠉", "⠙", "⠸", "⢰", "⣠", "⣄", "⡆", "⠇", "⠃", "⠁"}

// Download gets a file from a given string and stores it.
func Download(url string) {
	// consider adding parameters for aurora state and target destination
	client := grab.NewClient()
	req, _ := grab.NewRequest(".", url)

	urlStr := strings.SplitN(req.URL().String(), "/", -1)
	fmt.Print(aur.Bold(aur.Cyan("  Downloading ")))
	fmt.Println(urlStr)

	response := client.Do(req)

	fmt.Print(aur.Bold(aur.Cyan("      Status: ")))
	fmt.Println(response.HTTPResponse.Status)

	t := time.NewTicker(100 * time.Millisecond)
	defer t.Stop()

	frame := 0
Loop:
	for {
		select {
		case <-t.C:
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

		case <-response.Done:
			break Loop
		}
	}

	if err := response.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Download failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Print("\033[2K")
	fmt.Print(aur.Bold(aur.Green("  Downloaded ")))
	fmt.Println(response.Filename)
}
