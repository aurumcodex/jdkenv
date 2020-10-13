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
	"runtime"
	"strings"

	"github.com/aurumcodex/jdkenv/util"

	"github.com/logrusorgru/aurora/v3"
	"github.com/spf13/cobra"
	"gopkg.in/ini.v1"
)

var (
	jdkVer  int
	noColor bool
	spinner bool
	au      aurora.Aurora

	cfg    *ini.File
	cfgErr error
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Version: "0.1.0",
	Use:     "jdkenv",
	Short:   "A simple Go program to manage and install (if not found) various JDKs.",
	Long: `
jdkenv :: version 0.1.0

A simple Go program to manage and install (if not found) various JDKs.
Running this program without a subcommand will only print the set JDK
and Java version given.

(default JDK and Java version is AdoptOpenJDK 8)`,
	Run: func(cmd *cobra.Command, args []string) {
		util.CheckRuntime()

		switch jdkVer {
		case 8, 11, 14, 15: // will add in support for other Java versions in later updates
			break
		default:
			fmt.Fprintln(os.Stderr, "(e:1)ErrVer - Unknown Java version given:", jdkVer)
			os.Exit(util.ErrVer)
		}

		cfg, cfgErr = ini.LoadSources(ini.LoadOptions{
			SkipUnrecognizableLines: true,
		}, "config.ini")
		if cfgErr != nil {
			fmt.Fprintln(os.Stderr, "(e:2)ErrConf - Config file unable to be read; err =", cfgErr)
			os.Exit(util.ErrConf)
		}

		// consider putting this into another util file to get this info.
		// and to have it be much more flexible
		jdkRI := cfg.Section("").Key("JDK_RI").String()
		jdk := strings.SplitN(jdkRI, ".", -1)
		openj9, _ := cfg.Section("").Key("OPENJ9").Bool()

		fmt.Println("Arch ::", runtime.GOARCH)
		fmt.Println("OS   ::", runtime.GOOS)

		fmt.Println("Java environment:")
		fmt.Printf("Implementation :: %v\n", jdk[0])
		fmt.Printf("  Java Version :: %v\n", jdkVer)
		fmt.Printf("  Using OpenJ9 :: %v\n", openj9)
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
		"spinner",
		true,
		"disables the activity spinner",
	)

	rootCmd.PersistentFlags().IntVarP(
		&jdkVer,
		"java",
		"j",
		8,
		"use specific Java version (valid: 8, 11, 15)",
	)
}
