package ssh

import (
	"fmt"

	"github.com/spf13/cobra"
)

var addServerCli = &cobra.Command{
	Use:     "add [user]@[remote_ip]:[port]",
	Short:   "Add a new SSH session",
	Long:    "Add a new SSH session. Default port is 22.",
	Example: ROOT_CMD + " add root@127.0.0.1\n" + ROOT_CMD + " add test@10.254.11.5:550",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("add - sub")
		fmt.Println("args: ", args)
	},
}

func init() {
	rootCmd.AddCommand(addServerCli)
}
