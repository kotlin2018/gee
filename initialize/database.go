package initialize

import (
	"context"
	"gee/config"
	"gee/global"
	"github.com/go-redis/redis"
	"go.uber.org/zap"
	"gorm.io/gorm/logger"
	"os"
	//"github.com/jinzhu/gorm"
	//"github.com/go-sql-driver/mysql"
	//"github.com/jinzhu/gorm/dialects/mysql"
	//"github.com/jinzhu/gorm/dialects/sqlite"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func init(){
	s := config.System{}
	switch s.DatabaseType {
	case "mysql":
		Mysql()
	case "mongodb":
		MongoDB()
	case "sqLite":
		SqLite()
	default:
		Mysql()
	}
	// 启用redis
	if s.RedisOk {
		Redis()
	}
}
// 使用 "github.com/jinzhu/gorm" 这个库
// 初始化Mysql数据库
//func Mysql() {
//	admin := global.Config.Mysql
//	db,err := gorm.Open("mysql",admin.Username+":"+admin.Password+"@("+admin.HostPort+")/"+admin.Database+"?"+admin.Config)
//	if err !=nil {
//		panic(err)
//	}
//	db.DB().SetMaxIdleConns(admin.MaxIdleConn)
//	db.DB().SetMaxOpenConns(admin.MaxOpenConn)
//	db.LogMode(admin.LogMode)
//	global.DB = db
//}
//
//func SqLite() {
//	admin := global.Config.SqLite
//	db, err := gorm.Open("sqlite3", fmt.Sprintf("%s?%s", admin.HostPort, admin.Config))
//	if err !=nil {
//
//	}
//	db.LogMode(admin.LogMode)
//	global.DB = db
//}

// gormConfig 根据配置决定是否开启日志
func gOrmConfig(mod bool) *gorm.Config {
	if mod {
		return &gorm.Config{
			Logger:                                   logger.Default.LogMode(logger.Info),
			DisableForeignKeyConstraintWhenMigrating: true,
		}
	} else {
		return &gorm.Config{
			Logger:                                   logger.Default.LogMode(logger.Silent),
			DisableForeignKeyConstraintWhenMigrating: true,
		}
	}
}

// 使用 "gorm.io/gorm" 这个库
func Mysql () {
	m := global.Config.Mysql
	dsn := m.Username + ":" + m.Password + "@tcp(" + m.HostPort + ")/" + m.Database + "?" + m.Config
	mysqlConfig := mysql.Config{
		DSN:                       dsn,   // DSN data source name
		DefaultStringSize:         191,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据版本自动配置
	}
	if db, err := gorm.Open(mysql.New(mysqlConfig), gOrmConfig(m.LogMode)); err != nil {
		global.Log.Error("MySQL启动异常", zap.Any("err", err))
		os.Exit(0)
	} else {
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(m.MaxIdleConn)
		sqlDB.SetMaxOpenConns(m.MaxOpenConn)
		global.DB = db
	}

}

func SqLite() {
	db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	if err !=nil {
		global.Log.Error("Connect SQLite fail")
	}
	global.DB = db
}

// 初始化MongoDB数据库
func MongoDB() {
	// 1、建立连接
	client,err :=mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://"+global.Config.MongoDB.HostPort).SetConnectTimeout(global.Config.MongoDB.Timeout))
	if err !=nil {
		panic(err)
	}
	// 2、检查连接
	err = client.Ping(context.TODO(), nil)
	if err !=nil {

	}

	// 3、选择数据库
	db := client.Database(global.Config.MongoDB.Database)

	// 4、选择表
	//db.Collection("")
	global.MongoDB = db
}



func Redis(){
	redisCli := global.Config.Redis
	client := redis.NewClient(&redis.Options{
		Addr:     redisCli.HostPort,
		Password: redisCli.Password,
		DB:       redisCli.Database,
	})
	pong, err := client.Ping().Result()
	if err !=nil {
		global.Log.Error(pong)
	}else {

	}
	global.Redis = client
}
