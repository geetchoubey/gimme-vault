/*
Copyright © 2023 Geet Choubey <geetchoubey@gmail.com>
*/
package cli

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/geetchoubey/gimme-vault/shared/configuration"
)

type Reader struct {
	Config configuration.Config
}

func read() string {
	reader := bufio.NewReader(os.Stdin)
	str, _ := reader.ReadString('\n')
	return strings.Trim(str, "\n")
}

func (r *Reader) ReadValue(key string) {
	fmt.Printf("%s[%s] ", key, r.Config.GetString(key))
	value := read()
	if len(value) > 0 {
		r.Config.Set(key, value)
	}
}
