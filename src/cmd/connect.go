package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var connectCli = &cobra.Command{
	Use:   "connect [user]@[remote_ip]",
	Short: "Connect to a server via SSH. Usage: sshfam connect ",
	Long:  "Connect to a server via SSH. Save crendentials on successful login. If the credentials change override them.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("connect")
		fmt.Println("args: ", args)
	},
}

func init() {
	rootCmd.AddCommand(connectCli)
}
