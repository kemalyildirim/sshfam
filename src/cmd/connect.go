package cmd

import (
	"log"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

var connectCli = &cobra.Command{
	Use:   "connect [user]@[remote_ip]",
	Short: "Connect to a server via SSH. Usage: sshfam connect ",
	Long:  "Connect to a server via SSH. Save crendentials on successful login. If the credentials change override them.",
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("connect")
		log.Println("args: ", args)
		execCmd := exec.Command("ssh", args[0])
		execCmd.Stdin = os.Stdin
		execCmd.Stdout = os.Stdout
		execCmd.Stderr = os.Stderr
		log.Println(execCmd.String())
		err := execCmd.Run()
		if err != nil {
			log.Fatal(err)
		}
		log.Println("finished")
	},
}

func init() {
	rootCmd.AddCommand(connectCli)
}
