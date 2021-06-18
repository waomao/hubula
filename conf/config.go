package conf

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

//配置初始化包

//Httpconf 服务地址服务端口
type Httpconf struct {
	Host string `json:"host1"`
	Port int    `json:"port1"`
}

//DriverName 数据库类型
const DriverName = "mysql"

//DbConfig 连接数据库参数的结构体
type DbConfig struct {
	Host      string `json:"host"`
	Port      int    `json:"port"`
	User      string `json:"user"`
	Password  string `json:"password"`
	Database  string `json:"database"`
	Charset   string `json:"charset"`
	IsRunning bool   `json:"isRunning"` //状态 是否正常运行
}

//RdsConfig redis连接参数结构体
type RdsConfig struct {
	Host      string `json:"host"`
	Port      int    `json:"port"`
	User      string `json:"user"`
	Pwd       string `json:"pwd"`
	IsRunning bool   `json:"isRunning"` // 是否正常运行
}

//Amqp ss
type Amqp struct {
	Host   string `json:"host"`
	Port   int    `json:"port"`
	Vhost  string `json:"vhost"`
	User   string `json:"user"`
	Passwd string `json:"passwd"`
}

//读取的配置文件解析到这里
type baseConfig struct {
	Httpconf
	//DbMasterList 系统中所有mysql主库
	DbMasterList []DbConfig `json:"dbConfig"`
	//RdsCacheList 系统中用到的所有redis缓存资源
	RdsCacheList []RdsConfig `json:"rdsConfig"`
	Amqp         `json:"rabbitmqConfig"`
}

var (
	HttpConfig *Httpconf
	//DbMaster1 取出第一个配置结构
	DbMaster *DbConfig

	//RdsCache1 拿到配置参数
	RdsCache  *RdsConfig
	AmqpConfig *Amqp
)

//InitConfig 传入文件名
func InitConfig(filename string) (err error) {
	var (
		content []byte
		conf    baseConfig
	)
	//读取文件
	if content, err = ioutil.ReadFile(filename); err != nil {
		fmt.Println(err)
		return
	}
	//Json Unmarshal：将json字符串解码到相应的数据结构
	if err = json.Unmarshal(content, &conf); err != nil {
		fmt.Println(err)
		return
	}
	//return &conf
	//fmt.Printf("%+v", conf)
	HttpConfig = &conf.Httpconf
	//root:root@tcp(127.0.0.1:3306)/lottery?charset=utf-8
	DbMaster = &conf.DbMasterList[0]

	//RdsCache 拿到配置参数
	RdsCache = &conf.RdsCacheList[0]
	AmqpConfig = &conf.Amqp
	return
}

/*用法
err := conf.InitConfig("E:/go_code/src/waomao.com/web/config.json")
	if err != nil {
		return
	}

	fmt.Println(conf.DbMaster1.User)
*/
