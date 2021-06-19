package routes

import (
	"github.com/kataras/iris/v12/mvc"
	"github.com/waomao/hubula/bootstrap"
	"github.com/waomao/hubula/services"
	"github.com/waomao/hubula/web/controllers/fronted/demo"
	"github.com/waomao/hubula/web/controllers/fronted/www"
)

// Configure 和 bootstrap 里定义的一样
func Configure(b *bootstrap.Bootstrapper) {
	//主要是把indexcontrollers放进去 里面定义了很多service
	ademoService := services.NewADemoService()


	//---------------------------------------------------------------------
	//用mvc创建一个新的路径
	index := mvc.New(b.Party("/"))
	//把 Service 都注册进去
	index.Register(ademoService)
	//路径发给Handle
	index.Handle(new(www.IndexController))

	//演示
	demoURL := mvc.New(b.Party("/demo"))
	demoURL.Register(ademoService)
	demoURL.Handle(new(demo.IndexController))
	demoURL.Controllers[0].GetRoute("Get").Name = "返回html页面"
	demoURL.Controllers[0].GetRoute("GetInfo").Name = "返回json数据"
	demoURL.Controllers[0].GetRoute("GetBoom").Name = "返回错误页面"

	//虚假后台
	//adminURL := mvc.New(b.Party("/admin"))
	//adminURL.Router.Use(middleware.BasicAuth)
	//adminURL.Register(ademoService)
	//adminURL.Handle(new(admin.IndexController))
}