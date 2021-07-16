package shuju

import (
	"fmt"
	"github.com/kataras/golog"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"github.com/waomao/hubula/common"
	"github.com/waomao/hubula/conf"
	"github.com/waomao/hubula/models"
	"github.com/waomao/hubula/services"
	"github.com/waomao/hubula/web/viewsmodels"
	"html/template"
	_ "regexp"
	"strconv"
	"time"
)

type IndexController struct {
	Ctx         iris.Context
	ServiceADemo services.ADemoService
}

var (
	pathModels  =  "demo/shuju"
)

//Get http://localhost:8080/demo/muban
func (c *IndexController) Get() string {
	c.Ctx.Header("Content-Type", "text/html")
	return "welcome to Go+iris<br><a href='http://www.hubula.com/'>网站首页</a><br>" +
		"<a href='/demo'>演示页面</a><br>" +
		"<a href='/demo/shuju/chaandcha'>查询和插入</a><br>" +
		"<a href='/demo/shuju/gaiandchu'>修改和删除</a><br>" +
		"<a href='/demo/shuju/route'>路由</a><hr><hr>" +
		"<a href='/demo/shuju/tianjia'>添加数据</a><br>"+
		"<a href='/demo/shuju/chaxun'>查询数据</a><br>"+
		"<a href='/demo/shuju/huishouzhan'>回收站</a><br>"+
		"<a href='/demo/shuju/usercookie'>读取用户状态</a><br>"+
		"<a href='/demo/shuju/fenlei'>分类</a><br>"+
		"<a href='/demo/shuju/shijian'>时间</a><br>"+
		"<a href='/demo/shuju/yydengji'>yy等级</a><br>"+
		"<a href='/demo/shuju/qqdengji'>qq等级</a><br>"
}

//Get http://localhost:8080/demo/shuju/chaandcha
//查询和插入演示
func (c *IndexController) GetChaandcha() mvc.Result {
	path := common.PathAutoOrMobile(c.Ctx.Request(),false,"",conf.MubanIndexPath,pathModels)
	//总数
	zongshu := c.ServiceADemo.CountAll()
	//fmt.Println(zongshu)

	//所有数据
	page, _ := strconv.Atoi(c.Ctx.URLParam("page"))
	pageSize,_ := strconv.Atoi(c.Ctx.URLParam("pageSize"))

	//查询变量
	query := make(map[string]interface{})

	q := c.Ctx.URLParam("q")
	a := common.BlogKeyWord(q)

	//query["ip=?"] = 1
	//query["is_del=?"] = 1
	query["email REGEXP ? "] = 55
	if len(q) > 0 {
		//username LIKE  ? "] = "CONCAT ('%" + q + "%',%" + q + "%
		//query["username LIKE  ? "] = "%" + q + "%"
		query["username REGEXP ? "] =  a
	}

	sql := common.SqlWhere{
		TableName:   "",
		Conditions:  query,
		Fields:      []string{},
		OrderBy:     "to_id desc",
		Currentpage: int64(page),
		PageSize:    int64(pageSize),
		Uri : c.Ctx.Request().RequestURI,
	}                   //定义sql

	//fmt.Println(sql)

	data, err := c.ServiceADemo.GetAll(&sql)
	//错误输出
	if err != nil {
		fmt.Println("err", err)
		//c.Error(err.Error())
		//return
	}
	//fmt.Println(data)

	//根据Id查询
	by, err:= c.ServiceADemo.GetById(int64(18))
	if err != nil {
		fmt.Println(err)
	}

	//fmt.Println(by)

	//插入
	var sucNum string
	// 将服务器UTC转成北京时间
	uTime := time.Now()
	aa:=uTime.In(time.FixedZone("CST", 8*60*60))

	data2 := &models.ADemo{
		ToId:          0,
		LoginName:     "777",
		Password:      "ooooooooo",
		Vsername:      "fgfgd",
		Mobile:        "14244755",
		Email:         "5555@qq.com",
		GenTime:       aa,
		LoginTime:     aa,
		LastLoginTime: aa,
		Count:         0,
		IsDel:         0,
	}

	_, err = c.ServiceADemo.Create(data2)
	if err != nil {
		sucNum = "失败"
		panic(err)
	} else {
		// 成功导入数据库，还需要导入到缓存中
		sucNum = "成功"
	}
	//fmt.Println("插入数据",sucNum)

	//----------------------------
	return mvc.View{
		Name: path + "chaandcha.html",
		Data: iris.Map{
			"Title":    "插入和查询",
			"Zs":       zongshu,
			"pagings" : data,
			"By" : by,
			"sucNum":sucNum,
		},
		Layout: path + "layout.html",
	}
}

//Get http://localhost:8080/demo/shuju/gaiandchu
//修改和删除演示
func (c *IndexController) GetGaiandchu() mvc.Result {
	path := common.PathAutoOrMobile(c.Ctx.Request(),false,"",conf.MubanIndexPath,pathModels)

	xiugaiid := 1

	if c.Ctx.URLParam("demoid") != "" {
		xiugaiid, _ = strconv.Atoi(c.Ctx.URLParam("demoid"))
	}
	shanchuid := xiugaiid + 1

	fmt.Println(xiugaiid)
	fmt.Println(shanchuid)

	//修改
	var xiugai string
	_, err2 := c.ServiceADemo.Update(&models.ADemo{Id: int64(xiugaiid), Password: "18988",Email:"8888"}, nil)
	if err2 != nil {
		xiugai = "失败"
	}else{
		// 成功导入数据库，还需要导入到缓存中
		xiugai = "成功"
	}
	fmt.Println("修改数据",xiugai)

	//删除
	var shanchu string
	_, err :=c.ServiceADemo.Delete(int64(shanchuid))
	if err != nil {
		shanchu = "失败"
	}else{
		// 成功导入数据库，还需要导入到缓存中
		shanchu = "成功"
	}
	fmt.Println("删除数据",shanchu)

	return mvc.View{
		Name: path + "gaiandchu.html",
		Data: iris.Map{
			"Title":    "删除和修改",
			"xiugai":xiugai,
			"shanchu":shanchu,
			"xiugaiid":xiugaiid,
			"shanchuid":shanchuid,
		},
		Layout: path + "layout.html",
	}
}

//Get http://localhost:8080/demo/shuju/route
//打印路由的标签
func (c *IndexController) GetRoute() string {
	a := c.Ctx.Application().GetRoutesReadOnly()

	for _,only := range a {
		golog.Info(
			"name: %s, path:%s, method:%s",
			only.Name(),only.Path(),only.Method(),
		)
	}

	b := c.Ctx.Application().GetRouteReadOnly("返回html页面")
	fmt.Println(b.Name(),b.Path(),b.Method())

	c.Ctx.Header("Content-Type", "text/html")
	return "welcome to Go+iris <a href='/'>网站首页</a>"
}

//Get http://localhost:8080/demo/shuju/tianjia
//插入数据页面
func (c *IndexController) GetTianjia() mvc.Result {
	path := common.PathAutoOrMobile(c.Ctx.Request(),false,"",conf.MubanIndexPath,pathModels)
	return mvc.View{
		Name: path + "tianjia.html",
		Data: iris.Map{
			"Title":    "添加数据",
		},
		Layout: path + "layout.html",
	}
}

//Post http://localhost:8080/demo/shuju/tianjia
//插入数据处理页面
func (c *IndexController) PostTianjia() mvc.Result {
	// 将服务器UTC转成北京时间
	uTime := time.Now()
	aa:=uTime.In(time.FixedZone("CST", 8*60*60))
	fmt.Print(c.Ctx.PostValue("GenTime"))

	data2 := &models.ADemo{
		ToId:          0,
		LoginName:     c.Ctx.PostValue("LoginName"),
		Password:      c.Ctx.PostValue("Password"),
		Vsername:      c.Ctx.PostValue("Vsername"),
		Mobile:        c.Ctx.PostValue("Mobile"),
		Email:         c.Ctx.PostValue("Email"),
		GenTime:       aa,
		LoginTime:     aa,
		LastLoginTime: aa,
		Count:         0,
		IsDel:         0,
	}

	var sucNum string

	_, err := c.ServiceADemo.Create(data2)
	if err != nil {
		sucNum = "失败"
		panic(err)
	} else {
		// 成功导入数据库，还需要导入到缓存中
		sucNum = "成功"
	}

	_, _ = c.Ctx.HTML(fmt.Sprintf("\n插入 %d %#v，<a href='/demo/shuju/tianjia'>继续</a>", sucNum, "order"))
	return mvc.Response{
		Object: map[string]interface{}{
			"aa": "sucNum",
		},
	}
}

//Get http://localhost:8080/demo/shuju/chaxun
//查询和插入演示
func (c *IndexController) GetChaxun() mvc.Result {
	path := common.PathAutoOrMobile(c.Ctx.Request(),false,"",conf.MubanIndexPath,pathModels)

	a :=  c.Ctx.URLParam("LoginName")
	b :=  c.Ctx.URLParam("Vsername")
	e :=  c.Ctx.URLParam("Mobile")
	d :=  c.Ctx.URLParam("Email")
	page, _ :=  strconv.Atoi(c.Ctx.URLParam("page"))
	pageSize,_ := strconv.Atoi(c.Ctx.URLParam("pageSize"))

	//查询变量
	query := make(map[string]interface{})
	sousuoziduan := make(map[string]interface{})
	sousuoziduan["loginName"] = a
	sousuoziduan["vsername"] = b
	sousuoziduan["mobile"] = e
	sousuoziduan["email"] = d

	if a != ""{
		query["login_name REGEXP ? "] = a
	}
	if b != ""{
		query["vsername REGEXP ? "] = b
	}
	if e != ""{
		query["mobile REGEXP ? "] = e
	}
	if d != ""{
		query["email REGEXP ? "] = d
	}

	//as := common.BlogKeyWord(q)

	//if len(q) > 0 {
	//	//username LIKE  ? "] = "CONCAT ('%" + q + "%',%" + q + "%
	//	//query["username LIKE  ? "] = "%" + q + "%"
	//	query["username REGEXP ? "] =  a
	//}

	sql := common.SqlWhere{
		TableName:   "",
		Conditions:  query,
		Fields:      []string{},
		OrderBy:     "to_id desc",
		Currentpage: int64(page),
		PageSize:    int64(pageSize),
		Uri : c.Ctx.Request().RequestURI,
	}                   //定义sql

	//fmt.Println(sql)


	data, err := c.ServiceADemo.GetAll(&sql)
	//错误输出
	if err != nil {
		fmt.Println("err", err)
		//c.Error(err.Error())
		//return
	}
	//fmt.Println(data)


	return mvc.View{
		Name: path + "chaxun.html",
		Data: iris.Map{
			"Title":    "查询",
			"pagings" : data,
			"sousuo" : sousuoziduan,
		},
		Layout: path + "layout.html",
	}
}

//Get http://localhost:8080/demo/shuju/edit
//编辑页面
func (c *IndexController) GetEdit() mvc.Result {
	path := common.PathAutoOrMobile(c.Ctx.Request(),false,"",conf.MubanIndexPath,pathModels)

	id := c.Ctx.URLParamIntDefault("id", 0)

	//根据Id查询
	by, err:= c.ServiceADemo.GetById(int64(id))
	if err != nil {
		fmt.Println(err)
	}

	return mvc.View{
		Name: path + "giftEdit.html",
		Data: iris.Map{
			"Title":    "查询",
			"by" : by,
		},
		Layout: path + "layout.html",
	}
}

//Posthttp://localhost:8080/demo/shuju/cru
//提交修改
func (c *IndexController) PostCru() {
	visitor := models.ADemo{}
	err := c.Ctx.ReadForm(&visitor)
	if err != nil {
		c.Ctx.StatusCode(iris.StatusInternalServerError)
		c.Ctx.WriteString("填入结构体失败" + err.Error())
	}

	//修改
	var xiugai string
	_, err2 := c.ServiceADemo.Update(&visitor, []string{"pwd", "username", "email"})
	if err2 != nil {
		xiugai = "失败"
	}else{
		// 成功导入数据库，还需要导入到缓存中
		xiugai = "成功"
	}
	fmt.Println("修改数据",xiugai)


	c.Ctx.HTML(fmt.Sprintf("插入 %d %#v，<a href='/demo/shuju/chaxun'>继续</a>", xiugai, visitor))
}

//Get http://localhost:8080/demo/shuju/ruandel
//删除页面
func (c *IndexController) GetRuandel() {
	var shanchu string
	id := c.Ctx.URLParamIntDefault("id", 0)
	fmt.Print(id)
	_, err :=c.ServiceADemo.RuanDelete(int64(id))
	if err != nil {
		shanchu = "失败"
	}else{
		// 成功导入数据库，还需要导入到缓存中
		shanchu = "成功"
	}
	fmt.Println("软删除数据",shanchu)


	c.Ctx.HTML(fmt.Sprintf("软删除数据 %d，<a href='/demo/shuju/chaxun'>继续</a>", shanchu))
}

//Get http://localhost:8080/demo/shuju/huishouzhan
//查询和插入演示
func (c *IndexController) GetHuishouzhan() mvc.Result {
	path := common.PathAutoOrMobile(c.Ctx.Request(),false,"",conf.MubanIndexPath,pathModels)

	page, _ :=  strconv.Atoi(c.Ctx.URLParam("page"))
	pageSize,_ := strconv.Atoi(c.Ctx.URLParam("pageSize"))

	//查询变量
	query := make(map[string]interface{})
	query["is_del"] = "1"

	sql := common.SqlWhere{
		TableName:   "",
		Conditions:  query,
		Fields:      []string{},
		OrderBy:     "to_id desc",
		Currentpage: int64(page),
		PageSize:    int64(pageSize),
		Uri : c.Ctx.Request().RequestURI,
	}

	//fmt.Println(sql)

	data, err := c.ServiceADemo.GetAll(&sql)
	//错误输出
	if err != nil {
		fmt.Println("err", err)
	}

	return mvc.View{
		Name: path + "huishouzhan.html",
		Data: iris.Map{
			"Title":    "回收站",
			"pagings" : data,
		},
		Layout: path + "layout.html",
	}
}

//Get http://localhost:8080/demo/shuju/del
//删除页面
func (c *IndexController) GetDel() {
	var shanchu string
	id := c.Ctx.URLParamIntDefault("id", 0)
	_, err :=c.ServiceADemo.Delete(int64(id))
	if err != nil {
		shanchu = "失败"
	}else{
		// 成功导入数据库，还需要导入到缓存中
		shanchu = "成功"
	}
	fmt.Println("删除数据",shanchu)


	c.Ctx.HTML(fmt.Sprintf("删除 %d，<a href='/demo/shuju/chaxun'>继续</a>", shanchu))
}

//Get http://localhost:8080/demo/shuju/del2
//Get 删除 直接跳转页面
func (c *IndexController) GetDel2() mvc.Result {
	_, _ = c.ServiceADemo.Delete(int64(4))
	return mvc.Response{
		Path: "/all",
	}
}

//Get  http://localhost:8080/demo/shuju/fenlei
func (c *IndexController) GetFenlei() mvc.Result {
	path := common.PathAutoOrMobile(c.Ctx.Request(),false,"",conf.MubanIndexPath,pathModels)

	wuxianjifenlei := c.ServiceADemo.GetMenu(0)
	fmt.Println("8888888888888888", wuxianjifenlei)

	//tree := c.ServiceWebsiteMap.WuXian(data)
	tree := common.Tree(wuxianjifenlei)
	fmt.Println("666666666666666", tree)

	return mvc.View{
		Name: path + "fenlei.html",
		Data: iris.Map{
			"Title":   "所有模块",
			"fenlei":  wuxianjifenlei,
			"tree":    tree,
		},
		Layout: path + "layout.html",
	}
}

//Get  http://localhost:8080/demo/shuju/shijian
func (c *IndexController) GetShijian() mvc.Result {
	path := common.PathMobile(c.Ctx.Request(),conf.MubanIndexPath,pathModels)

	//随机数
	suijishu := common.Random(9999)
	//获取时间戳 1626276224
	t := common.NowUnix()
	//时间格式化
	timeLayout := "2006-01-02 15:04:05"
	// 2020-05-01 00:00:00 +0000 UTC
	shoudong, _ := time.Parse(timeLayout, "2020-05-01 00:00:00")
	// 1588291200
	shoudongt := shoudong.Unix()
	//当前时间 2021-07-14 23:23:44
	shijian := common.FormatFromUnixTime(int64(t))

	miao := time.Now().Unix()
	haomiao := time.Now().UnixNano() / 1e6
	namiao := time.Now().UnixNano()
	natomiao := time.Now().UnixNano() / 1e9

	lingd := common.NextDayDuration()
	fmt.Print(lingd)

	return mvc.View{
		Name: path + "shijian.html",
		Data: iris.Map{
			"Title":     "demo",
			"suijishu":  suijishu,
			"t":         t,
			"shoudong": shoudong,
			"shoudongt": shoudongt,
			"shijian":   shijian,
			"miao":      miao,
			"haomiao":   haomiao,
			"namiao":    namiao,
			"natomiao":  natomiao,
			"ling" : lingd,
		},
		Layout: path + "layout.html",
	}
}

//Get http://localhost:8080/demo/shuju/usercookie
//读取用户cookie
func (c *IndexController) GetUsercookie() mvc.Result {
	path := common.PathAutoOrMobile(c.Ctx.Request(),false,"",conf.MubanIndexPath,pathModels)

	loginuser := viewsmodels.ObjLoginuser{
		Uid:      5,
		Username: fmt.Sprintf("admin-%d", 8),
		Now:      common.NowUnix(),
		Ip:       common.ClientIP(c.Ctx.Request()),
	}

	//设置cookie
	common.SetLoginuser(c.Ctx.ResponseWriter(), &loginuser)

	//读取cookie
	datalist := common.GetLoginUser(c.Ctx.Request())
	return mvc.View{
		Name: path + "usercookie.html",
		Data: iris.Map{
			"Title":    "读取用户状态",
			"Zs":       datalist,
		},
		Layout: path + "layout.html",
	}
}

//
func (c *IndexController) GetQqdengji() mvc.Result{
	path := common.PathAutoOrMobile(c.Ctx.Request(),false,"",conf.MubanIndexPath,pathModels)

	dengji := 47
	dpic := common.QQLeveToSeveral(int64(dengji))
	b := common.QQLevePic(dpic)
	a :=template.HTML(b)
	return mvc.View{
		Name: path + "dengji.html",
		Data: iris.Map{
			"Title":    "QQ等级",
			"leixing":       "QQ",
			"miaosu":       dengji,
			"Pc":       a,
		},
		Layout: path + "layout.html",
	}
}

func (c *IndexController) GetYydengji() mvc.Result{
	path := common.PathAutoOrMobile(c.Ctx.Request(),false,"",conf.MubanIndexPath,pathModels)

	dengji := 3071
	dpic := common.YYLeveToSeveral(int64(dengji))
	b := common.YYLevePic(dpic)
	a :=template.HTML(b)
	return mvc.View{
		Name: path + "dengji.html",
		Data: iris.Map{
			"Title":    "YY等级",
			"leixing":       "YY",
			"miaosu":       dengji,
			"Pc":       a,
		},
		Layout: path + "layout.html",
	}
}
//抖音号
//快手号
//微信号
//陌陌号
//繁星号