package common

import (
	"html/template"
	"strconv"
	"strings"
	"time"
)

/**
* 返回总记录条数,总页数,当前页面数据,分页html
* 根据分页选项,生成分页连接 下面是一个实例:
    func (this *MainController) Test() {
       var po util.PageOptions
       po.EnablePreNexLink = true
       po.EnableFirstLastLink = true
       po.LinkItemCount = 7
       po.TableName = "help_topic"
       cp, _ := this.GetInt("pno")
       po.Currentpage = int(cp)
       _,_,_ pager := util.GetPagerLinks(&po, this.Ctx)
       this.Data["Email"] = html.HTML(pager)
       this.TplNames = "test.html"
   }
*/

//GetAll 列表查询传入的结构体
type SqlWhere struct {
	TableName           string //表名  -----------------[必填]
	Conditions          map[string]interface{} //条件
	Fields              []string //字段
	OrderBy             string //排序
	Currentpage         int64    //当前页 ,默认1 每次分页,必须在前台设置新的页数,不设置始终默认1.在控制器中使用方式:cp, _ := this.GetInt("pno")   po.Currentpage = int(cp)
	PageSize            int64    //页面大小,默认20
	Uri                 string
}

//GetAll 列表查询返回的结构体
//返回总记录条数,总页数,以及当前请求的数据RawSeter
type SqlReturn struct {
	Page        int64   //当前页		`json:"page"`
	PageSize    int64   //每页条数
	TotalCount  int64   //总条数
	TotalPage   int64   //总页码
	Data        []interface{}        `json:"data"`  //数据
	Href        string //A标签的链接地址  ---------[不需要设置]
	Str         template.HTML  //分页
}

//分页标签信息
type PageOptions struct {
	FirstPageText       string //首页文字  默认"首页"
	LastPageText        string //尾页文字  默认"尾页"

	PrePageText         string //上一页文字 默认"上一页"
	NextPageText        string //下一页文字 默认"下一页"
	Currentpage         int64    //当前页 ,默认1 每次分页,必须在前台设置新的页数,不设置始终默认1.在控制器中使用方式:cp, _ := this.GetInt("pno")   po.Currentpage = int(cp)

	LinkItemCount       int64    //生成A标签的个数 默认10个
	PageSize            int64    //页面大小,默认20
	TotalPage           int64   //总页码
	ParamName           string //参数名称  默认是page

	Href                string //A标签的链接地址  ---------[不需要设置]

	EnableFirstLastLink bool   //是否启用首尾连接 默认false 建议开启
	EnablePreNexLink    bool   //是否启用上一页,下一页连接 默认false 建议开启
}

//设置默认值
func SetDefault(po *PageOptions) *PageOptions {
	if len(po.FirstPageText) <= 0 {
		po.FirstPageText = "首页"
	}
	if len(po.LastPageText) <= 0 {
		po.LastPageText = "尾页"
	}
	if len(po.PrePageText) <= 0 {
		po.PrePageText = "&lt; 上一页"
	}
	if len(po.NextPageText) <= 0 {
		po.NextPageText = "下一页 &gt;"
	}
	if po.Currentpage >= po.TotalPage {
		po.Currentpage = po.TotalPage
	}
	if po.Currentpage <= 1 {
		po.Currentpage = 1
	}
	if po.LinkItemCount == 0 {
		po.LinkItemCount = 10
	}
	if po.PageSize == 0 {
		po.PageSize = 20
	}
	if len(po.ParamName) <= 0 {
		po.ParamName = "page"
	}
	po.EnableFirstLastLink = false
	po.EnablePreNexLink = true
	return po
}

//第一步 传入总条数 totalCount 返回分页标签信息 *PageOptions
func GetPages(sqlwhere *SqlWhere,totalCount int64) (po *PageOptions) {
	po = new(PageOptions)

	//当前页
	if sqlwhere.Currentpage <= 1 {
		sqlwhere.Currentpage = 1
	}

	//每页数量
	if sqlwhere.PageSize == 0 {
		sqlwhere.PageSize = 20
	}

	//总页数
	totalPage  := int64(0)

	if totalCount <= (sqlwhere.PageSize) {
		totalPage = 1
	} else if totalCount > (sqlwhere.PageSize) {
		temp := totalCount / (sqlwhere.PageSize)
		if (totalCount % (sqlwhere.PageSize)) != 0 {
			temp = temp + 1
		}
		totalPage = temp
	}

	//当前页
	po.Currentpage = sqlwhere.Currentpage
	//每页数量
	po.PageSize = sqlwhere.PageSize
	//总页数
	po.TotalPage = totalPage

	//对比并设置默认值
	po = SetDefault(po)
	po = DealUri(po,sqlwhere.Uri)
	return
}

/**
 * 处理url,目的是保存参数
 * ParamName  string //参数名称  默认是pno
 */
//第二步 处理 PageOptions.Href A标签的链接地址
func DealUri(po *PageOptions,uri string) *PageOptions{
	var rs string

	//如果url ？ 后面带参数
	if strings.Contains(uri, "?") {
		arr := strings.Split(uri, "?")
		rs = arr[0] + "?" + po.ParamName + "time=" + strconv.Itoa(time.Now().Second())
		arr2 := strings.Split(arr[1], "&")
		for _, v := range arr2 {
			if !strings.Contains(v, po.ParamName) {
				rs += "&" + v
			}
		}
	} else {
		rs = uri + "?" + po.ParamName + "time=" + strconv.Itoa(time.Now().Second())
	}
	po.Href = rs
	return po
}

//最后 组装分页标签
func H(po *PageOptions) string {
	str := ""
	if po.TotalPage <= po.LinkItemCount {
		str = fun1(po, po.TotalPage) //显示完全  12345678910
	} else if po.TotalPage > po.LinkItemCount {
		if po.Currentpage < po.LinkItemCount {
			str = fun2(po, po.TotalPage) //123456789...200
		} else {
			if po.Currentpage+po.LinkItemCount < po.TotalPage {
				str = fun3(po, po.TotalPage)
			} else {
				str = fun4(po, po.TotalPage)
			}
		}
	}
	return str
}


/**
 * 1...197 198 199 200
 */
func fun4(po *PageOptions, totalpages int64) string {
	var rs = ""
	rs += getHeader(po, totalpages)
	rs += "<a href='" + po.Href + "&" + po.ParamName + "=" + strconv.Itoa(1) + "'>" + strconv.Itoa(1) + "</a>"
	rs += "<a href=''>...</a>"
	for i := totalpages - po.LinkItemCount; i <= totalpages; i++ {
		if po.Currentpage != i {
			rs += "<a href='" + po.Href + "&" + po.ParamName + "=" + strconv.Itoa(int(i)) + "'>" + strconv.Itoa(int(i)) + "</a>"
		} else {
			//rs += "<span class=\"current\">" + strconv.Itoa(int(i)) + "</span>"
			//<span class=\"fk\"><i class=\"pic\"></i></span>
			rs += "<strong><span class=\"pc\">" + strconv.Itoa(int(i)) + "</span></strong>"
		}
	}
	rs += getFooter(po, totalpages)
	return rs

}

/**
 * 1...6 7 8 9 10 11 12  13  14 15... 200
 */
func fun3(po *PageOptions, totalpages int64) string {
	var rs = ""
	rs += getHeader(po, totalpages)
	rs += "<a href='" + po.Href + "&" + po.ParamName + "=" + strconv.Itoa(1) + "'>" + strconv.Itoa(1) + "</a>"
	rs += "<a href=''>...</a>"
	for i := po.Currentpage - po.LinkItemCount/2 + 1; i <= po.Currentpage+po.LinkItemCount/2-1; i++ {
		if po.Currentpage != i {
			rs += "<a href='" + po.Href + "&" + po.ParamName + "=" + strconv.Itoa(int(i)) + "'>" + strconv.Itoa(int(i)) + "</a>"
		} else {
			//rs += "<span class=\"current\">" + strconv.Itoa(int(i)) + "</span>"
			rs += "<strong><span class=\"pc\">" + strconv.Itoa(int(i)) + "</span></strong>"
		}
	}
	rs += "<a href=''>...</a>"
	rs += "<a href='" + po.Href + "&" + po.ParamName + "=" + strconv.Itoa(int(totalpages)) + "'>" + strconv.Itoa(int(totalpages)) + "</a>"
	rs += getFooter(po, totalpages)
	return rs

}

/**
 * totalpages > po.LinkItemCount   po.Currentpage < po.LinkItemCount
 * 123456789...200
 */
func fun2(po *PageOptions, totalpages int64) string {
	var rs = ""
	rs += getHeader(po, totalpages)
	for i := int64(1); i <= po.LinkItemCount+1; i++ {
		if i == po.LinkItemCount {
			rs += "<a href=\"" + po.Href + "&" + po.ParamName + "=" + strconv.Itoa(int(i)) + "\">...</a>"
		} else if i == po.LinkItemCount+1 {
			rs += "<a href=\"" + po.Href + "&" + po.ParamName + "=" + strconv.Itoa(int(totalpages)) + "\">" + strconv.Itoa(int(totalpages)) + "</a>"
		} else {
			if po.Currentpage != i {
				rs += "<a href='" + po.Href + "&" + po.ParamName + "=" + strconv.Itoa(int(i)) + "'>" + strconv.Itoa(int(i)) + "</a>"
			} else {
				//rs += "<span class=\"current\">" + strconv.Itoa(int(i)) + "</span>"
				rs += "<strong><span class=\"pc\">" + strconv.Itoa(int(i)) + "</span></strong>"
			}
		}
	}
	rs += getFooter(po, totalpages)
	return rs
}

/**
 * totalpages <= po.LinkItemCount
 * 显示完全  12345678910
 */
func fun1(po *PageOptions, totalpages int64) string {

	var rs = ""
	rs += getHeader(po, totalpages)
	for i := int64(1); i <= totalpages; i++ {
		if po.Currentpage != i {
			rs += "<a href='" + po.Href + "&" + po.ParamName + "=" + strconv.Itoa(int(i)) + "'>" + strconv.Itoa(int(i)) + "</a>"
		} else {
			//rs += "<span class=\"current\">" + strconv.Itoa(int(i)) + "</span>"
			rs += "<strong><span class=\"pc\">" + strconv.Itoa(int(i)) + "</span></strong>"
		}
	}
	rs += getFooter(po, totalpages)
	return rs
}

/**
 * 头部
 */
func getHeader(po *PageOptions, totalpages int64) string {
	var rs = "<div class='page-inner'>"
	if po.EnableFirstLastLink {
		//当首页,尾页都设定的时候,就显示

		rs += "<a " + judgeDisable(po, totalpages, 0) + " href='" + po.Href + "&" + po.ParamName + "=" + strconv.Itoa(1) + "' class = 'n'>" + po.FirstPageText + "</a>"
	}
	if po.EnablePreNexLink {
		// disabled=\"disabled\"
		var a = po.Currentpage - 1
		if po.Currentpage == 1 {
			a = 1
		}
		rs += "<a " + judgeDisable(po, totalpages, 0) + " href='" + po.Href + "&" + po.ParamName + "=" + strconv.Itoa(int(a)) + "' class = 'n'>" + po.PrePageText + "</a>"
	}
	return rs
}

/**
 * 尾部
 */
func getFooter(po *PageOptions, totalpages int64) string {
	var rs = ""
	if po.EnablePreNexLink {
		var a = po.Currentpage + 1
		if po.Currentpage == totalpages {
			a = totalpages
		}
		rs += "<a " + judgeDisable(po, totalpages, 1) + "  href='" + po.Href + "&" + po.ParamName + "=" + strconv.Itoa(int(a)) + "' class = 'n'>" + po.NextPageText + "</a>"
	}
	if po.EnableFirstLastLink {
		//当首页,尾页都设定的时候,就显示
		rs += "<a " + judgeDisable(po, totalpages, 1) + " href='" + po.Href + "&" + po.ParamName + "=" + strconv.Itoa(int(totalpages)) + "' class = 'n'>" + po.LastPageText + "</a>"
	}
	rs += "</div>"
	return rs
}

/**
 *判断首页尾页  上一页下一页是否能用
 */
func judgeDisable(po *PageOptions, totalpages int64, hf int64) string {
	var rs = ""
	//判断头部
	if hf == 0 {
		if po.Currentpage == 1 {
			rs = "disabled=\"disabled\"  style='pointer-events:none;'"
		}
	} else {
		if po.Currentpage == totalpages {
			rs = "disabled=\"disabled\"  style='pointer-events:none;'"
		}
	}
	return rs
}
