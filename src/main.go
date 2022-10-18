package main

import (
	"github.com/kemalyildirim/sshfam/src/cmd"
	"github.com/kemalyildirim/sshfam/src/config"
)

func main() {
	config.InitViper()
	cmd.Execute()
}
