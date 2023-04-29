/*
Copyright © 2023 Jacob Biehler <jacob@biehlerj.xyz>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"fmt"

	"github.com/adrg/xdg"

	"github.com/spf13/cobra"
)

// importCmd represents the import command
var importCmd = &cobra.Command{
	Use:   "import",
	Short: "Import GNOME settings",
	Long: `Imports GNOME settings. The default behavior is to look for gsb-settings.csv in the $XDG_DATA_HOME directory.
If $XDG_DATA_HOME is not defined it defaults to looking in ~/.local/share/gsb.

You can specify where to find the file by setting backup-file in the config file or passing the --file flag.`,
	Run: importSettings,
}

func init() {
	rootCmd.AddCommand(importCmd)

	importCmd.Flags().StringVarP(&backupFile, "file", "f", "", "specify the file to import from")
}

func importSettings(cmd *cobra.Command, args []string) {
	var importFile string
	if backupFile != "" {
		importFile = backupFile
	} else {
		importFile = fmt.Sprintf("%s/gsb/gsb-settings.csv", xdg.DataHome)
	}
	fmt.Println(importFile)
}
