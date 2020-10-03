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

var useOpenJ9 bool

// openjdkCmd represents the openjdk command
var openjdkCmd = &cobra.Command{
	Use:   "openjdk",
	Short: "Set usage of AdoptOpenJDK",
	Long: `Downloads and extracts (if not already existing) the tarball containing
the AdoptOpenJDK implementation with the set JDK version parameter. 
(OpenJ9 VM only available for Java versions 8, 11, and 14)

Java versions supported:
  - 8   (default)
  - 11
  - 14
  - 15`,
	Run: func(cmd *cobra.Command, args []string) {
		util.CheckRuntime()

		valid := util.CheckValidJDK(util.OpenJDK, jdkVer)

		if !valid {
			fmt.Fprintln(os.Stderr, "Invalid Java version passed. Exiting...")
			os.Exit(2)
		}
		fmt.Println("openjdk called")
	},
}

func init() {
	rootCmd.AddCommand(openjdkCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// openjdkCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// openjdkCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	// correttoCmd.Flags().BoolP("8", "", false, "set JDK version to Java 8")
	// correttoCmd.Flags().BoolP("11", "", false, "set JDK version to Java 11")
	openjdkCmd.Flags().BoolVar(&useOpenJ9, "openj9", false, "set the option of using OpenJ9 VM JDK instead of default")
}
