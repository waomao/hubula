package blog

import (
	"github.com/kataras/iris/v12"
	"github.com/waomao/hubula/services"
	_ "regexp"
)

type IndexController struct {
	Ctx         iris.Context
	ServiceADemo services.ADemoService
}

//Get http://localhost:8080/blog
func (c *IndexController) Get() {
	c.Ctx.StatusCode(404)
	return
}
