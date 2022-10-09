package main

import (
	"log"

	"github.com/kemalyildirim/sshfam/src/cmd/ssh"
	"github.com/kemalyildirim/sshfam/src/commons"
	"github.com/spf13/viper"
)

func initConfig() {
	viper.SetConfigName(commons.ROOT_CMD)
	viper.SetConfigType("properties")
	// TODO
	// homeDir, err := os.UserHomeDir()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// viper.AddConfigPath(homeDir)
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; ignore error if desired
			log.Println("viper: config not found")
			viper.SafeWriteConfig()
		} else {
			// Config file was found but another error was produced
			log.Println("viper: config found with error")
			log.Fatal(err)
		}
	}

	// Config file found and successfully parsed
}

func main() {
	initConfig()
	ssh.Execute()
}
