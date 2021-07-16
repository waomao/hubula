package muban

import (
	"fmt"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"github.com/waomao/hubula/common"
	"github.com/waomao/hubula/conf"
	"github.com/waomao/hubula/services"
	_ "regexp"
)

type IndexController struct {
	Ctx         iris.Context
	ServiceADemo services.ADemoService
}

var (
	pathModels  =  "demo/muban"
)

//Get http://localhost:8080/demo/muban
func (c *IndexController) Get() string {
	c.Ctx.Header("Content-Type", "text/html")
	return "welcome to Go+iris - <a href='http://www.hubula.com/'>网站首页</a> - " +
		"<a href='/demo'>演示页面</a><br>" +
		"<a href='/demo/muban/ceshi'>测试模板</a><br>" +
		"<a href='/demo/muban/ismobile'>判断设备</a><br>" +
		"<a href='/demo/muban/style1'>模板样式 1 </a><br>" +
		"<a href='/demo/muban/style2'>模板样式 2 </a><br>" +
		"<a href='/demo/muban/style3'>模板样式 3 </a><br>"+
		"<a href='/demo/muban/daohang'>导航</a><br>"
}

//Get http://localhost:8080/demo/muban/ceshi
//模板路径是否正常
func (c *IndexController) GetCeshi() mvc.Result{
	path := common.PathAutoOrMobile(c.Ctx.Request(),false,"",conf.MubanIndexPath,pathModels)
	//fmt.Println(path)

	return mvc.View{
		Name: path + "muban.html",
		Data: iris.Map{
			"Title":    "模板测试",
		},
		Layout: path + "layout.html",
	}
}

//Get http://localhost:8080/demo/muban/ismobile
//是否是手机等移动设备访问
func (c *IndexController) GetIsmobile() mvc.Result {
	path := common.PathAutoOrMobile(c.Ctx.Request(),false,"",conf.MubanIndexPath,pathModels)
	fmt.Println(path)
	types := "电脑登录"
	if common.Mobile(c.Ctx.Request()) {
		fmt.Println("手机登录")
		types = "手机登录"
	}

	return mvc.View{
		Name: path + "ismobile.html",
		Data: iris.Map{
			"Title":    "是否是手机等移动设备访问",
			"Types":types,
		},
		Layout: path + "layout.html",
	}
}

//Get http://localhost:8080/demo/muban/style1
func (c *IndexController) GetStyle1() mvc.Result {
	path := common.PathAutoOrMobile(c.Ctx.Request(),false,"",conf.MubanIndexPath,pathModels)

	return mvc.View{
		Name: path + "style1.html",
		Data: iris.Map{
			"Title":    "HTML模板使用方法一",
		},
		Layout: path + "layout1.html",
	}
}

//Get http://localhost:8080/demo/muban/style2
func (c *IndexController) GetStyle2() mvc.Result {
	path := common.PathAutoOrMobile(c.Ctx.Request(),false,"",conf.MubanIndexPath,pathModels)

	return mvc.View{
		Name: path + "style2.html",
		Data: iris.Map{
			"Title":    "HTML模板使用方法二",
		},
		Layout: path + "layout2.html",
	}
}

//Get http://localhost:8080/demo/muban/style3
func (c *IndexController) GetStyle3() mvc.Result {
	path := common.PathAutoOrMobile(c.Ctx.Request(),false,"",conf.MubanIndexPath,pathModels)

	return mvc.View{
		Name: path + "style3.html",
		Data: iris.Map{
			"Title":    "HTML模板使用方法",
		},
		Layout: path + "layout3.html",
	}
}

//
func (c *IndexController) GetDaohang() mvc.Result {
	path := common.PathAutoOrMobile(c.Ctx.Request(),false,"",conf.MubanIndexPath,pathModels)

	datalist := common.GetLoginUser(c.Ctx.Request())
	return mvc.View{
		Name: path + "daohang.html",
		Data: iris.Map{
			"Title":    "读取用户状态",
			"Zs":       datalist,
		},
		Layout: path + "layout_daohang.html",
	}
}