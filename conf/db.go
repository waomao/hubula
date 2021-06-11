package conf

//mysql 数据库配置信息

//DriverName 数据库类型
const DriverName = "mysql"

//DbConfig 连接数据库的参数结构体
type DbConfig struct {
	Host      string
	Port      int
	User      string
	Pwd       string
	Database  string
	IsRunning bool //状态 是否正常运行
}

//DbMasterList 系统中所有mysql主库
//root:root@tcp(127.0.0.1:3306)/lottery?charset=utf-8
var DbMasterList = []DbConfig{
	{
		Host:      "127.0.0.1",
		Port:      3306,
		User:      "root",
		Pwd:       "root",
		Database:  "test",
		IsRunning: true,
	},
}

//DbMaster 取出第一个配置结构
var DbMaster DbConfig = DbMasterList[0]
