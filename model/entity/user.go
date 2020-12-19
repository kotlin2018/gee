package entity

import (
	"gee/global"
	uuid "github.com/satori/go.uuid"
	"time"
)

type User struct {
	global.Model
	UUID  		uuid.UUID			//`json:"uuid" gorm:"comment:用户UUID"`
	Username 	string 				//`json:"username" gorm:"comment:用户登录名"`
	Password 	string				//`json:"password" gorm:"comment:用户密码"`
	Nickname 	string 				//`json:"nickname" gorm:"default:系统用户;comment:用户昵称"`
	Mobile  	string				//`json:"mobile" gorm:"default:null;unique;index;comment:用户手机号"`
	Email 		string				//`json:"email" gorm:"default:null;unique;index;comment:用户邮箱"`
	Avatar   	string				//`json:"avatar" gorm:"default:http://qmplusimg.henrongyi.top/head.png;comment:用户头像"`
	Authority 	Authority    	 	`json:"authority" gorm:"foreignKey:Authority;reference:AuthorityId;comment:用户角色"`
	AuthorityId string 				//`json:"authorityId" gorm:"default:888;comment:用户角色ID"`
}

type Authority struct {
	CreateAt 		time.Time
	UpdateAt		time.Time
	DeletedAt 		time.Time 	`sql:"index"`
	AuthorityId 	string		`json:"authorityId" gorm:"not null;unique;primary_key;comment:角色ID;size:90"`
	AuthorityName 	string		`json:"authorityName" gorm:"comment:角色名"`
	ParentId 		string		`json:"parentId" gorm:"comment:父角色ID"`
	DataAuthorityId []Authority	`json:"dataAuthorityId" gorm:"many2many:data_authority_id"`
	Children 		[]Authority	`json:"children" gorm:"-"`
	//BaseMenus 		[]BaseMenu 	`json:"menus" gorm:"many2many:authority_menus;"`
}