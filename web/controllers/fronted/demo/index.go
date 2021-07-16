package demo

import (
	"github.com/kataras/iris/v12"
	"github.com/waomao/hubula/services"
)

type IndexController struct {
	Ctx         iris.Context
	ServiceADemo services.ADemoService
}

var (
	//pathModels  =  "demo"
)

//Get http://localhost:8080/demo
//返回html页面
func (c *IndexController) Get() string {
	c.Ctx.Header("Content-Type", "text/html")
	return "welcome to Go+iris<br><a href='http://www.hubula.com/'>网站首页</a><br>" +
		"<a href='/demo/info'>返回json</a><br>" +
		"<a href='/demo/boom'>返回404</a><br>" +
		"<a href='/demo/muban'>模板</a><br>" +
		"<a href='/demo/shuju'>数据库操作</a><br>" +
		"<a href='/demo/html5'>Html5演示</a><br>"
}

//Get http://localhost:8080/demo/info
//返回json数据
func (c *IndexController) GetInfo() map[string]interface{} {
	rs := make(map[string]interface{})
	rs["code"] = 200
	rs["msg"] = "成功"
	return rs
}

//Get http://localhost:8080/demo/boom
//返回错误页面
func (c *IndexController) GetBoom() {
	c.Ctx.StatusCode(404)
	return
}