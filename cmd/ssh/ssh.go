package ssh

import (
	"fmt"

	"github.com/spf13/cobra"
)

var addServerCli = &cobra.Command{
	Use:     "add",
	Short:   "Add a new SSH session",
	Example: "ssh add root@127.0.0.1 -p 22",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("add - sub")
		fmt.Println("args: ", args)
	},
}

func init() {
	rootCmd.AddCommand(addServerCli)
}
