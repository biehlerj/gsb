package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// exportCmd represents the command to export GNOME settings
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
		exportFile = viper.GetString("backup-file")
	}
	fmt.Println(exportFile)
}
