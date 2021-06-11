package conf

//RdsConfig redis连接参数结构体
type RdsConfig struct {
	Host      string
	Port      int
	User      string
	Pwd       string
	IsRunning bool // 是否正常运行
}

//RdsCacheList 系统中用到的所有redis缓存资源
var RdsCacheList = []RdsConfig{
	{
		Host:      "127.0.0.1",
		Port:      6379,
		User:      "",
		Pwd:       "",
		IsRunning: true,
	},
}

//RdsCache 拿到配置参数
var RdsCache RdsConfig = RdsCacheList[0]