package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Display or change configuration settings for gh.",
	Long: `Current respected settings:
- backup-file
- config-file
`,
}

var configFile string

func init() {
	rootCmd.AddCommand(configCmd)

	configCmd.Flags().StringVarP(&configFile, "file", "f", "", "specify the config file to use")
	configCmd.AddCommand(setConfigKey, getConfigKey, listConfigKey)
}

var setConfigKey = &cobra.Command{
	Use:   "set",
	Short: "Update configuration with a value for the given key",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var getConfigKey = &cobra.Command{
	Use:   "get",
	Short: "Print the value of a given configuration key",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(viper.Get(args[0]))
	},
}

var listConfigKey = &cobra.Command{
	Use:   "list",
	Short: "Print a list of configuration keys and values",
	Run: func(cmd *cobra.Command, args []string) {
		allSettings := viper.AllSettings()
		for k, v := range allSettings {
			fmt.Printf("%s=%s\n", k, v)
		}
	},
}
