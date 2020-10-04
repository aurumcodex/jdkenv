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

// localCmd represents the local command
var localCmd = &cobra.Command{
	Use:   "local",
	Short: "Set usage of a locally installed JDK",
	Long: `Sets the Java environment to use whichever JDK option you pass to it,
assuming that it has a name and a path given in a local.toml configuration file.
Whatever you choose, it *must* match whatever you named it within the local.toml file
in the configuration directory.`,
	Run: func(cmd *cobra.Command, args []string) {
		util.CheckRuntime()

		fmt.Println("local called")
	},
}

func init() {
	rootCmd.AddCommand(localCmd)
}
