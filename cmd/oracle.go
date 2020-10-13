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
	"gopkg.in/ini.v1"
)

// oracleCmd represents the oracle command
var oracleCmd = &cobra.Command{
	Use:   "oracle",
	Short: "Set usage of Oracle OpenJDK Reference Implementation",
	Long: `Downloads and extracts (if not already existing) the archive containing
the Oracle OpenJDK implementation with the set JDK version parameter.

Java versions supported:
  - 8   (default)
  - 11
  - 15`,
	Run: func(cmd *cobra.Command, args []string) {
		util.CheckRuntime()

		valid := util.CheckValidJDK(util.OpenJDK, jdkVer)
		if !valid {
			fmt.Fprintln(os.Stderr, "(e:1)ErrVer - ")
			os.Exit(util.ErrVer)
		}

		cfg, cfgErr = ini.LoadSources(ini.LoadOptions{
			SkipUnrecognizableLines: true,
		}, "./config.ini")
		if cfgErr != nil {
			fmt.Fprintln(os.Stderr, "(e:2)ErrConf - Config file unable to be read; err =", cfgErr)
			os.Exit(util.ErrConf)
		}

		dest := cfg.Section("paths").Key("target").String()

		errToml, errDL, errExtr := util.SetOracle(dest, jdkVer, spinner, noColor, au)
		if errToml != nil {
			fmt.Fprintf(os.Stderr, "(e:2)ErrConf - Unable to read jdk_list.toml; error: %v", errToml)
			os.Exit(util.ErrConf)
		} else if errDL != nil {
			fmt.Fprintf(os.Stderr, "(e:3)ErrDL - Error downloading archive; error: %v", errDL)
			os.Exit(util.ErrDL)
		} else if errExtr != nil {
			fmt.Fprintf(os.Stderr, "(e:4)ErrExtr - Error extracting archive; error: %v", errExtr)
			os.Exit(util.ErrExtr)
		}
	},
}

func init() {
	rootCmd.AddCommand(oracleCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// oracleCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// oracleCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	// correttoCmd.Flags().BoolP("8", "", false, "set JDK version to Java 8")
	// correttoCmd.Flags().BoolP("11", "", false, "set JDK version to Java 11")
}
