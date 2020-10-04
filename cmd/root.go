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

	"github.com/logrusorgru/aurora/v3"
	toml "github.com/pelletier/go-toml"
	"github.com/spf13/cobra"
)

var jdkVer int
var noColor bool
var spinner bool
var au aurora.Aurora

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Version: "0.1.0",
	Use:     "jdkenv",
	Short:   "A simple Go program to manage and install (if not found) various JDKs.",
	Long: `jdkenv :: version 0.1.0
A simple Go program to manage and install (if not found) various JDKs.
Running this program without a subcommand will only print the set JDK and Java version.

(default JDK and Java version is AdoptOpenJDK 8)`,
	Run: func(cmd *cobra.Command, args []string) {
		util.CheckRuntime()

		switch jdkVer {
		case 8, 11, 14, 15: // will add in support for other Java versions in later updates
			break
		default:
			fmt.Fprintln(os.Stderr, "Error: Unknown JDK type")
			os.Exit(1)
		}

		jdks, err := toml.LoadFile("jdk_list.toml")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Unable to read JDK list file; error = %v", err)
			os.Exit(1)
		}

		test, _ := toml.LoadFile("test.toml")

		directOpenJDKBaseURL := jdks.Get("openjdk.base_url").(string)
		testArray := jdks.GetArray("test")
		testArray2 := jdks.GetArrayPath([]string{"name", "path"})
		mapT := test.ToMap()

		fmt.Println("oracle base_url:", directOpenJDKBaseURL)
		fmt.Println(testArray)
		fmt.Println(testArray2)
		fmt.Println(mapT)
		fmt.Println(mapT["test.name"])
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().BoolVar(
		&noColor,
		"no-color",
		false,
		"use monochrome output",
	)
	au = aurora.NewAurora(!noColor)

	rootCmd.PersistentFlags().BoolVar(
		&spinner,
		"no-spinner",
		false,
		"disables the activity spinner (useful for CI or testing)",
	)

	rootCmd.PersistentFlags().IntVarP(
		&jdkVer,
		"java",
		"j",
		8,
		"use specific Java version (valid: 8, 11, 14, 15)",
	)
}
