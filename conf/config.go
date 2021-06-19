package conf

import "time"

//配置文件结构体

type Config struct {
	Title    string
	Runmode string
	CustomLogger bool
	Password_salt string
	Aes_key string
	Stage    map[string]stage `toml:"stage"`
	CookieSession cookieSession
	App      app
	Webvar webvar
	DB       map[string]mysql `toml:"mysql"`
	Redis    map[string]redis
	Releases releases
	Company  Company
	Song    []song
}

type stage struct {
	Admin_load bool
	Http string
}

type cookieSession struct {
	Sessionon bool
	Session_prefix string
	Session_name string
}

type app struct {
	Name string
	Owner string
	Author  string
	Release time.Time
	Port int
	Org     string `toml:"organization"`
	Mark    string
}

type webvar struct {
	Views  string
}

type mysql struct {
	//数据库类型，默认 MYSQL
	DriverName string
	//主机ip
	Host string
	//端口，默认 3306
	Port int
	//数据库登录帐号
	User string
	//登录密码
	Password string
	//数据库名
	Database string
	//向服务器请求连接所使用的字符集，默认：无
	Charset  string
	//MySQL的时区设置
	Loc string
	//显示sql
	ShowSql  bool
	//日志级别
	LogLevel string
	//最大空闲数
	MaxIdleConns int
	//最大连接数
	MaxOpenConns int
	//状态 是否启用
	IsRunning bool
	//查询结果是否自动解析为时间
	ParseTime       bool
	////SetConnMaxLifetime(time.Second * 500) //设置连接超时500秒
	//ConnMaxLifetime int
	////是否启用 SSL 连接模式，默认：MySqlSslMode.None
	//SslMode         string
}

type redis struct {
	Host string
	Port int
}

type releases struct {
	Release []string
	Tags    [][]interface{}
}

type Company struct {
	Name   string
	Detail detail
}

type detail struct {
	Type string
	Addr string
	ICP  string
}

type song struct {
	Name string
	Dur  duration `toml:"duration"`
}

type duration struct {
	time.Duration
}

func (d *duration) UnmarshalText(text []byte) error {
	var err error
	d.Duration, err = time.ParseDuration(string(text))
	return err
}