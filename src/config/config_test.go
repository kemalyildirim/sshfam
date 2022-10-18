package config

import (
	"errors"
	"log"
	"os"
	"testing"
	"time"

	"github.com/fsnotify/fsnotify"
)

var VIPER_INSTANCE_TEST = CFG

func setup(tb testing.TB) func(tb testing.TB) {
	log.Println("setup test")
	InitViper()
	return teardown
}

func teardown(tb testing.TB) {
	log.Println("teardown test")
	configPath := VIPER_INSTANCE_TEST.ConfigFileUsed()
	log.Println("config path: ", configPath)
	config, _ := os.ReadFile(configPath)
	log.Println("config:\n", string(config))
	err := os.RemoveAll(CONFIG_PATH)
	if err != nil {
		panic("[teardown] Can't remove the config file")
	}
}

func TestConfigInit(t *testing.T) {
	teardown := setup(t)
	defer teardown(t)

	myMap := make(map[string]string)
	myMap["user1"] = "mysecretencryptedpassword"
	VIPER_INSTANCE_TEST.Set("testMap", myMap)
	VIPER_INSTANCE_TEST.WriteConfig()
	if _, err := os.Stat(CONFIG_PATH); errors.Is(err, os.ErrNotExist) {
		t.Fatal("config file is not created")
	}
}

func TestWatchConfig(t *testing.T) {
	teardown := setup(t)
	defer teardown(t)
	configChanged := false
	myMap := make(map[string]string)
	myMap["user1"] = "mysecretencryptedpassword"
	VIPER_INSTANCE_TEST.Set("testMap", myMap)
	VIPER_INSTANCE_TEST.OnConfigChange(func(in fsnotify.Event) {
		log.Println("event: ", in.Name)
		configChanged = true
	})
	time.Sleep(time.Second * 2)
	VIPER_INSTANCE_TEST.Set("mydummy", "configchanged")
	VIPER_INSTANCE_TEST.WriteConfig()
	time.Sleep(time.Second * 1)
	if !configChanged {
		t.Fatal("config didn't change")
	}
}

// TODO: Some errors when tests run, not failing the test though.
