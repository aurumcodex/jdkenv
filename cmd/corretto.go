/*
Package cmd contians various subcommands for jdkenv to run correctly.

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
package cmd

import (
	"fmt"
	"os"

	"github.com/aurumcodex/jdkenv/util"

	"github.com/spf13/cobra"
)

// correttoCmd represents the corretto command
var correttoCmd = &cobra.Command{
	Use:   "corretto",
	Short: "Set usage of Amazon Corretto JDK",
	Long: `Downloads and extracts (if not already existing) the tarball containing
the Amazon Corretto JDK with set JDK version parameter.

Java versions supported:
  - 8    (default)
  - 11`,
	Run: func(cmd *cobra.Command, args []string) {
		util.CheckRuntime()

		valid := util.CheckValidJDK(util.OpenJDK, jdkVer)
		if !valid {
			fmt.Fprintln(os.Stderr, "Invalid Java version passed. Exiting...")
			os.Exit(2)
		}

		errToml, errDL, errExtr := util.SetCorretto("", jdkVer, spinner, noColor, au)
		if errToml != nil {
			fmt.Fprintf(os.Stderr, "(e:2)ErrConf - Unable to read jdk_list.toml; error: %v", errToml)
			os.Exit(util.ErrConf)
		} else if errDL != nil {
			fmt.Fprintf(os.Stderr, "(e:3)ErrDL - Error downloading archive; error: %v", errDL)
			os.Exit(util.ErrDL)
		} else if errExtr != nil {
			fmt.Fprintf(os.Stderr, "(e:2)ErrConf - Error extracting archive; error: %v", errExtr)
			os.Exit(util.ErrExtr)
		}
	},
}

func init() {
	rootCmd.AddCommand(correttoCmd)
}
