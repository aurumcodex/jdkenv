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

	"github.com/logrusorgru/aurora"
	toml "github.com/pelletier/go-toml"
	"github.com/spf13/cobra"

	"github.com/aurumcodex/jdkenv/util"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Version: "0.1.0",
	Use:     "jdkenv",
	Short:   "A simple Go program to manage and install (if not found) various JDKs.",
	Long: `jdkenv :: version 0.1.0
A simple Go program to manage and install (if not found) various JDKs.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		type Config struct {
			oracle   util.Oracle
			corretto util.Corretto
		}

		tomlCfg, _ := toml.LoadFile("./jdk_list.toml")

		var conf Config
		tomlCfg.Unmarshal(&conf)

		// fmt.Println(conf)
		fmt.Println("oracle base_url:", conf.oracle.Base)
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

var color bool
var au aurora.Aurora

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.jdkenv.yaml)")

	rootCmd.PersistentFlags().BoolVar(&color, "no-color", false, "set colorized or monochrome output")
	au = aurora.NewAurora(!color)

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	// rootCmd.Flags().BoolP("version", "v", false, "Prints version")
}
