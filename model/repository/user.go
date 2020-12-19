package repository

import (
	"errors"
	"gee/global"
	"gee/model/entity"
	"gee/utils"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

func Register(u entity.User) (err error,user entity.User) {
	if !errors.Is(global.DB.Where("username = ?", u.Username).First(&user).Error, gorm.ErrRecordNotFound) { // 判断用户名是否注册
		return errors.New("用户名已注册"), entity.User{}
	}
	if !errors.Is(global.DB.Where("mobile = ?", u.Mobile).First(&user).Error, gorm.ErrRecordNotFound) {
		return errors.New("手机号已注册"), entity.User{}
	}
	if !errors.Is(global.DB.Where("email = ?", u.Email).First(&user).Error, gorm.ErrRecordNotFound){
		return errors.New("邮箱已注册"), entity.User{}
	}
	// 否则 附加uuid 密码md5简单加密 注册
	u.Password = utils.MD5VByte([]byte(u.Password))
	u.UUID = uuid.NewV4()
	err = global.DB.Create(&u).Error
	return err, u
}

func Login(u *entity.User) (err error,userInter *entity.User) {
	var user entity.User // 这里不能定义成指针，否则报不知名的错误
	u.Password = utils.MD5VByte([]byte(u.Password))
	err = global.DB.Where("password = ?",u.Password).Preload("Authority").First(&user).Error
	if err == nil {
		if err1 := global.DB.Where("username = ?",u.Username).Preload("Authority").First(&user).Error;err1 !=nil { //这里不预加载user表中的authority表也不影响程序正常执行
			goto ERR1
		}
		if err2 := global.DB.Where("mobile = ?",u.Mobile).Preload("Authority").First(&user).Error;err2 !=nil {
			goto ERR2
		}
		if err2 := global.DB.Where("email = ?",u.Email).Preload("Authority").First(&user).Error;err2 !=nil {
			goto ERR3
		}
		return err,&user
	}else {
		return errors.New("密码错误"), &user
	}
	ERR1:
		return errors.New("用户名错误"), &user
	ERR2:
		return errors.New("手机号错误"), &user
	ERR3:
		return errors.New("邮箱错误"), &user

}