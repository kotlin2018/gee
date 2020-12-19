package request

import (
	"gee/global"
	"github.com/dgrijalva/jwt-go"
	uuid "github.com/satori/go.uuid"
)

// JWT 鉴权用的结构体
//type CustomClaims struct {
//	UUID 	    uuid.UUID
//	ID 			uint
//	Username    string      // json:"username"
//	Nickname  	string
//	Mobile 	    string      // json:"mobile"
//	AuthorityId string
//	BufferTime  int64
//	jwt.StandardClaims
//}

type CustomClaims struct {
	UUID 	    uuid.UUID
	Username    string      // json:"username"
	Mobile 	    string      // json:"mobile"
	BufferTime  int64
	jwt.StandardClaims
}


// 转换请求参数中的 pageNum，pageSize
func Page(pageNum,pageSize int)(offset,limit int){
	if pageNum <= 0 || pageSize <=0{
		pageNum = 1
		pageSize = 1
	}else if pageSize >= 100 {
		pageSize = global.Config.System.PageSize
	}
	limit = pageSize
	offset = pageSize * (pageNum-1)
	return
}

// 注册请求参数，要存储到数据库表中的
type Register struct {
	Username    string   	//`json:"username"`
	Password 	string		//`json:"password"`
	NickName 	string		//`json:"nickName" gorm:"default:'admin'"`
	Avatar		string		//`json:"avatar" gorm:"default:'http://www.henrongyi.top/avatar/lufu.jpg'"` //头像
	AuthorityId string		//`json:"authorityId" gorm:"default:888"`			// 角色权限ID
	Mobile 		string 		//`json:"mobile" gorm:"default:null;not null"`
	Email 		string		//`json:"email" gorm:"default:null"`
}

// 登陆请求参数，不存储到数据库表中
type Login struct {
	Username  string	//`json:"username"`
	Password  string 	//`json:"password"`
	Captcha   string 	//`json:"captcha"`
	CaptchaId string	//`json:"captchaId"`
	Mobile 	  string 	//`json:"mobile"`
	Email 	  string    //`json:"email"`
}

type ChangePassword struct {
	Username 	string //`json:"username"`
	Password 	string //`json:"password"`
	NewPassword string //`json:"newPassword"`
}

type SetUserAuth struct {
	UUID 		uuid.UUID	//`json:"uuid"`
	AuthorityId string 		//`json:"authorityId"`
}