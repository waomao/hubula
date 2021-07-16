package html5

import (
	"fmt"
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
	pathModels  =  "demo/html5"
)

//Get http://localhost:8080/demo/html5
//返回html页面
func (c *IndexController) Get() string {
	c.Ctx.Header("Content-Type", "text/html")
	return "welcome to Go+iris<br><a href='http://www.hubula.com/'>网站首页</a><br>" +
		"<a href='/demo/html5/v1'>V1</a><br>" +
		"<a href='/demo/html5/v2'>V2</a><br>" +
		"<a href='/demo/html5/tianjiahaoma'>添加号码</a><br>" +
		"<a href='/demo/html5/gaoliang'>导航</a><br>" +
		"<a href='/demo/html5/nihao'>导航</a><br>" +
		"<a href='/demo/html5/wohao'>导航</a><br>" +
		"<a href='/demo/html5/wohao/a'>导航</a><br>" +
		"<a href='/demo/html5/wohao/b'>导航</a><br>" +
		"<a href='/demo/html5/xuanze'>选择</a><br>"
}

//Get http://localhost:8080/demo/html5/info

func (c *IndexController) GetV1() mvc.Result {
	path := common.PathAutoOrMobile(c.Ctx.Request(),false,"",conf.MubanIndexPath,pathModels)

	a := c.Ctx.Request().URL.Path
	v := c.Ctx.URLParam("value")
	paths :=c.Ctx.Request().Host + "/" +c.Ctx.Request().URL.String()
	fmt.Print(a, v ,paths)

	return mvc.View{
		Name: path + "v1.html",
		Data: iris.Map{
			"Title":  "请选择管理项目",
			"mssage": "当前路径Passport",
		},
		Layout: path + "layout.html",
	}
}

//首页第一个页面
func (c *IndexController) GetV2() mvc.Result {
	path := common.PathAutoOrMobile(c.Ctx.Request(),false,"",conf.MubanIndexPath,pathModels)

	return mvc.View{
		Name: path + "v2.html",
		Data: iris.Map{
			"Title":  "请选择管理项目",
			"mssage": "当前路径Passport",
		},
		Layout: path + "layout.html",
	}
}

func (c *IndexController) GetTianjiahaoma() mvc.Result {
	path := common.PathAutoOrMobile(c.Ctx.Request(),false,"",conf.MubanIndexPath,pathModels)

	return mvc.View{
		Name: path + "tianjiahaoma.html",
		Data: iris.Map{
			"Title":  "养眼-添加号码",
		},
		Layout: path + "layout.html",
	}
}

func (c *IndexController) GetGaoliang() mvc.Result {
	path := common.PathAutoOrMobile(c.Ctx.Request(),false,"",conf.MubanIndexPath,pathModels)

	return mvc.View{
		Name: path + "indexxx.html",
		Data: iris.Map{
			"Title":  "请选择管理yangyan项目",
		},
		Layout: path + "layoutdaohang.html",
	}
}

func (c *IndexController) GetNihao() mvc.Result {
	path := common.PathAutoOrMobile(c.Ctx.Request(),false,"",conf.MubanIndexPath,pathModels)

	return mvc.View{
		Name: path + "indexxx.html",
		Data: iris.Map{
			"Title":  "请选择管理yangyan项目",
		},
		Layout: path + "layoutdaohang.html",
	}
}

func (c *IndexController) GetWohao() mvc.Result {
	path := common.PathAutoOrMobile(c.Ctx.Request(),false,"",conf.MubanIndexPath,pathModels)

	return mvc.View{
		Name: path + "indexxx.html",
		Data: iris.Map{
			"Title":  "请选择管理yangyan项目",
		},
		Layout: path + "layoutdaohang.html",
	}
}

func (c *IndexController) GetWohaoA() mvc.Result {
	path := common.PathAutoOrMobile(c.Ctx.Request(),false,"",conf.MubanIndexPath,pathModels)

	return mvc.View{
		Name: path + "indexxx.html",
		Data: iris.Map{
			"Title":  "请选择管理yangyan项目",
		},
		Layout: path + "layoutdaohang.html",
	}
}

func (c *IndexController) GetWohaoB() mvc.Result {
	path := common.PathAutoOrMobile(c.Ctx.Request(),false,"",conf.MubanIndexPath,pathModels)

	return mvc.View{
		Name: path + "indexxx.html",
		Data: iris.Map{
			"Title":  "请选择管理yangyan项目",
		},
		Layout: path + "layoutdaohang.html",
	}
}

func (c *IndexController) GetXuanze() mvc.Result {
	path := common.PathAutoOrMobile(c.Ctx.Request(),false,"",conf.MubanIndexPath,pathModels)

	//<?php
	/**
	 *加入搜索条件
	 */
	//$where ="1";
	//$year_array = array(1=>'2015',2=>'2014');
	//$ctype_array = array(1=>'0',2=>'1');
	//$colors_array = array(1=>'0',2=>'1',3=>'2',4=>'3',5=>'4',6=>'5');
	//$lengths_array = array(1=>'0',2=>'1',3=>'2',4=>'3',5=>'4',6=>'5',7=>'6');
	//$micronaire_array = array(1=>'0',2=>'1',3=>'2',4=>'3',5=>'4');

	//if(isset($year)&&($year!=0)) $where .= " AND year=".$year_array[$year];
	//if(isset($ctype)&&($ctype!=0)) $where .= " AND ctype=".$ctype_array[$ctype];
	//if(isset($colors)&&($colors!=0)) $where .= " AND colors=".$colors_array[$colors];
	//if(isset($lengths)&&($lengths!=0)) $where .= " AND lengths=".$lengths_array[$lengths];
	//if(isset($micronaire)&&($micronaire!=0)) $where .= " AND micronaire=".$micronaire_array[$micronaire];
	/**
	 *加入搜索条件
	 */
	//?>

	return mvc.View{
		Name: path + "xuanze.html",
		Data: iris.Map{
			"Title":  "请选择管理yangyan项目",
		},
		Layout: path + "layout.html",
	}
}








//func (c *ManagerYangyanController) GetYyhaoma() mvc.Result {
//	//rs := make(map[string]interface{})
//	haoma := c.Ctx.URLParam("haoma")
//	sql := "SELECT * FROM `yy_haoma` WHERE `haoma` LIKE '%" + haoma + "%' "
//	qq := c.ServiceYyHaoma.GetWhere(sql)
//	var resList []interface{}
//	for _, shop := range qq {
//		resList = append(resList, shop.ShopToRespDesc())
//	}
//	return mvc.Response{
//		Object: map[string]interface{}{
//			"aa": 0,
//			"bb": resList,
//		},
//	}
//}

//func (c *ManagerYangyanController) PostTianjiahaoma() mvc.Result {
//	yyxinxi := &models.YyHaoma{}
//
//	yhaoma,_ := (strconv.Atoi(c.Ctx.PostValue("yyhaoma")))
//	ydengji,_ := (strconv.Atoi(c.Ctx.PostValue("yydengji")))
//	vipdengj,_ := (strconv.Atoi(c.Ctx.PostValue("vipdengji")))
//	nianf,_ := (strconv.Atoi(c.Ctx.PostValue("nianfei")))
//	chengzhangz,_ := (strconv.Atoi(c.Ctx.PostValue("chengzhangzhi")))
//	vipt,_ := (strconv.Atoi(c.Ctx.PostValue("viptime")))
//	juew,_ := (strconv.Atoi(c.Ctx.PostValue("juewei")))
//	renz,_ := (strconv.Atoi(c.Ctx.PostValue("renzheng")))
//	jiag,_ := (strconv.Atoi(c.Ctx.PostValue("jiage")))
//
//	yyxinxi.Haoma = int64(yhaoma)
//	yyxinxi.Dengji = int64(ydengji)
//	yyxinxi.VipDengji = int64(vipdengj)
//	yyxinxi.IsNianfei = nianf
//	yyxinxi.Chengzhangzh = int64(chengzhangz)
//	yyxinxi.Viptime = vipt
//	yyxinxi.Juewei = juew
//	yyxinxi.IsZhengjian = renz
//	yyxinxi.Jiage = jiag
//
//	yyxinxi.HaomaJiwei = 1
//	yyxinxi.IsYincang= 1
//	yyxinxi.IsTuijian= 1
//	yyxinxi.Status= 1
//	yyxinxi.Inputtime= 1
//	yyxinxi.Updatetime= 1
//	yyxinxi.Youxiaoqi= 1
//	yyxinxi.Selltime= 1
//	yyxinxi.Userid= 1
//	yyxinxi.Username= "444"
//	yyxinxi.Ip= "555"
//	yyxinxi.Hits= 1
//
//	yhaoid, err := c.ServiceYyHaoma.Createid(yyxinxi)
//	if err != nil {
//		//sucNum = "失败"
//		c.Ctx.WriteString(err.Error())
//		c.Ctx.HTML(fmt.Sprintf("\n插入yy号码失败"))
//	} else {
//		// 成功导入数据库，还需要导入到缓存中
//		//sucNum = "成功"
//		c.Ctx.HTML(fmt.Sprintf("\n插入yy号码成功"))
//	}
//
//	c.Ctx.HTML(fmt.Sprintf("\n插入yy号码 %#v，<a href='/cr'>继续</a>", "成功"))
//	return mvc.Response{
//		Object: map[string]interface{}{
//			"bb": yyxinxi,
//			"aa": yhaoid,
//		},
//	}
//
//}
//
//
//func (c *ManagerYangyanController) GetHaomachaxun() mvc.Result {
//	//sql语句初始化
//	condition := "SELECT * FROM `yy_haoma` WHERE 1"
//
//	//如果有传入搜索的号码 if 不等于空
//	if(c.Ctx.URLParam("haoma") != ""){
//		condition += " AND `haoma` LIKE '%" + c.Ctx.URLParam("haoma") + "%'"
//	}
//
//	//if GET['taocan'] > 0
//	if(c.Ctx.URLParam("taocan") > "0" ){
//		condition += " AND taocan=" +  c.Ctx.URLParam("taocan")
//	}
//
//	//if GET['yunyingshang'] > 0
//	if(c.Ctx.URLParam("yunyingshang") > "0" ){
//		condition += " AND yunyingshang=" +  c.Ctx.URLParam("yunyingshang")
//	}
//
//	//if GET['status'] > 0
//	if(c.Ctx.URLParam("status") > "0" ){
//		condition += " AND status=" +  c.Ctx.URLParam("status")
//	}
//
//	//if GET['tuijian'] > 0
//	if(c.Ctx.URLParam("tuijian") > "0" ){
//		condition += " AND tuijian=" +  c.Ctx.URLParam("tuijian")
//	}
//
//	//sql----------------------------------------------------------------
//	// haoduan >0  " AND ".['sql'];
//	if (c.Ctx.URLParam("haoduan") !="" && common.SqlHaoduanSJH(c.Ctx.URLParam("haoduan"))["sql"] != ""){
//		condition += " AND " +  common.SqlHaoduanSJH(c.Ctx.URLParam("haoduan"))["sql"]
//	}
//
//	// tezheng
//	if (c.Ctx.URLParam("tezheng") !="" && common.SqlTeZheng(c.Ctx.URLParam("tezheng"))["sql"] != ""){
//		condition += " AND " +  common.SqlTeZheng(c.Ctx.URLParam("tezheng"))["sql"]
//	}
//
//	// shengri $condition .= " AND ".$array_sql_shengri[$_GET['shengri']]['sql'];
//	if (c.Ctx.URLParam("shengri") !="" && common.SqlShengRi(c.Ctx.URLParam("shengri"))["sql"] != ""){
//		condition += " AND " +  common.SqlShengRi(c.Ctx.URLParam("shengri"))["sql"]
//	}
//
//	// niandai $condition .= " AND ".$array_sql_niandai[$_GET['niandai']]['sql'];
//	if (c.Ctx.URLParam("niandai") !="" && common.SqlNianDai(c.Ctx.URLParam("niandai"))["sql"] != ""){
//		condition += " AND " +  common.SqlNianDai(c.Ctx.URLParam("niandai"))["sql"]
//	}
//
//	// jili  $condition .= " AND ".$array_sql_jili[$_GET['jili']]['sql'];
//	if (c.Ctx.URLParam("jili") !="" && common.SqlJiLiHao(c.Ctx.URLParam("jili"))["sql"] != ""){
//		condition += " AND " +  common.SqlJiLiHao(c.Ctx.URLParam("jili"))["sql"]
//	}
//
//	// jiage  $condition .= " AND ".$array_sql_kafei[$_GET['jiage']]['sql'];
//	if (c.Ctx.URLParam("jiage") !="" && common.SqlJiaGe(c.Ctx.URLParam("jiage"))["sql"] != ""){
//		condition += " AND " +  common.SqlJiaGe(c.Ctx.URLParam("jiage"))["sql"]
//	}
//
//	// jiaoduo $condition .= " AND ".$array_sql_jiaoduo[$_GET['jiaoduo']]['sql'];
//	if (c.Ctx.URLParam("jiaoduo") !="" && common.SqlJiaoDuo(c.Ctx.URLParam("jiaoduo"))["sql"] != ""){
//		condition += " AND " +  common.SqlJiaoDuo(c.Ctx.URLParam("jiaoduo"))["sql"]
//	}
//
//	// buhan $condition .= " AND ".$array_sql_buhan[$_GET['buhan']]['sql'];
//	if (c.Ctx.URLParam("buhan") !="" && common.SqlBuHan(c.Ctx.URLParam("buhan"))["sql"] != ""){
//		condition += " AND " +  common.SqlBuHan(c.Ctx.URLParam("buhan"))["sql"]
//	}
//
//	//整体11位手机号----------------------------------------------------------------
//	for i := 1; i <= 11; i++ {
//		if c.Ctx.URLParam(fmt.Sprint("kuang%v", i)) != "" {
//			condition += " AND floor(mid(haoma," + strconv.Itoa(i) + ",1)) = " + c.Ctx.URLParam(fmt.Sprint("kuang%v", i))
//		}
//	}
//
//	//$condition = substr($condition, 4);
//
//	//排序 if !$_GET['order']
//	if (c.Ctx.URLParam("order") !="" && common.SqlOrder(c.Ctx.URLParam("order"))["sql"] != ""){
//		condition += common.SqlOrder(c.Ctx.URLParam("order"))["sql"]
//	}
//	//?haoma=&taocan=&yunyingshang=&status=&tuijian=&haoduan=&tezheng= &shengri=&niandai=
//	// &jili=&jiage=&jiaoduo=&buhan=&order=
//	//分页 isset($_GET['page']) && intval($_GET['page']) ? intval($_GET['page']) : 1
//	page := 1
//	if c.Ctx.URLParam("page") !="" {
//		pages, _ := strconv.Atoi(c.Ctx.URLParam("page"))
//		page = pages
//	}
//	fmt.Println(page)
//
//	a:=fmt.Sprintf("sql%v%v,%v,pages=%v",condition,page,20)
//
//	qq := c.ServiceYyHaoma.GetWhere(condition)
//	var resList []interface{}
//	for _, shop := range qq {
//		resList = append(resList, shop.ShopToRespDesc())
//	}
//
//	//$infos = $this->db->listinfo($condition,$order,$page, $pages = '20');
//	//$pages = $this->db->pages;
//	//include $this->admin_tpl('haoma_list');
//
//
//	//yunyingshang = getcache('yunyingshang', 'sjh');
//	//taocan = getcache('taocan', 'sjh');
//	//sql_tezheng = getcache('sql_tezheng', '../sqls/'.ROUTE_M);//号码特征
//	//sql_shengri = getcache('sql_shengri', '../sqls/'.ROUTE_M);//生日
//	//sql_niandai = getcache('sql_niandai', '../sqls/'.ROUTE_M);//年代号
//	//sql_jili = getcache('sql_jili', '../sqls/'.ROUTE_M);//吉利号
//	//sql_kafei = getcache('sql_kafei', '../sqls/'.ROUTE_M);//价格
//	//sql_jiaoduo = getcache('sql_jiaoduo', '../sqls/'.ROUTE_M);//较多
//	//ql_buhan = getcache('sql_buhan', '../sqls/'.ROUTE_M);//不包含
//	//sql_haoduan = getcache('sql_haoduan', '../sqls/'.ROUTE_M);//号段
//	//sql_order = getcache('sql_order', '../sqls/'.ROUTE_M);//排序
//	//
//	//sfields = array('手机号','会员名');
//	//dfields = array('haoma', 'username');
//	//status_array = array(99=>'出售中',2=>'已售出',1=>'隐藏');
//	//tuijian_array = array(1=>'推荐',0=>'不推荐');
//
//	ha := common.SqlShengRi("15")["sql"]
//	has := common.SqlShengRi("12")["sql"]
//
//	return mvc.Response{
//		Object: map[string]interface{}{
//			"bb": resList,
//			"aa": a,
//			"ha":ha,
//			"has":has,
//		},
//	}
//
//}


