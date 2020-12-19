package response

import (
	"gee/model/entity"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code int         	//json:"code"
	Data interface{} 	//json:"data"
	Msg  string      	//json:"msg"
}

const (
	ERROR   = 7
	SUCCESS = 0
)

func Result(code int, data interface{}, msg string, c *gin.Context) {
	// 开始时间
	c.JSON(http.StatusOK, Response{
		code,
		data,
		msg,
	})
}

func Ok(c *gin.Context) {
	Result(SUCCESS, map[string]interface{}{}, "操作成功", c)
}

func OkWithMsg(message string, c *gin.Context) {
	Result(SUCCESS, map[string]interface{}{}, message, c)
}

func OkWithData(data interface{}, c *gin.Context) {
	Result(SUCCESS, data, "操作成功", c)
}

func OkDetailed(data interface{}, message string, c *gin.Context) {
	Result(SUCCESS, data, message, c)
}

func Fail(c *gin.Context) {
	Result(ERROR, map[string]interface{}{}, "操作失败", c)
}

func FailWithMsg(message string, c *gin.Context) {
	Result(ERROR, map[string]interface{}{}, message, c)
}

func FailWithDetailed(code int, data interface{}, message string, c *gin.Context) {
	Result(code, data, message, c)
}

type UserResponse struct {
	User entity.User 	//`json:"user"`
}

type CaptchaResponse struct {
	CaptchaId string 	// json:"captchaId"
	PicPath   string 	// json:"picPath"
}

type LoginResponse struct {
	User      entity.User 	//`json:"user"`
	Token     string        //`json:"token"`
	ExpiresAt int64         //`json:"expiresAt"`
}