package routes

import (
	"github.com/kataras/iris/v12/mvc"
	"github.com/waomao/hubula/bootstrap"
	"github.com/waomao/hubula/services"
	"github.com/waomao/hubula/web/controllers/backend/admin"
	"github.com/waomao/hubula/web/controllers/backend/groupadmin"
	"github.com/waomao/hubula/web/controllers/backend/groupadmin/manager"
	"github.com/waomao/hubula/web/controllers/fronted/blog"
	"github.com/waomao/hubula/web/controllers/fronted/demo"
	"github.com/waomao/hubula/web/controllers/fronted/demo/html5"
	"github.com/waomao/hubula/web/controllers/fronted/demo/muban"
	"github.com/waomao/hubula/web/controllers/fronted/demo/shuju"
	"github.com/waomao/hubula/web/controllers/fronted/errorDiy"
	"github.com/waomao/hubula/web/controllers/fronted/open"
	"github.com/waomao/hubula/web/controllers/fronted/passport"
	"github.com/waomao/hubula/web/controllers/fronted/passport/member"
	"github.com/waomao/hubula/web/controllers/fronted/www"
	"github.com/waomao/hubula/web/controllers/fronted/yangyan"
	"github.com/waomao/hubula/web/middleware"
)

// Configure 和 bootstrap 里定义的一样
func Configure(b *bootstrap.Bootstrapper) {
	//主要是把indexcontrollers放进去 里面定义了很多service
	var (
		ademoService = services.NewADemoService()
	)


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

	//演示 - 模板布局
	dMubanURL := demoURL.Party("/muban")
	dMubanURL.Register(ademoService)
	dMubanURL.Handle(new(muban.IndexController))

	//演示 - 数据库操作
	dShujuURL := demoURL.Party("/shuju")
	dShujuURL.Register(ademoService)
	dShujuURL.Handle(new(shuju.IndexController))

	//演示 - 数据库操作
	dHtml5URL := demoURL.Party("/html5")
	dHtml5URL.Register(ademoService)
	dHtml5URL.Handle(new(html5.IndexController))

	//二级域名示例
	demoErji := mvc.New(b.Party("demo."))
	//groupadminSId.Router.Use(before,Session)
	demoIndex := demoErji.Party("/")
	demoIndex.Register(ademoService)
	demoIndex.Handle(new(demo.IndexController))

	/*或者二级目录 用nginx进行解析
	demo. 的二级域名
	*/

	//虚假后台
	adminURL := mvc.New(b.Party("/admin"))
	adminURL.Router.Use(middleware.BasicAuth)
	adminURL.Register(ademoService)
	adminURL.Handle(new(admin.IndexController))

	//后台
	groupadminURL := mvc.New(b.Party("/groupadmin"))
	groupadminURL.Router.Use(middleware.BasicAuth)
	groupadminURL.Register(ademoService)
	groupadminURL.Handle(new(groupadmin.IndexController))

	//后台首页
	gManagerURL := groupadminURL.Party("/manager")
	gManagerURL.Register(ademoService)
	gManagerURL.Handle(new(manager.IndexController))

	////博客后台
	//gMBlobURL := gManagerURL.Party("/blog")
	//gMBlobURL.Register(adminService)
	//gMBlobURL.Handle(new(groupadmin.ManagerBlogController))
	//
	//gMYangyanURL := gManagerURL.Party("/yangyan")
	//gMYangyanURL.Register(adminService)
	//gMYangyanURL.Handle(new(groupadmin.ManagerYangyanController))

	//前台-------------------------------------------------------------------

	//帐号
	passportURL := mvc.New(b.Party("/passport"))
	passportURL.Register(ademoService)
	passportURL.Handle(new(passport.IndexController))

	pMemberURL := passportURL.Party("/member")
	pMemberURL.Register(ademoService)
	pMemberURL.Handle(new(member.IndexController))

	//假定错误
	errorURL := mvc.New(b.Party("/error"))
	errorURL.Register(ademoService)
	errorURL.Handle(new(errorDiy.IndexController))

	//博客
	blogURL := mvc.New(b.Party("/blog"))
	blogURL.Register(ademoService)
	blogURL.Handle(new(blog.IndexController))

	//养眼
	yangyanURL := mvc.New(b.Party("/yangyan"))
	yangyanURL.Register(ademoService)
	yangyanURL.Handle(new(yangyan.IndexController))

	//订单
	//order := mvc.New(b.Party("/order"))
	//order.Register(userService)
	//order.Handle(new(fronted.OrderController))

	//开放平台
	openURL := mvc.New(b.Party("/open"))
	openURL.Register(ademoService)
	openURL.Handle(new(open.IndexController))
}


//func loginNameHandler(ctx iris.Context){
//	name := ctx.Params().Get("name")
//	println(name)
//	ctx.Next()
//}
//
//func before(ctx iris.Context){
//	println("before")
//	ctx.Next() //继续执行下一个handler，这本例中是mainHandler
//}
//
//func mainHandler(ctx iris.Context){
//	println("mainHandler")
//	ctx.Next()
//}
//
//func after(ctx iris.Context){
//	println("after")
//	ctx.Next()
//}
//