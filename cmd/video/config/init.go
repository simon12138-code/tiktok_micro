package config

import (
	"github.com/fatih/color"
	"github.com/spf13/viper"
)

func InitConfig() {
	//实例化Viper
	v := viper.New()
	//文件路径设置
	v.SetConfigFile("/root/dockerfile/go/go/go_workspace/tiktok_micro/cmd/video/video-settings.yml")
	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}
	//创建配置变量
	serverConfig := ServerConfig{}
	//初始化 将绑定文件的配置信息反序列化到变量当中，完成文件信息到配置变量的转换
	if err := v.Unmarshal(&serverConfig); err != nil {
		panic(err)
	}
	//传递全局变量 调用关系 settings-dev.yaml -> viper -> serverConfig(局部变量) -> global.Settings(全局变量)
	Settings = serverConfig
	color.Blue("11111111", Settings.LogAddress)
}
