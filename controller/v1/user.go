package v1

import (
	"fmt"
	"gee/global"
	"gee/middleware"
	"gee/model/entity"
	"gee/model/repository"
	"gee/model/request"
	"gee/model/response"
	"gee/utils"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	//"github.com/thinkeridea/go-extend/exnet"
	"go.uber.org/zap"
	"net/http"
	"time"
)

//var store = base64Captcha.DefaultMemStore
//
//func Captcha(c *gin.Context) {
//	//字符,公式,验证码配置
//	// 生成默认数字的driver
//	driver := base64Captcha.NewDriverDigit(global.Config.Captcha.ImgHeight, global.Config.Captcha.ImgWidth, global.Config.Captcha.KeyLong, 0.7, 80)
//	cp := base64Captcha.NewCaptcha(driver, store)
//	id, b64s, err := cp.Generate()
//	if err != nil {
//		response.FailWithMsg(fmt.Sprintf("获取数据失败，%v", err), c)
//	} else {
//		response.OkDetailed(response.CaptchaResponse{
//			CaptchaId: id,
//			PicPath:   b64s,
//		}, "验证码获取成功", c)
//	}
//}

//// RemoteIp 返回远程客户端的 IP，如 192.168.1.1
//func RemoteIp(req *http.Request) string {
//	remoteAddr := req.RemoteAddr
//	if ip := exnet.ClientPublicIP(req); ip != "" {
//		remoteAddr = ip
//	} else if ip := exnet.ClientIP(req); ip != "" {
//		remoteAddr = ip
//	} else if ip := req.Header.Get("X-Real-IP"); ip != "" {
//		remoteAddr = ip
//	} else if ip = req.Header.Get("X-Forwarded-For"); ip != "" {
//		remoteAddr = ip
//	} else {
//		remoteAddr, _, _ = net.SplitHostPort(remoteAddr)
//	}
//
//	//if remoteAddr == "::1" {
//	//	remoteAddr = "127.0.0.1"
//	//}
//
//	return remoteAddr
//}

func Register(c *gin.Context) {
	var R request.Register
	_ = c.ShouldBindJSON(&R)

	if err := utils.Verify(R, utils.RegisterVerify);err !=nil {
		response.FailWithMsg(err.Error(),c)
		return
	}
	user := &entity.User{Username: R.Username, Nickname: R.NickName, Password: R.Password, Avatar: R.Avatar, AuthorityId: R.AuthorityId,Mobile: R.Mobile,Email: R.Email}
	err, userReturn := repository.Register(*user)
	if err != nil {
		global.Log.Error("注册失败",zap.Any("err",err))
		response.FailWithDetailed(response.ERROR, response.UserResponse{User: userReturn}, fmt.Sprintf("%v", err), c)
	} else {
		response.OkDetailed(response.UserResponse{User: userReturn}, "注册成功", c)
	}
}

func Login(c *gin.Context) {
	var l request.Login
	_ = c.ShouldBindJSON(&l)
	//if err := utils.Verify(l,utils.LoginVerify);err !=nil {
	//	response.FailWithMsg(err.Error(),c)
	//	return
	//}
	//
	//// 验证码正确
	//if store.Verify(l.CaptchaId,l.Captcha,true) {
	//	u := &entity.User{Username: l.Username,Password: l.Password}
	//	if err, user := repository.Login(u);err !=nil {
	//		global.Log.Error("登陆失败! 用户名不存在或者密码错误", zap.Any("err", err))
	//		response.FailWithMsg("用户名不存在或者密码错误", c)
	//	}else {
	//		tokenNext(c,*user)
	//	}
	//}else {
	//	response.FailWithMsg("验证码错误", c)
	//}
	u := &entity.User{Username: l.Username,Password: l.Password,Mobile: l.Mobile,Email: l.Email}
		if err, user := repository.Login(u);err !=nil {
			//global.Log.Error("登陆失败!", zap.Any("err", err))
			//response.FailWithMsg("用户名不存在或者密码错误", c)
			c.JSON(500,gin.H{
				"msg":err.Error(),
			})
		}else {
			tokenNext(c,*user)
		}
}

// 签发 jwt
func tokenNext(c *gin.Context,user entity.User) {
	// 唯一签名 (key)密钥
	//j := []byte(global.Config.JWT.SigningKey)

	claims := request.CustomClaims{
		//UUID: 			user.UUID,    // 注释掉的是可选项
		//ID:   			user.ID,
		Username:       user.Username,
		Mobile: 		user.Mobile,
		//AuthorityId: 	user.AuthorityId,
		BufferTime: 60 * 60 * 24,
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix() - 1000,       //签名生效时间
			ExpiresAt: time.Now().Unix() + 60*60*24*7, //过期时间7天
			Issuer:    "kotlin",					   //签名的发行者
		},
	}
	// 1、生产token
	token, err := middleware.NewJWT().CreateToken(claims)
	if err != nil {
		response.FailWithMsg("获取token失败", c)
		return
	}
	c.JSON(http.StatusOK,gin.H{
		"code":200,
		"msg":"登陆成功",
		"data":response.LoginResponse{
			User: user,
			Token: token,
			ExpiresAt: claims.StandardClaims.ExpiresAt * 1000,
		},
	})
}
