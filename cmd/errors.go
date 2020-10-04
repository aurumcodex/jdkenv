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
	"github.com/aurumcodex/jdkenv/util"

	"github.com/spf13/cobra"
)

// errorCmd represents the error command
var errorsCmd = &cobra.Command{
	Use:   "errors",
	Short: "Print error codes and their meanings",
	Long: `Prints out the error codes and what they mean.
Mostly used for debugging and shouldn't be needed for day to day use.`,
	Run: func(cmd *cobra.Command, args []string) {
		util.PrintErrorList()
	},
}

func init() {
	rootCmd.AddCommand(errorsCmd)
}
