package cmd

import (
	"os"

	"github.com/charmbracelet/log"
	"github.com/spf13/cobra"
	"github.com/yurifrl/cli/pkg/config"
)

var debugFlag bool
var appConfig *config.Config

var rootCmd = &cobra.Command{
	Use:   "cli",
	Short: "TODO.",
	Long:  `TODO.`,
	PersistentPreRun: func(cmd *cobra.Command, _ []string) {
		if debugFlag || (appConfig != nil && appConfig.App.Debug) {
			log.SetLevel(log.DebugLevel)
		}
	},
}

func Execute() {
	var err error
	
	// Load configuration
	appConfig, err = config.LoadConfig()
	if err != nil {
		log.Fatal("Failed to load config:", err)
	}

	err = rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().BoolVarP(&debugFlag, "debug", "d", true, "Enable debug mode")
}
