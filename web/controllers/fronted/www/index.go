package www

import (
	"fmt"
	mobiledetect "github.com/Shaked/gomobiledetect"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"github.com/waomao/hubula/common"
	"github.com/waomao/hubula/conf"
	"github.com/waomao/hubula/services"
)

type IndexController struct {
	Ctx         iris.Context
	ServiceADemo services.ADemoService
}

var (
	pathModels  =  "www"
)

//Get http://localhost:8080/
func (c *IndexController) Get() string {
	c.Ctx.Header("Content-Type", "text/html")
	return "welcome to Go+iris<br><a href='http://www.hubula.com/'>网站首页</a><br>" +
		"<a href='/groupadmin'>后台</a><br>" +
		"<a href='/demo'>演示</a><br>" +
		"<a href='ce'>测试移动类</a><br>" +
		"<a href='zs'>json</a><br>" +
		"<a href='dingdan'>订单</a><hr><hr>" +
		"<a href='biaodan'>表单</a><br>"+
		"<a href='baidu'>百度地图</a><br>"+
		"<a href='ditu1'>地图1</a><br>"+
		"<a href='ditu2'>地图2</a><br>"
}

func (c *IndexController) GetCe(){
	path := common.PathMobile(c.Ctx.Request(),conf.MubanIndexPath,pathModels)
	fmt.Println(path)

	detect := mobiledetect.NewMobileDetect(c.Ctx.Request(), nil)
	fmt.Println(detect)

	requestValue := c.Ctx.Request().URL.Query().Get("r")
	fmt.Println( "is(request)?", detect.Is(requestValue))
	fmt.Println( "isKey(request)?",  detect.IsKey(mobiledetect.IPHONE))
}

//Get json
func (c *IndexController) GetZs() mvc.Result {
	datalist := c.ServiceADemo.CountAll()

	return mvc.Response{
		Object: map[string]interface{}{
			"Path":     "path",
			"DataList": datalist,
		},
	}
}

func (c *IndexController) GetDingdan() mvc.Result {
	path := common.PathMobile(c.Ctx.Request(),conf.MubanIndexPath,pathModels)
	return mvc.View{
		Name: path + "dingdan.html",
		Data: iris.Map{
			"Title":    "map后台",
		},
		Layout: path + "layoutdd.html",
	}
}

func (c *IndexController) GetBiaodan() mvc.Result {
	path := common.PathMobile(c.Ctx.Request(),conf.MubanIndexPath,pathModels)

	return mvc.View{
		Name: path + "js_zsgc.html",
		Data: iris.Map{
			"Title":    "map后台",
		},
		Layout: path + "layoutdd.html",
	}
}

func (c *IndexController) GetBaidu() mvc.Result {
	path := common.PathMobile(c.Ctx.Request(),conf.MubanIndexPath,pathModels)
	return mvc.View{
		Name: path + "baidu.html",
		Data: iris.Map{
			"Title":    "map后台",
		},
		Layout: path + "layoutdd.html",
	}
}


func (c *IndexController) GetDitu1() mvc.Result {
	path := common.PathMobile(c.Ctx.Request(),conf.MubanIndexPath,pathModels)

	return mvc.View{
		Name: path + "ditu1.html",
		Data: iris.Map{
			"Title":    "map后台",
		},
		Layout: path + "layoutditu.html",
	}
}

func (c *IndexController) GetDitu2() mvc.Result {
	path := common.PathMobile(c.Ctx.Request(),conf.MubanIndexPath,pathModels)
	return mvc.View{
		Name: path + "ditu2.html",
		Data: iris.Map{
			"Title":    "map后台",
		},
		Layout: path + "layoutditu.html",
	}
}

//
////全自动事务托管.
//func (c *IndexController) GetTest() {
//	time.LoadLocation("Asia/Shanghai")
//
//	session := datasource.InstanceDbMaster().NewSession()
//	defer session.Close()
//
//	tx, e := session.BeginTrans()
//	if e != nil {
//		panic(e)
//	}
//
//
//	user := &models.Admin{
//		Username:   "李伟",
//		IsDel:    1,
//	}
//
//
//	txsession := tx.Session()
//
//	_, err := txsession.InsertOne(user)
//
//	if err != nil {
//		panic(err)
//		txsession.Rollback()
//	}
//
//	fmt.Println("111")
//	bool, err := txsession.ID(user.Aid).Get(user)
//	if !bool {
//		panic(err)
//		txsession.Rollback()
//	}
//	fmt.Println(user)
//	j, err := json.Marshal(user)
//	if err != nil {
//		panic(err)
//	}
//	fmt.Println(string(j))
//
//	res2, err := txsession.Update(&models.Admin{
//		Username:       "hhh",
//	}, &models.Admin{
//		Aid:        23,
//	})
//	if err != nil || res2 == 0 {
//		panic(err)
//		txsession.Rollback()
//	}
//	txsession.Commit()
//	fmt.Println(res2)
//}