package v1

import (
	"fmt"
	"gee/global"
	"gee/model/response"
	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"

)

var store = base64Captcha.DefaultMemStore

func Captcha(c *gin.Context) {
	//字符,公式,验证码配置
	// 生成默认数字的driver
	driver := base64Captcha.NewDriverDigit(global.Config.Captcha.ImgHeight, global.Config.Captcha.ImgWidth, global.Config.Captcha.KeyLong, 0.7, 80)
	cp := base64Captcha.NewCaptcha(driver, store)
	id, b64s, err := cp.Generate()
	if err != nil {
		response.FailWithMsg(fmt.Sprintf("获取数据失败，%v", err), c)
	} else {
		response.OkDetailed(response.CaptchaResponse{
			CaptchaId: id,
			PicPath:   b64s,
		}, "验证码获取成功", c)
	}
}
