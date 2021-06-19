package www

import (
	"github.com/kataras/iris/v12"
	"github.com/waomao/hubula/services"
)

var (
//pathModels  =  "www/"
)

type IndexController struct {
	Ctx         iris.Context
	ServiceADemo services.ADemoService
}

//Get http://localhost:8080/
func (c *IndexController) Get() string {
	c.Ctx.Header("Content-Type", "text/html")
	return "welcome to <a href='http://localhost:8080/'>网站DEMO</a>"
}
