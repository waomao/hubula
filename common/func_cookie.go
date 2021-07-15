package common

import (
	"crypto/md5"
	"fmt"
	"github.com/waomao/hubula/conf"
	"github.com/waomao/hubula/web/viewsmodels"
	"log"
	"net/http"
	"net/url"
	"strconv"
)

// 根据登录用户信息生成加密字符串 签名
func createLoginuserSign(loginuser *viewsmodels.ObjLoginuser) string {
	str := fmt.Sprintf("uid=%d&username=%s&secret=%s&now=%d", loginuser.Uid, loginuser.Username, conf.CookieSecret, loginuser.Now)
	//fmt.Println(str)
	//md5
	sign := fmt.Sprintf("%x", md5.Sum([]byte(str)))
	//fmt.Println(sign)
	return sign
	//return CreateSign(str)
}

// 从cookie中得到当前登录的用户
func GetLoginUser(request *http.Request) *viewsmodels.ObjLoginuser {
	//获取cookie lottery_loginuser
	c, err := request.Cookie("waomao_loginuser")
	if err != nil {
		return nil
	}
	//用url格式解析cookie
	params, err := url.ParseQuery(c.Value)
	if err != nil {
		return nil
	}
	//转换uid
	uid, err := strconv.Atoi(params.Get("uid"))
	if err != nil || uid < 1 {
		return nil
	}
	// Cookie最长使用时长
	now, err := strconv.Atoi(params.Get("now"))
	//当前时间 - cookie 时间 大于 30天 cookie失效
	if err != nil || NowUnix()-now > 86400*30 {
		return nil
	}
	//// IP修改了是不是要重新登录
	//ip := params.Get("ip")
	//if ip != ClientIP(request) {
	//	return nil
	//}
	// 登录信息
	loginuser := &viewsmodels.ObjLoginuser{}
	loginuser.Uid = uid
	loginuser.Username = params.Get("username")
	loginuser.Now = now
	loginuser.Ip = ClientIP(request)
	//签名
	loginuser.Sign = params.Get("sign")
	// if err != nil {
	// 	log.Println("fuc_web GetLoginUser Unmarshal ", err)
	// 	return nil
	// }
	//验证签名
	sign := createLoginuserSign(loginuser)
	if sign != loginuser.Sign {
		log.Println("fuc_web GetLoginUser createLoginuserSign not sign", sign, loginuser.Sign)
		return nil
	}

	return loginuser
}


// 将登录的用户信息设置到cookie中
func SetLoginuser(writer http.ResponseWriter, loginuser *viewsmodels.ObjLoginuser) {
	//如果不存在或id小于1 清理cookie
	if loginuser == nil || loginuser.Uid < 1 {
		c := &http.Cookie{
			Name:  "waomao_loginuser",
			Value: "",
			Path:  "/",
			//过期
			MaxAge: -1,
			Domain: "paodj.com",
		}
		http.SetCookie(writer, c)
		return
	}
	//如果签名为空 则生成签名
	if loginuser.Sign == "" {
		loginuser.Sign = createLoginuserSign(loginuser)
	}
	//构造一个参数
	params := url.Values{}
	//构造好写入cookie
	params.Add("uid", strconv.Itoa(loginuser.Uid))
	params.Add("username", loginuser.Username)
	params.Add("now", strconv.Itoa(loginuser.Now))
	params.Add("ip", loginuser.Ip)
	params.Add("sign", loginuser.Sign)
	c := &http.Cookie{
		Name:   "waomao_loginuser",
		Value:  params.Encode(),
		Path:   "/",
		Domain: "hubula.com",
	}
	http.SetCookie(writer, c)
}


