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

// systemCmd represents the system command
var systemCmd = &cobra.Command{
	Use:   "system",
	Short: "Set usage of system installed JDK",
	Long: `Searches for a JDK implementation that was installed via a
package manager on your operating system. Default searches /usr/lib64/ for
an "openjdk-*" directory or similar. Can be set in a toml file called "system.toml"
if your package manager does not install directly to /usr/lib64 (or /usr/lib32 if using a multilib version)`,
	Run: func(cmd *cobra.Command, args []string) {
		util.CheckRuntime()

		fmt.Println("system called")
	},
}

func init() {
	rootCmd.AddCommand(systemCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// systemCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// systemCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	// correttoCmd.Flags().BoolP("8", "", false, "set JDK version to Java 8 (if installed via package manager)")
	// correttoCmd.Flags().BoolP("11", "", false, "set JDK version to Java 11 (if installed via package manager)")
}
