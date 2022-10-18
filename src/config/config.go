package config

import (
	"io/fs"
	"log"
	"os"

	"github.com/kemalyildirim/sshfam/src/commons"
	"github.com/spf13/viper"
)

var CFG = viper.New()
var homeDir, _ = os.UserHomeDir()
var CONFIG_PATH = homeDir + string(os.PathSeparator) + "." + commons.APP_NAME + string(os.PathSeparator)

func initViper() {
	CFG.SetConfigName(commons.APP_NAME)
	CFG.SetConfigType(commons.CONFIG_TYPE)
	os.Mkdir(CONFIG_PATH, fs.ModeDir)
	CFG.AddConfigPath(CONFIG_PATH)
	CFG.WatchConfig()

	if err := CFG.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; ignore error if desired
			log.Println("viper: config not found")
			CFG.SafeWriteConfig()
		} else {
			// Config file was found but another error was produced
			log.Println("viper: config found with error")
			panic(err)
		}
	}
}

func SaveCred(con commons.Connection) {
	CFG.Set("mycon", con)
	CFG.WriteConfig()
}
