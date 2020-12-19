package initialize

import (
	v1 "gee/controller/v1"
	"gee/global"
	"gee/middleware"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// 初始化总路由
func Router() (*gin.Engine,*gin.RouterGroup){
	r := gin.New()

	// 为用户头像和文件提供静态地址
	//r.StaticFS(global.Config.Local.AvatarPath,http.Dir(global.Config.AvatarPath))
	//r.StaticFS(global.Config.Local.FilePath,http.Dir(global.Config.FilePath))

	r.Use(middleware.LoadTLS())  // 使用https
	global.Log.Debug("use middleware https")

	r.Use(middleware.Cors())// 跨域
	global.Log.Debug("use middleware cors")

	r.Use(middleware.LoggerToFile())
	global.Log.Debug("use middleware logger，使用日志中间件")

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	global.Log.Debug("register swagger handler")

	b := base(r)  //注册基础功能路由 不做鉴权
	b.Use(middleware.JWTAuth())  //之后所有的路由都做鉴权
	global.Log.Debug("use middleware JWTAuth")

	g := b.Group("")  //方便统一添加路由组前缀 多服务器上线使用
	return r,g
}

// 基础功能路由 (不做鉴权)
func base(r *gin.Engine) (R *gin.Engine) {
	{
		r.POST("/register", v1.Register)
		r.POST("/login", v1.Login)
		r.POST("/captcha", v1.Captcha)
	}
	return r
}