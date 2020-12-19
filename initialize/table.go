package initialize

import (
	"gee/global"
	"gee/model/entity"
)


func init(){
	if global.MongoDB != nil {
		global.Log.Info("Connection mongoDB database success !")
	}
	if global.Redis != nil {
		global.Log.Info("Connection redis database success !")
	}

	// 注册数据库表专用
	if global.DB != nil {
		// 表名为单数
		//global.DB.SingularTable(true)
		// 自动迁移 (自动建表)
		global.DB.AutoMigrate(
			&entity.User{},






		)
		global.Log.Debug("register table success")
	}else {
		global.Log.Error("Connection SQL database failed !")
	}
}


