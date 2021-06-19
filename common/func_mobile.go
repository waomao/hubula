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

func PathMobile(r *http.Request,indexoradmin string,model string) string{
	detect := mobiledetect.NewMobileDetect(r, nil)
	if detect.IsMobile() || detect.IsTablet(){
		if detect.IsMobile() && detect.IsTablet(){
			//fmt.Println("Hello, this is Tablet")
		}else {
			//fmt.Println("Hello, this is Mobile")
		}
		//return true
		return indexoradmin + conf.WapPath + model
	}else {
		//fmt.Println("Hello, this is Desktop")
		//return false
		return indexoradmin + conf.PcPath + model
	}
}
