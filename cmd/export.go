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

// exportCmd represents the export command
var exportCmd = &cobra.Command{
	Use:   "export",
	Short: "Export GNOME Settings",
	Long: `Exports all GNOME Settings. The default behavior is to save the
settings to $XDG_DATA_HOME/gsb/gsb-settings.csv. If $XDG_DATA_HOME is not
defined then it defaults to saving the settings to $HOME/.local/share/gsb/gsb-settings.csv

To use a different file location use the --file flag or setting backup-file in the config file.`,
	Run: exportSettings,
}

func init() {
	rootCmd.AddCommand(exportCmd)

	exportCmd.Flags().StringVarP(&backupFile, "file", "f", "", "specify the file to import from")
}

func exportSettings(cmd *cobra.Command, args []string) {
	var exportFile string
	if backupFile != "" {
		exportFile = backupFile
	} else {
		exportFile = fmt.Sprintf("%s/gsb/gsb-settings.csv", xdg.DataHome)
	}
	fmt.Println(exportFile)
}
