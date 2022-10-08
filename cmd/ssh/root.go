package ssh

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   ROOT_CMD,
	Short: ROOT_CMD + " - Manage ssh saved instances",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("root")
		fmt.Println("args: ", args)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "There was an error while executing '%s'", err)
		os.Exit(1)
	}
}
