package cmd

import (
	"log"
	"os"
	"os/exec"

	"github.com/kemalyildirim/sshfam/src/commons"
	"github.com/kemalyildirim/sshfam/src/config"
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
		channel := make(chan error)
		go func() {
			err := execCmd.Run()
			channel <- err
		}()
		go func(err error) {
			if err != nil {
				log.Println("test")
				var con commons.Connection = commons.Connection{}
				con.Cred = args[0]
				con.Pass = "todo"
				config.SaveCred(con)
			} else {
				log.Fatal("connection failed")
			}
		}(<-channel)

		log.Println("finished")
	},
}

func init() {
	rootCmd.AddCommand(connectCli)
}
