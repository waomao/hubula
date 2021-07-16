package manager

import (
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
	pathModels  =  "groupadmin/manager"
)

//Get http://localhost:8080/admin
func (c *IndexController) Get() string {
	c.Ctx.Header("Content-Type", "text/html")
	return "管理员welcome to Go+iris <a href='/groupadmin/'>后台</a>" +
		"<a href='/groupadmin/manager/yangyan'>养眼</a>" +
		"<a href='/groupadmin/manager/blog'>博客</a>" +
		"<a href='/groupadmin/manager/member/register'>注册</a>" +
		"<a href='/groupadmin/manager/member/login'>登录</a>"
}

func (c *IndexController) GetMemberRegister() mvc.Result {
	path := common.PathAutoOrMobile(c.Ctx.Request(),false,"",conf.MubanAdminPath,pathModels)

	return mvc.View{
		Name: path + "register.html",
		Data: iris.Map{
			"Title":    "注册管理员账号",
			"DataList": "jq",
		},
		Layout: path + "layout.html",
	}
}

//登录
func (c *IndexController) GetMemberLogin() mvc.Result {
	path := common.PathAutoOrMobile(c.Ctx.Request(),false,"",conf.MubanAdminPath,pathModels)

	return mvc.View{
		Name: path + "login.html",
		Data: iris.Map{
			"Title":    "管理员登录 - 欢迎",
			"DataList": "jq",
		},
		Layout: path + "layout.html",
	}
}

/*
用户登录和退出
基于cookie的用户状态
创建ObjLoginuser登录用户对象
登录用户对象与cookie的读写
cookie的安全效验值，不能被篡改 v
*/
//验证
//func (c *IndexController) PostMemberLogin() {
//	username := c.Ctx.PostValue("username")
//	password := c.Ctx.PostValue("password")
//	//fmt.Println("username:",username)
//
//	//验证
//	adm, err := c.ServiceAdmin.Auth(username, password)
//	//错误检测
//	if err != nil {
//		fmt.Errorf(err.Error())
//		fmt.Println(err)
//
//		common.SessionSet(c.Ctx,"ISLOGIN", false)
//		c.Ctx.WriteString("账户登录失败，请重新尝试")
//	} else {
//		fmt.Println("登录成功：",adm)
//		//设置Session
//		common.LoginSessionSet(c.Ctx,"jname",adm)
//		c.Ctx.WriteString("账户登录成功 ")
//		//返回
//		//c.Success("操作成功")
//
//		//app.Logger().Info(" 查询信息 path :", path)
//		session := common.Sess.Start(c.Ctx)
//		isLogin, err := session.GetBoolean("ISLOGIN")
//		if err != nil {
//			c.Ctx.WriteString("账户未登录,请先登录 ")
//			return
//		}
//		if isLogin {
//			//app.Logger().Info(" 账户已登录 ")
//			c.Ctx.WriteString("账户已登录"+ "isLogin")
//
//			logi := common.LoginSession(c.Ctx)
//
//			c.Ctx.WriteString("---")
//			c.Ctx.WriteString(logi.Username)
//			common.Redirect(c.Ctx.ResponseWriter(), "http://www.paodj.com/public/index.html?form=tuichu")
//
//		} else {
//			//app.Logger().Info(" 账户未登录 ")
//			c.Ctx.WriteString("账户未登录")
//		}
//
//	}
//
//}

//
//// session 填充
//func (c *IndexController) SSet(sessionname string , sessiona *viewsmodels.AdminSession) {
//
//	SESSION_NAME := sessionname
//
//	//存入 Session
//	str2,_ := json.Marshal(sessiona)
//	session := sess.Start(c.Ctx)
//	//用户名
//	session.Set(SESSION_NAME, string(str2))
//	//登录状态
//	session.Set("ISLOGIN", true)
//
//	//格式化显示
//	buff,_ := json.MarshalIndent(sessiona,"","   ")
//	fmt.Println("str => ?", string(buff))
//}
