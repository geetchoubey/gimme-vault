/*
Copyright © 2023 Geet Choubey <geetchoubey@gmail.com>
*/
package cmd

import (
	"os"

	"github.com/geetchoubey/gimme-vault/shared/cli"
	"github.com/geetchoubey/gimme-vault/shared/configuration"
	"github.com/spf13/cobra"
)

var (
	profile string = "default"
	config  configuration.Config
	reader  cli.Reader
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:     "gimme-vault",
	Short:   "gimme-vault command to login to vault and generate AWS token",
	Long:    "",
	Version: "1.0.0",
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVar(&profile, "profile", profile, "Profile name")
	config = configuration.Config{
		Namespace: &profile,
	}
	reader = cli.Reader{
		Config: config,
	}
}
