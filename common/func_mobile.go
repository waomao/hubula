package common

import (
	"fmt"
	mobiledetect "github.com/Shaked/gomobiledetect"
	"github.com/waomao/hubula/conf"
	"net/http"
)

func Mobile(r *http.Request) bool{
	detect := mobiledetect.NewMobileDetect(r, nil)
	if detect.IsMobile() || detect.IsTablet(){
		if detect.IsMobile() && detect.IsTablet(){
			fmt.Println("Hello, this is Tablet")
		}else {
			fmt.Println("Hello, this is Mobile")
		}
		return true
	}else {
		fmt.Println("Hello, this is Desktop")
		return false
	}
}

//判断pc还是移动设备，进行模板路径拼接，不考虑自适应模板
func PathMobile(r *http.Request,indexoradmin string,model string) string{
	detect := mobiledetect.NewMobileDetect(r, nil)
	if detect.IsMobile() || detect.IsTablet(){
		if detect.IsMobile() && detect.IsTablet(){
			//fmt.Println("Hello, this is Tablet")
		}else {
			//fmt.Println("Hello, this is Mobile")
		}
		//return true
		return indexoradmin + "/" + conf.WapPath + "/" + model + "/"
	}else {
		//fmt.Println("Hello, this is Desktop")
		//return false
		return indexoradmin + "/" + conf.PcPath + "/" + model + "/"
	}
}

// http.Request 是否启用自适应方式默认flase cookie设备信息 前台或者后台 模块路径
func PathAutoOrMobile(r *http.Request, b bool, cookieMobile string, indexoradmin string, model string) string{
	if b {
		return indexoradmin + "/" + conf.AutoPath + "/" + model + "/"
	}

	if cookieMobile == conf.WapPath || cookieMobile == conf.PcPath {
		return indexoradmin + "/" + cookieMobile + "/" + model + "/"
	}

	detect := mobiledetect.NewMobileDetect(r, nil)
	if detect.IsMobile() || detect.IsTablet(){
		if detect.IsMobile() && detect.IsTablet(){
			//fmt.Println("Hello, this is Tablet")
		}else {
			//fmt.Println("Hello, this is Mobile")
		}
		//return true
		return indexoradmin + "/" + conf.WapPath + "/" + model + "/"
	}else {
		//fmt.Println("Hello, this is Desktop")
		//return false
		return indexoradmin + "/" + conf.PcPath + "/" + model + "/"
	}
}
