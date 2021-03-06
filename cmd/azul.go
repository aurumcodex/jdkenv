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

	"github.com/spf13/cobra"
)

// azulCmd represents the azul command
var azulCmd = &cobra.Command{
	Use:   "azul",
	Short: "Set usage of Azul Zulu Community JDK",
	Long: `Downloads and extracts (if not already existing) the archive containing
the Azul Zulu Community JDK with set JDK version parameter.

Java versions supported:
  - 8    (default)
  - 11
  - 15`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("azul called")
	},
}

func init() {
	rootCmd.AddCommand(azulCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// azulCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// azulCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
