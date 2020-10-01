package util

import (
	"fmt"
	"time"

	spin "github.com/briandowns/spinner"
	aur "github.com/logrusorgru/aurora/v3"
	arc "github.com/mholt/archiver/v3"
)

// Extract wraps `archiver.Unarchive()` for simplicity, with added spinner effect.
// Takes path parameter and extracts to dest parameter.
func Extract(path, dest string) error {
	s := spin.New(charSet, 100*time.Millisecond)
	s.Prefix = "  "
	s.Suffix = fmt.Sprintf("%v %v", aur.Bold(aur.Cyan(" Extracting:")), path)
	s.FinalMSG = fmt.Sprintf("   %v %v\n", aur.Bold(aur.Green(" Extracted")), path)
	s.Color("bold", "yellow")

	s.Start()
	err := arc.Unarchive(path, dest)
	s.Stop()

	return err
}
