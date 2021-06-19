package conf

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"path/filepath"
	"sync"
)

var (
	cfg *Config
	once sync.Once
	cfgLock = new(sync.RWMutex)
)

//var Rds = `
//# 全局信息
//title = "TOML格式配置文件示例"
//#开发阶段
//runmode = "dev"
//`

func Configs() *Config {
	once.Do(ReloadConfig)
	cfgLock.RLock()
	defer cfgLock.RUnlock()
	return cfg
}

func ReloadConfig() {
	//config := new(Config)
	//读取文件
	filePath, err := filepath.Abs("./configs/gateway.toml")
	if err != nil {
		fmt.Println("load config error: ", err)
		panic(err)
	}
	//fmt.Printf("parse toml file once. filePath: %s\n", filePath)
	if _ , err := toml.DecodeFile(filePath, &cfg); err != nil {
		fmt.Println("Para config failed: ", err)
		panic(err)
	}

	//读取字符串
	//if _, err := toml.Decode(Rds, &cfg); err != nil {
	//	panic(err)
	//}

	cfgLock.Lock()
	defer cfgLock.Unlock()
	//cfg = cfg
}