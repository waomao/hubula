package member

import (
	"fmt"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"github.com/waomao/hubula/common"
	"github.com/waomao/hubula/conf"
	"github.com/waomao/hubula/services"
	"github.com/waomao/hubula/web/viewsmodels"
	_ "regexp"
)

type IndexController struct {
	Ctx         iris.Context
	ServiceADemo services.ADemoService
}

var (
	pathModels  =  "passport/member"
)

//Get http://localhost:8080/passport/member
func (c *IndexController) Get() {
	c.Ctx.StatusCode(404)
	return
}

//Get http://localhost:8080/passport/member/register
func (c *IndexController) GetRegister() mvc.Result {
	path := common.PathAutoOrMobile(c.Ctx.Request(),false,"",conf.MubanIndexPath,pathModels)

	return mvc.View{
		Name: path + "register.html",
		Data: iris.Map{
			"Title":    "注册管理员账号",
			"DataList": "jq",
		},
		Layout: path + "layout.html",
	}
}

//Get http://localhost:8080/passport/member/login
func (c *IndexController) GetLogin() mvc.Result {
	path := common.PathAutoOrMobile(c.Ctx.Request(),false,"",conf.MubanIndexPath,pathModels)

	return mvc.View{
		Name: path + "login.html",
		Data: iris.Map{
			"Title":    "管理员登录 - 欢迎",
			"DataList": "jq",
		},
		Layout: path + "layout.html",
	}
}

//https://graph.qq.com/oauth2.0/show?which=Login&display=pc&response_type=code&client_id=101487368&redirect_uri=https%3A%2F%2Fpacaio.match.qq.com%2Fqq%2FloginBack%3Fsurl%3Dhttps%3A%2F%2Fwww.qq.com%2F&state=5b481c68e379d
//https://ssl.zc.qq.com/v3/index-chs.html?from=pt


/*
用户登录和退出
基于cookie的用户状态
创建ObjLoginuser登录用户对象
登录用户对象与cookie的读写
cookie的安全效验值，不能被篡改 v
*/

func (c *IndexController) GetLoginDenglu() {
	uid := common.Random(10000)
	loginuser := viewsmodels.ObjLoginuser{
		Uid:      uid,
		Username: fmt.Sprintf("admin-%d", uid),
		Now:      common.NowUnix(),
		Ip:       common.ClientIP(c.Ctx.Request()),
	}
	common.SetLoginuser(c.Ctx.ResponseWriter(), &loginuser)
	common.Redirect(c.Ctx.ResponseWriter(), "http://www.hubula.com/public/index.html?form=denglu"+loginuser.Username)
}

func (c *IndexController) GetLoginTuichu() {
	common.SetLoginuser(c.Ctx.ResponseWriter(), nil)
	common.Redirect(c.Ctx.ResponseWriter(), "http://www.hubula.com/public/index.html?form=tuichu")

}

//GetAll s
func (c *IndexController) GetUser() mvc.Result {
	path := common.PathAutoOrMobile(c.Ctx.Request(),false,"",conf.MubanIndexPath,pathModels)
	common.SetLoginuser(c.Ctx.ResponseWriter(), nil)
	return mvc.View{
		Name: path + "all.html",
		Data: iris.Map{
			"Title":    "退出",
		},
		Layout: path + "layout.html",
	}
}

//GetAll s
func (c *IndexController) GetUserc() mvc.Result {
	path := common.PathAutoOrMobile(c.Ctx.Request(),false,"",conf.MubanIndexPath,pathModels)
	datalist := common.GetLoginUser(c.Ctx.Request())
	fmt.Print(datalist)
	return mvc.View{
		Name: path + "all.html",
		Data: iris.Map{
			"Title":    "退出",
			"Zs":       datalist,
		},
		Layout: path + "layout.html",
	}
}