package main

import (
	"log"
	"os"

	"github.com/kemalyildirim/sshfam/cmd/ssh"
	"github.com/spf13/viper"
)

func initConfig() {
	viper.SetConfigName(ssh.ROOT_CMD)
	viper.SetConfigType("properties")
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	viper.AddConfigPath(homeDir)
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; ignore error if desired
			log.Println("viper: config not found")
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
