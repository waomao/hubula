package controllers

/**
 * 首页根目录的Controller
 * http://localhost:8080/
 */
import (
	"github.com/kataras/iris/v12"
	"github.com/waomao/hubula/services"
)

//IndexController s
type IndexController struct {
	Ctx         iris.Context
	ServiceUser services.UserService
}

//Get http://localhost:8080/
func (c *IndexController) Get() string {
	c.Ctx.Header("Content-Type", "text/html")
	return "welcome to Go，<a href='/public/index.html'>开始</a>"
}

//GetInfo http://localhost:8080/info
func (c *IndexController) GetInfo() map[string]interface{} {
	rs := make(map[string]interface{})
	rs["code"] = 0
	rs["msg"] = "123"
	return rs
}
