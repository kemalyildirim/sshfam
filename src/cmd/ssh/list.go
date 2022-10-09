package ssh

import (
	"fmt"

	"github.com/kemalyildirim/sshfam/src/commons"
	"github.com/spf13/cobra"
)

var listCli = &cobra.Command{
	Use:   commons.ROOT_CMD + " list",
	Short: "List saved sessions. Usage: sshfam list",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("list")

		fmt.Println("args: ", args)
	},
}

func init() {
	rootCmd.AddCommand(listCli)
}
