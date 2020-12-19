package router

//// 注册路由
//func init(){
//	_,g:= initialize.Router()
//	user(g)    // 注册用户路由
//}
//
//func user(g *gin.RouterGroup) {
//	r := g.Group("user").
//		Use(middleware.JWTAuth())
//	{
//		r.POST("changePassword", v1.ChangePassword)     // 修改密码
//		r.POST("uploadHeaderImg", v1.UploadHeaderImg)   // 上传头像
//		r.POST("getUserList", v1.GetUserList)           // 分页获取用户列表
//		//r.POST("setUserAuthority", v1.SetUserAuthority) // 设置用户权限
//		r.DELETE("deleteUser", v1.DeleteUser)           // 删除用户
//	}
//}