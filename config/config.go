package config

import "time"

// 服务总配置
type Server struct {
	System
	Mysql
	SqLite
	MongoDB
	Redis
	JWT
	Captcha
	Log
	Zap
	Local
	QiNiu
	Map
	Email
}

type System struct {
	UseMultipoint bool      //是否多点登陆
	EnableZip     bool      //是否开启 gzip 支持
	RunMode       string    //应用的模式，默认是 dev 开发模式，pro 为发布模式
	Host          string
	Port          int
	PageSize      int       //每页显示多少条，默认20条
	DatabaseType  string    //数据库类型,默认使用mysql数据库
	RedisOk 	  bool		//是否使用redis数据库，默认不使用
	AppVersion    float64   //版本控制
	SSLHost 	  string    //使用https协议时的主机和端口
}

type Mysql struct {
	Username    string
	Password    string
	HostPort    string
	Database    string
	Config      string
	MaxIdleConn int
	MaxOpenConn int
	LogMode     bool   // 设置日志模式: “true”表示详细的日志，“false”表示没有日志，默认情况下，将只打印错误日志。
}

type SqLite struct {
	Username string
	Password string
	HostPort string
	Config   string
	LogMode  bool
}

type MongoDB struct {
	Username string
	Password string
	HostPort string
	Database string
	Timeout  time.Duration
}

type Redis struct {
	Password    string
	HostPort    string
	Database    int
}

type JWT struct {
	SigningKey string
}

type Captcha struct {
	KeyLong 	int
	ImgWidth 	int
	ImgHeight 	int
}

type Log struct {
	Prefix      string
	LogFile     bool
	Stdout      string
	File        string
	OutPutDir  	string  // log输出目录
	LogSoftLink string
	Module		string  // 项目名称
	FileName    string  // 日志文件名
}

type Zap struct {
	Level 			string  // 默认 info
	Format  		string  // 默认 console
	Prefix  		string  // 项目名 默认 '[Gin-App]'
	Director 		string // 日志输出路径 默认 'log'
	LinkName 		string // 日志输出路径软链接 默认 'latest_log'
	ShowLine 		bool   // 显示行 默认 true
	EncodeLevel 	string // 默认 'LowercaseColorLevelEncoder'
	StacktraceKey 	string // 默认 'stacktrace'
	LogInConsole  	bool   // 默认 true
}

type Local struct {
	UploadPath   string	    // 本地文件上传地址
	DownloadPath string     // 本地文件下载地址
	AvatarPath   string		// 用户头像地址
	FilePath     string     // 本地文件地址
}

type QiNiu struct {
	Zone          string
	Bucket        string
	UseHttps      bool
	UseCdnDomains bool
	AccessKey     string
	SecretKey     string
	ImgPath       string
}

type Map struct {
	AMapKey  string
	AMapUrl  string
	BaiduKey string
	BaiduUrl string
	QQKey    string
	QQIpUrl  string
	QQIpUrl2 string
	QQSearch string
}

type Email struct {
	To   		string
	Port 		int
	From 		string
	Host 		string
	IsSSL 		bool
	Secret 		string
	Nickname 	string
}