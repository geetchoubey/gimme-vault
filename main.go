/*
Copyright © 2023 Geet Choubey <geetchoubey@gmail.com>
*/
package main

import (
	"github.com/geetchoubey/gimme-vault/cmd"
	"github.com/geetchoubey/gimme-vault/shared/configuration"
)

const appName = "gimme-vault"

func main() {
	configuration.Init(appName)
	cmd.Execute()
}
