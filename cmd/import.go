package cmd

import (
	"fmt"

	"github.com/adrg/xdg"

	"github.com/spf13/cobra"
)

// importCmd represents the command to import GNOME settings
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
