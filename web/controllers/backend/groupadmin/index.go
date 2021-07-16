package groupadmin

import (
	"github.com/kataras/iris/v12"
	"github.com/waomao/hubula/services"
	_ "regexp"
)

type IndexController struct {
	Ctx         iris.Context
	ServiceADemo services.ADemoService
}

var (
	//pathModels  =  "groupadmin/"
)

//Get http://localhost:8080/groupadmin
func (c *IndexController) Get() string{
	c.Ctx.Header("Content-Type", "text/html")
	return "管理员welcome to Go+iris <a href='/groupadmin/manager/'>后台首页</a>"
}
