package core

import (
	"fmt"
	"gee/global"
	"gee/initialize"
	"go.uber.org/zap"
	"time"
)
type server interface {
	ListenAndServe() error
}

// 初始化服务
func init() {
	r,_:= initialize.Router() // 初始化路由
	r.Static("/form-generator", "./resource/page")
	address := fmt.Sprintf(":%d", global.Config.System.Port)
	s := initServer(address, r)
	// 保证文本顺序输出
	// In order to ensure that the text order output can be deleted
	time.Sleep(10 * time.Microsecond)
	global.Log.Info("server run success on ", zap.String("address", address))

	fmt.Printf("欢迎使用 %s\n当前版本%s\n默认自动化文档地址%s\n默认前端文件运行地址%s\n后台服务接口%s\n",
		"{Appname}",global.Config.System.AppVersion,"http://127.0.0.1%s/swagger/index.html",
		global.Config.System.Host +":" +address)
	global.Log.Error(s.ListenAndServe().Error())
}
