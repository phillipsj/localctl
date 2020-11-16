/*
Copyright © 2020 Jamie Phillips

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
	"github.com/mitchellh/go-homedir"
	"github.com/phillipsj/localctl/services"
	"github.com/spf13/cobra"
	"os"
	"path"
	"path/filepath"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Installs app provided by URL.",
	Long: `Installs the application or script provided by the URL into your .local/bin folder:

The goal is to make it easier to download and install applications/scripts into your home 
directory's .local/bin folder'`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		url := args[0]
		filename := path.Base(url)
		home, _ := homedir.Dir()
		expanded, _ := homedir.Expand(home)
		local := filepath.Join(expanded, ".local", "bin")

		if _, err := os.Stat(local); os.IsNotExist(err) {
			os.MkdirAll(local, os.ModePerm)
		}

		fmt.Println("Installing to", local, "...")

		err := services.DownloadFile(url, filepath.Join(local, filename))
		if err != nil {
			fmt.Println("Download failed!")
			os.Exit(1)
		}

		fmt.Println("Download finished!")
	},
}

func init() {
	rootCmd.AddCommand(getCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
