package types

import (
	"fmt"
	"strings"

	"github.com/kemalyildirim/sshfam/src/commons"
	"github.com/spf13/viper"
)

type RemoteConnection struct {
	ServerIp string
	Port     int
	User     string
	Password string
}

func Parse(in string) *RemoteConnection {
	var con *RemoteConnection = new(RemoteConnection)
	splited := strings.Split(in, "@")
	con.User = splited[0]
	con.ServerIp = splited[1]
	con.Port = 22
	// con.Password = nil
	return con
}

func Save(con RemoteConnection) {
	conMap := viper.GetStringMap(commons.VIPER_CONNECTIONS_MAP)
	fmt.Println(conMap)
	// conMap[con.ServerIp][con.User] = con.Password
}
