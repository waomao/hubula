package common

type YYSeveralPic struct {
	Quanzhang int64
	Denglong int64
	Huangguan int64
	Dunpai int64
	Zhuanshi int64
	Shuye int64
}

type QQSeveralPic struct {
	Huangguan int64
	Sun int64
	Moon int64
	Star int64
}

func YYLeveToSeveral(leve int64) *YYSeveralPic{
	levePic := new(YYSeveralPic)
	levePic.Quanzhang = leve / 1024
	levePic.Denglong = leve % 1024 / 256
	levePic.Huangguan = leve % 1024 % 256 / 64
	levePic.Dunpai = leve % 1024 % 256 % 64 / 16
	levePic.Zhuanshi = leve % 1024 % 256 % 64 % 16 / 4
	levePic.Shuye = leve % 1024 % 256 % 64 % 16 % 4
	return levePic
}


func YYLevePic(leve *YYSeveralPic) string{
	picStr := ""

	if leve.Quanzhang > 0 {
		for i := 0; i < int(leve.Quanzhang); i++ {
			picStr += "<img src='/public/img/太阳.png'  width='40px' height='40px' alt=''>"
		}
	}

	if leve.Denglong > 0 {
		for i := 0; i < int(leve.Denglong); i++ {
			picStr += "<img src='/public/img/太阳.png'  width='40px' height='40px' alt=''>"
		}
	}

	if leve.Huangguan > 0 {
		for i := 0; i < int(leve.Huangguan); i++ {
			picStr += "<img src='/public/img/太阳.png'  width='40px' height='40px' alt=''>"
		}
	}

	if leve.Dunpai > 0 {
		for i := 0; i < int(leve.Dunpai); i++ {
			picStr += "<img src='/public/img/太阳.png'  width='40px' height='40px' alt=''>"
		}
	}

	if leve.Zhuanshi > 0 {
		for i := 0; i < int(leve.Zhuanshi); i++ {
			picStr += "<img src='/public/img/太阳.png'  width='40px' height='40px' alt=''>"
		}
	}

	if leve.Shuye > 0 {
		for i := 0; i < int(leve.Shuye); i++ {
			picStr += "<img src='/public/img/太阳.png'  width='40px' height='40px' alt=''>"
		}
	}

	return picStr
}

func QQLeveToSeveral(leve int64) *QQSeveralPic{
	levePic := new(QQSeveralPic)
	levePic.Huangguan = leve / 64
	levePic.Sun = leve % 64 / 16
	levePic.Moon = leve % 64 % 16 / 4
	levePic.Star = leve % 64 % 16 % 4
	return levePic
}


func QQLevePic(leve *QQSeveralPic) string{
	picStr := ""

	if leve.Huangguan > 0 {
		for i := 0; i < int(leve.Huangguan); i++ {
			picStr += "<img src='/public/img/皇冠.png'  width='40px' height='40px' alt=''>"
		}
	}

	if leve.Sun > 0 {
		for i := 0; i < int(leve.Sun); i++ {
			picStr += "<img src='/public/img/太阳.png'  width='40px' height='40px' alt=''>"
		}
	}

	if leve.Moon > 0 {
		for i := 0; i < int(leve.Moon); i++ {
			picStr += "<img src='/public/img/月亮.png' width='40px' height='40px' alt=''>"
		}
	}

	if leve.Star > 0 {
		for i := 0; i < int(leve.Star); i++ {
			picStr += "<img src='/public/img/星星.png'  width='40px' height='40px' alt=''>"
		}
	}

	return picStr
}
