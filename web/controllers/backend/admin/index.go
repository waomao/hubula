package admin

import (
	"fmt"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"github.com/waomao/hubula/services"
	_ "regexp"
)

type IndexController struct {
	Ctx         iris.Context
	ServiceADemo services.ADemoService
}

var (
//pathModels  =  "admin/"
)

//Get http://localhost:8080/admin
func (c *IndexController) Get() string {
	c.Ctx.Header("Content-Type", "text/html")
	return "welcome to <a href='/demo'>网站DEMO</a>"
}

func (c *IndexController) Gets() mvc.Result {
	a := c.Ctx.Request().URL.Path
	v := c.Ctx.URLParam("name")
	fmt.Print(a, v)
	if v != "admin" {
		err := iris.Map{
			"app":     "pc",
			"status":  404,
			"message": "访问的页面不存在",
		}

		return mvc.View{
			Name: "shared/error.html",
			Data: iris.Map{
				"Err":   err,
				"Title": "Error",
			},
		}
	}

	return mvc.View{
		Name: "admin/index.html",
		Data: iris.Map{
			"Title":    "管理后台",
			"Channel":  "",
			"DataList": 16,
		},
		Layout: "admin/layout1.html",
	}
}
