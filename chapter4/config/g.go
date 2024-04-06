package config

import (
	"fmt"
	"github.com/spf13/viper"
)

func init() {
	projectName := "go-mega"
	getConfig(projectName)
}
func getConfig(projectName string) {
	//设置配置文件的名称"config"
	viper.SetConfigName("config")
	//添加配置文件的搜索路径，包括当前路径，用户主目录和Docker配置目录
	viper.AddConfigPath(".")
	viper.AddConfigPath(fmt.Sprintf("$HOME/.%s", projectName))
	viper.AddConfigPath(fmt.Sprintf("/data/docker/config/%s", projectName))

	//读取配置文件
	err := viper.ReadInConfig()
	if err != nil {
		//如果读取配置文件错误
		panic(fmt.Errorf("Fatal error config file:%s", err))
	}
}

func GetMysqlConnectingString() string {
	usr := viper.GetString("mysql.user")
	pwd := viper.GetString("mysql.password")
	host := viper.GetString("mysql.host")
	db := viper.GetString("mysql.db")
	charset := viper.GetString("mysql.charset")
	return fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=%s&parseTime=true", usr, pwd, host, db, charset)
}
