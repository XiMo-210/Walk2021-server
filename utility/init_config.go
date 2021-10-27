package utility

import (
	"fmt"
	"github.com/spf13/viper"
)

var Config = viper.New()

func InitConfig() {
	Config.SetConfigType("yaml")
	Config.AddConfigPath("./config")
	Config.SetConfigName("config")
	Config.WatchConfig() // Watch the change of the config file

	err := Config.ReadInConfig()
	if err != nil {
		fmt.Println("配置读取错误! ")
		fmt.Println(err)
	}

	fmt.Println(Config.GetString("server.wechatAPPID"), Config.GetString("server.wechatSecret"))
}
