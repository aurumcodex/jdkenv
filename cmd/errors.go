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
	Long:  `Prints out the error codes and what they mean.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: add a PrintErrorList() function to util package to have this functionality
		// (even if it practice it's a big fmt.Printf call)
		util.PrintErrorList()
		// fmt.Println("errors called")
	},
}

func init() {
	rootCmd.AddCommand(errorsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// errorCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// errorCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
