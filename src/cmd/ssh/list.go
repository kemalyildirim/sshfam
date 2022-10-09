package ssh

import (
	"log"

	"github.com/spf13/cobra"
)

var listCli = &cobra.Command{
	Use:   "list",
	Short: "List saved sessions. Usage: sshfam list",
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("list")
		log.Println("args: ", args)
		connections, err := ListConnections()
		if err != nil {
			log.Fatal(err)
		}
		log.Println("connections:\n", connections)
	},
}

func init() {
	rootCmd.AddCommand(listCli)
}
