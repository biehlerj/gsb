/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os/exec"

	"github.com/spf13/cobra"
)

// exportCmd represents the export command
var exportCmd = &cobra.Command{
	Use:   "export",
	Short: "Exports GNOME settings",
	Long: `Exports the following GNOME settings:

	1. WM specific keybindings
	2. Non-WM specific keybindings
	3. Power related buttons
	4. Background
	5. Peripherals
	6. Default Applications
	6. Privacy settings
	7. WM preferences

Additionally you can also specifiy which settings to export
	`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("export called")
	},
}

func init() {
	rootCmd.AddCommand(exportCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// exportCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// exportCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func wmKeybindingExport() {
	// Step 1: Run gsettings list-recursively org.gnome.desktop.wm.keybindings
	gsettingsCmd, err := exec.LookPath("gsettings")
	if err != nil {

	}
	// Step 2: Format output into csv
	// Step 3: Make a file named keybindings.wm.csv in current directory
	// Step 4a: Put formatted csv output into exsiting or created file
	// Step 4b: If error opening or creating file exit
}
