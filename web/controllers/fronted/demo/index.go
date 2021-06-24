package demo

import (
	"fmt"
	"github.com/kataras/golog"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"github.com/waomao/hubula/common"
	"github.com/waomao/hubula/conf"
	"github.com/waomao/hubula/models"
	"github.com/waomao/hubula/services"
	"strconv"
	"time"
)

var (
	pathModels  =  "demo"
)

type IndexController struct {
	Ctx         iris.Context
	ServiceADemo services.ADemoService
}

//Get http://localhost:8080/demo
//返回html页面
func (c *IndexController) Get() string {
	c.Ctx.Header("Content-Type", "text/html")
	return "welcome to Go+iris<br><a href='http://localhost:8080/'>网站首页</a><br>" +
		"<a href='http://localhost:8080/demo/info'>返回json</a><br>" +
		"<a href='http://localhost:8080/demo/boom'>返回404</a><br>" +
		"<a href='http://localhost:8080/demo/muban'>验证模板</a><br>" +
		"<a href='http://localhost:8080/demo/ismobile'>判断设备</a><br>" +
		"<a href='http://localhost:8080/demo/chaandcha'>查询和插入</a><br>" +
		"<a href='http://localhost:8080/demo/gaiandchu'>修改和删除</a><br>" +
		"<a href='http://localhost:8080/demo/route'>路由</a><hr><hr>" +
		"<a href='http://localhost:8080/demo/tianjia'>添加数据</a><br>"+
		"<a href='http://localhost:8080/demo/chaxun'>查询数据</a><br>"+
		"<a href='http://localhost:8080/demo/huishouzhan'>回收站</a><br>"
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

//Get http://localhost:8080/demo/muban
//模板路径是否正常
func (c *IndexController) GetMuban() mvc.Result{
	path := common.PathAutoOrMobile(c.Ctx.Request(),false,"",conf.MubanIndexPath,pathModels)
	//fmt.Println(path)

	return mvc.View{
		Name: path + "muban.html",
		Data: iris.Map{
			"Title":    "模板测试",
		},
		Layout: path + "layout.html",
	}
}

//Get http://localhost:8080/demo/ismobile
//是否是手机等移动设备访问
func (c *IndexController) GetIsmobile() mvc.Result {
	path := common.PathAutoOrMobile(c.Ctx.Request(),false,"",conf.MubanIndexPath,pathModels)

	types := "电脑登录"
	if common.Mobile(c.Ctx.Request()) {
		fmt.Println("手机登录")
		types = "手机登录"
	}

	return mvc.View{
		Name: path + "ismobile.html",
		Data: iris.Map{
			"Title":    "是否是手机等移动设备访问",
			"Types":types,
		},
		Layout: path + "layout.html",
	}
}

//Get http://localhost:8080/demo/chaandcha
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

//Get http://localhost:8080/demo/gaiandchu
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

//Get http://localhost:8080/demo/route
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
	return "welcome to Go+iris <a href='http://localhost:8080/'>网站首页</a>"
}

//Get http://localhost:8080/demo/tianjia
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

//Post http://localhost:8080/demo/tianjia
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

	_, _ = c.Ctx.HTML(fmt.Sprintf("\n插入 %d %#v，<a href='/demo/tianjia'>继续</a>", sucNum, "order"))
	return mvc.Response{
		Object: map[string]interface{}{
			"aa": "sucNum",
		},
	}
}

//Get http://localhost:8080/demo/chaxun
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

	fmt.Print(a,b,e,d,page,pageSize)

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

//Get http://localhost:8080/demo/edit
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


//Posthttp://localhost:8080/demo/cru
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
	_, err2 := c.ServiceADemo.Update(&visitor, nil)
	if err2 != nil {
		xiugai = "失败"
	}else{
		// 成功导入数据库，还需要导入到缓存中
		xiugai = "成功"
	}
	fmt.Println("修改数据",xiugai)


	c.Ctx.HTML(fmt.Sprintf("插入 %d %#v，<a href='/chaxun'>继续</a>", xiugai, visitor))
}

//Get http://localhost:8080/demo/ruandel
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


	c.Ctx.HTML(fmt.Sprintf("软删除数据 %d，<a href='/demo/chaxun'>继续</a>", shanchu))
}

//Get http://localhost:8080/demo/huishouzhan
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

//Get http://localhost:8080/demo/del
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


	c.Ctx.HTML(fmt.Sprintf("删除 %d，<a href='/demo/chaxun'>继续</a>", shanchu))
}