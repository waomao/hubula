package main

import (
	"fmt"
	"github.com/waomao/hubula/bootstrap"
	"github.com/waomao/hubula/conf"
	"github.com/waomao/hubula/web/middleware/identity"
	"github.com/waomao/hubula/web/routes"
)

func newApp() *bootstrap.Bootstrapper {
	// 初始化应用
	app := bootstrap.New("Go+iris", "作者")
	app.Bootstrap()

	//中间件
	app.Configure(identity.Configure, routes.Configure)

	return app
}

func main() {
	// 服务器集群的时候才需要区分这项设置
	// 比如：根据服务器的IP、名称、端口号等，或者运行的参数
	if conf.Configs().App.Port == 8080 {
		conf.RunningCrontabService = true
	}

	app := newApp()
	app.Listen(fmt.Sprintf(":%d", conf.Configs().App.Port))
}