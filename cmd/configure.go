/*
Copyright © 2023 Geet Choubey <geetchoubey@gmail.com>
*/
package cmd

import (
	"github.com/geetchoubey/gimme-vault/shared/configuration"
	"github.com/spf13/cobra"
	jww "github.com/spf13/jwalterweatherman"
)

// configureCmd represents the configure command
var configureCmd = &cobra.Command{
	Use:   "configure",
	Short: "Configure command",
	Run: func(cmd *cobra.Command, args []string) {
		jww.INFO.Printf("Setting configuration for profile: [%s]\n", profile)
		for _, v := range configuration.Keys {
			reader.ReadValue(v)
		}
		if err := config.Save(); err != nil {
			jww.CRITICAL.Panicf("error saving configuration %v", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(configureCmd)
}

func IsProfileConfigured() {
	if isInit := config.IsConfigInitialized(); !isInit {
		jww.INFO.Printf("No configuration found for profile %s.", profile)
		jww.CRITICAL.Panicf("run `gimme-vault configure --profile`")

	}
}
