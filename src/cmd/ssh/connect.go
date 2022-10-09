package ssh

import (
	"log"

	. "github.com/kemalyildirim/sshfam/src/types"
	"github.com/spf13/cobra"
)

var connectCli = &cobra.Command{
	Use:   "connect [user]@[remote_ip]",
	Short: "Connect to a server via SSH. Usage: sshfam connect ",
	Long:  "Connect to a server via SSH. Save crendentials on successful login. If the credentials change override them.",
	Run: func(cmd *cobra.Command, args []string) {
		con := Parse(args[0])
		err := Connect(con)
		if err != nil {
			log.Fatal(err)
		}
		log.Println("[connect] success")
	},
}

func init() {
	rootCmd.AddCommand(connectCli)
}
