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

		fmt.Println("corretto called")
	},
}

func init() {
	rootCmd.AddCommand(correttoCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// correttoCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// correttoCmd.Flags().BoolP("8", "", false, "set JDK version to Java 8")
	// correttoCmd.Flags().BoolP("11", "", false, "set JDK version to Java 11")
}
