package routes

import (
	"github.com/kataras/iris/v12/mvc"
	"github.com/waomao/hubula/bootstrap"
	"github.com/waomao/hubula/services"
	"github.com/waomao/hubula/web/controllers"
)

// Configure 和 bootstrap 里定义的一样
func Configure(b *bootstrap.Bootstrapper) {
	//主要是把indexcontrollers放进去 里面定义了很多service
	userService := services.NewADemoService()


	//用mvc创建一个新的路径
	index := mvc.New(b.Party("/"))
	//把 Service 都注册进去
	index.Register(userService)
	//路径发给Handle
	index.Handle(new(controllers.IndexController))
}
