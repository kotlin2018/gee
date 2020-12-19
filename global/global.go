package global

import (
	"gee/config"
	"github.com/go-redis/redis"
	//"github.com/jinzhu/gorm"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"time"
)

var (
	Config 	config.Server
	DB 		*gorm.DB
	MongoDB *mongo.Database
	Redis 	*redis.Client
	Log    	*zap.Logger
)

type Model struct {
	ID        uint 			 //`gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt //`gorm:"index" json:"-"`
}
