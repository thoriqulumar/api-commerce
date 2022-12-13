package config

import(
	"github.com/spf13/viper"
	"fmt"
)

func init(){

	viper.SetConfigType("json")
	viper.AddConfigPath("./")
	viper.SetConfigName("config")
	
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
		return
	}
}

func AppName() string{
	return viper.GetString("appName")
}

func ServerPort() string{
	return viper.GetString("server.port")
}

func DatabaseHost() string {
	return viper.GetString("e-commerce.host")
}

func DatabaseName() string {
	return viper.GetString("e-commerce.name")
}

func DatabaseUser() string {
	return viper.GetString("e-commerce.user")
}

func DatabasePass() string {
	return viper.GetString("e-commerce.pass")
}

func DatabasePort() string {
	return viper.GetString("e-commerce.port")
}