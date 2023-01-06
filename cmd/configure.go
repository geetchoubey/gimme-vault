/*
Copyright © 2023 Geet Choubey <geetchoubey@gmail.com>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var keys = [2]string{"username", "accountnumber"}

// configureCmd represents the configure command
var configureCmd = &cobra.Command{
	Use:   "configure",
	Short: "Configure command",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Setting configuration for profile: [%s]\n", profile)
		for _, v := range keys {
			reader.ReadValue(v)
		}
		if err := config.Save(); err != nil {
			panic(fmt.Errorf("error saving configuration %v", err))
		}
	},
}

func init() {
	rootCmd.AddCommand(configureCmd)
}
