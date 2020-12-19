package initialize

import (
	"gee/config"
	"gee/global"
	"github.com/go-ini/ini"
)

func init(){
	c, err := ini.Load("config/config.ini") //读取配置文件
	if err !=nil {
		panic(err)
	}
	s := new(config.Server)
	c.Section("Server").MapTo(s) //将config.yaml所有配置信息映射到 Server结构体
	global.Config = *s
}

