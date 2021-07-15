package common

func Sqlyuju(biaoming string,keyValue string) string {
	return "SELECT * FROM `" + biaoming + "` WHERE `add` LIKE '%" + keyValue + "%' "
}

//生日 只能匹配末尾四位
func SqlShengRi(shengri string) (v map[string]string){
	shengris := make(map[string]map[string]string)

	shengris["nil"] = make(map[string]string)
	shengris["nil"]["key"] = ""
	shengris["nil"]["value"] = ""
	shengris["nil"]["sql"] = ""

	shengris["1"] = make(map[string]string)
	shengris["1"]["key"] = "一月"
	shengris["1"]["value"] = "1"
	shengris["1"]["sql"] = "floor(mid(right(`haoma`,4),1,2)) = '01' and floor(mid(right(`haoma`,2),1,2)) <= '31'"

	shengris["2"] = make(map[string]string)
	shengris["2"]["key"] = "二月"
	shengris["2"]["value"] = "2"
	shengris["2"]["sql"] = "floor(mid(right(haoma,4),1,2)) = 02 and floor(mid(right(haoma,2),1,2)) <= 31"

	shengris["3"] = make(map[string]string)
	shengris["3"]["key"] = "三月"
	shengris["3"]["value"] = "3"
	shengris["3"]["sql"] = "floor(mid(right(haoma,4),1,2)) = 03 and floor(mid(right(haoma,2),1,2)) <= 31"

	shengris["4"] = make(map[string]string)
	shengris["4"]["key"] = "四月"
	shengris["4"]["value"] = "4"
	shengris["4"]["sql"] = "floor(mid(right(haoma,4),1,2)) = 04 and floor(mid(right(haoma,2),1,2)) <= 31"

	shengris["5"] = make(map[string]string)
	shengris["5"]["key"] = "五月"
	shengris["5"]["value"] = "5"
	shengris["5"]["sql"] = "floor(mid(right(haoma,4),1,2)) = 05 and floor(mid(right(haoma,2),1,2)) <= 31"

	shengris["6"] = make(map[string]string)
	shengris["6"]["key"] = "六月"
	shengris["6"]["value"] = "6"
	shengris["6"]["sql"] = "floor(mid(right(haoma,4),1,2)) = 06 and floor(mid(right(haoma,2),1,2)) <= 31"

	shengris["7"] = make(map[string]string)
	shengris["7"]["key"] = "七月"
	shengris["7"]["value"] = "7"
	shengris["7"]["sql"] = "floor(mid(right(haoma,4),1,2)) = 07 and floor(mid(right(haoma,2),1,2)) <= 31"

	shengris["8"] = make(map[string]string)
	shengris["8"]["key"] = "八月"
	shengris["8"]["value"] = "8"
	shengris["8"]["sql"] = "floor(mid(right(haoma,4),1,2)) = 08 and floor(mid(right(haoma,2),1,2)) <= 31"

	shengris["9"] = make(map[string]string)
	shengris["9"]["key"] = "九月"
	shengris["9"]["value"] = "9"
	shengris["9"]["sql"] = "floor(mid(right(haoma,4),1,2)) = 09 and floor(mid(right(haoma,2),1,2)) <= 31"

	shengris["10"] = make(map[string]string)
	shengris["10"]["key"] = "十月"
	shengris["10"]["value"] = "10"
	shengris["10"]["sql"] = "floor(mid(right(haoma,4),1,2)) = 10 and floor(mid(right(haoma,2),1,2)) <= 31"

	shengris["11"] = make(map[string]string)
	shengris["11"]["key"] = "十一月"
	shengris["11"]["value"] = "11"
	shengris["11"]["sql"] = "floor(mid(right(haoma,4),1,2)) = 11 and floor(mid(right(haoma,2),1,2)) <= 31"

	shengris["12"] = make(map[string]string)
	shengris["12"]["key"] = "十二月"
	shengris["12"]["value"] = "12"
	shengris["12"]["sql"] = "floor(mid(right(haoma,4),1,2)) = 12 and floor(mid(right(haoma,2),1,2)) <= 31"

	//判断name字段是否存在
	//v,ok := users[name]
	if shengris[shengri] != nil {
		//存在
		v = shengris[shengri]
	} else {
		//不存在
		v = shengris["nil"]
	}
	return
}


//不含
func SqlBuHan(buhan string) (v map[string]string){
	buhans := make(map[string]map[string]string)

	buhans["nil"] = make(map[string]string)
	buhans["nil"]["key"] = ""
	buhans["nil"]["value"] = ""
	buhans["nil"]["sql"] = ""

	buhans["z0"] = make(map[string]string)
	buhans["z0"]["key"] = "不含0"
	buhans["z0"]["value"] = "z0"
	buhans["z0"]["sql"] = "`haoma` not like '%0%'"

	buhans["1"] = make(map[string]string)
	buhans["1"]["key"] = "不含1"
	buhans["1"]["value"] = "1"
	buhans["1"]["sql"] = "`haoma` not like '%1%'"

	buhans["2"] = make(map[string]string)
	buhans["2"]["key"] = "不含2"
	buhans["2"]["value"] = "2"
	buhans["2"]["sql"] = "`haoma` not like '%2%'"

	buhans["3"] = make(map[string]string)
	buhans["3"]["key"] = "不含3"
	buhans["3"]["value"] = "3"
	buhans["3"]["sql"] = "`haoma` not like '%3%'"

	buhans["4"] = make(map[string]string)
	buhans["4"]["key"] = "不含4"
	buhans["4"]["value"] = "4"
	buhans["4"]["sql"] = "`haoma` not like '%4%'"

	buhans["5"] = make(map[string]string)
	buhans["5"]["key"] = "不含5"
	buhans["5"]["value"] = "5"
	buhans["5"]["sql"] = "`haoma` not like '%5%'"

	buhans["6"] = make(map[string]string)
	buhans["6"]["key"] = "不含6"
	buhans["6"]["value"] = "6"
	buhans["6"]["sql"] = "`haoma` not like '%6%'"

	buhans["7"] = make(map[string]string)
	buhans["7"]["key"] = "不含7"
	buhans["7"]["value"] = "7"
	buhans["7"]["sql"] = "`haoma` not like '%7%'"

	buhans["47"] = make(map[string]string)
	buhans["47"]["key"] = "不含47"
	buhans["47"]["value"] = "47"
	buhans["47"]["sql"] = "`haoma` not like '%4%' and `haoma`  not like '%7%'"

	buhans["8"] = make(map[string]string)
	buhans["8"]["key"] = "不含8"
	buhans["8"]["value"] = "8"
	buhans["8"]["sql"] = "`haoma` not like '%8%'"

	buhans["9"] = make(map[string]string)
	buhans["9"]["key"] = "不含9"
	buhans["9"]["value"] = "9"
	buhans["9"]["sql"] = "`haoma` not like '%9%'"
	//判断name字段是否存在
	//v,ok := users[name]
	if buhans[buhan] != nil {
		//存在
		v = buhans[buhan]
	} else {
		//不存在
		v = buhans["nil"]
	}
	return
}

//较多
//直接匹配到至少4个
func SqlJiaoDuo(jiaoduo string) (v map[string]string){
	jiaoduos := make(map[string]map[string]string)

	jiaoduos["nil"] = make(map[string]string)
	jiaoduos["nil"]["key"] = ""
	jiaoduos["nil"]["value"] = ""
	jiaoduos["nil"]["sql"] = ""

	jiaoduos["10"] = make(map[string]string)
	jiaoduos["10"]["key"] = "0较多"
	jiaoduos["10"]["value"] = "10"
	jiaoduos["10"]["sql"] = "`haoma` like '%0%0%0%0%'"

	jiaoduos["1"] = make(map[string]string)
	jiaoduos["1"]["key"] = "1较多"
	jiaoduos["1"]["value"] = "1"
	jiaoduos["1"]["sql"] = "`haoma` like '%1%1%1%1%'"

	jiaoduos["2"] = make(map[string]string)
	jiaoduos["2"]["key"] = "2较多"
	jiaoduos["2"]["value"] = "2"
	jiaoduos["2"]["sql"] = "`haoma` like '%2%2%2%2%'"

	jiaoduos["3"] = make(map[string]string)
	jiaoduos["3"]["key"] = "3较多"
	jiaoduos["3"]["value"] = "3"
	jiaoduos["3"]["sql"] = "`haoma` like '%3%3%3%3%'"

	jiaoduos["4"] = make(map[string]string)
	jiaoduos["4"]["key"] = "4较多"
	jiaoduos["4"]["value"] = "4"
	jiaoduos["4"]["sql"] = "`haoma` like '%4%4%4%4%'"

	jiaoduos["5"] = make(map[string]string)
	jiaoduos["5"]["key"] = "5较多"
	jiaoduos["5"]["value"] = "5"
	jiaoduos["5"]["sql"] = "`haoma` like '%5%5%5%5%'"

	jiaoduos["6"] = make(map[string]string)
	jiaoduos["6"]["key"] = "6较多"
	jiaoduos["6"]["value"] = "6"
	jiaoduos["6"]["sql"] = "`haoma` like '%6%6%6%6%'"

	jiaoduos["7"] = make(map[string]string)
	jiaoduos["7"]["key"] = "7较多"
	jiaoduos["7"]["value"] = "7"
	jiaoduos["7"]["sql"] = "`haoma` like '%7%7%7%7%'"

	jiaoduos["8"] = make(map[string]string)
	jiaoduos["8"]["key"] = "8较多"
	jiaoduos["8"]["value"] = "8"
	jiaoduos["8"]["sql"] = "`haoma` like '%8%8%8%8%'"

	jiaoduos["9"] = make(map[string]string)
	jiaoduos["9"]["key"] = "9较多"
	jiaoduos["9"]["value"] = "9"
	jiaoduos["9"]["sql"] = "`haoma` like '%9%9%9%9%'"
	//判断name字段是否存在
	//v,ok := users[name]
	if jiaoduos[jiaoduo] != nil {
		//存在
		v = jiaoduos[jiaoduo]
	} else {
		//不存在
		v = jiaoduos["nil"]
	}
	return
}

//排序
func SqlOrder(order string) (v map[string]string){
	orders := make(map[string]map[string]string)

	orders["nil"] = make(map[string]string)
	orders["nil"]["key"] = ""
	orders["nil"]["value"] = " ORDER BY listorder DESC,id DESC"
	orders["nil"]["sql"] = " ORDER BY haoma DESC,id DESC"

	orders["price1_99"] = make(map[string]string)
	orders["price1_99"]["key"] = "价格低到高,初装费,月消费"
	orders["price1_99"]["value"] = "price1_99"
	orders["price1_99"]["sql"] = "ORDER BY kafei ASC"

	orders["price99_1"] = make(map[string]string)
	orders["price99_1"]["key"] = "价格高到低,初装费,月消费"
	orders["price99_1"]["value"] = "price99_1"
	orders["price99_1"]["sql"] = " ORDER BY kafei DESC"

	//判断name字段是否存在
	//v,ok := users[name]
	if orders[order] != nil {
		//存在
		v = orders[order]
	} else {
		//不存在
		v = orders["nil"]
	}
	return
}

//年代 末尾四位
func SqlNianDai(niandai string) (v map[string]string){
	niandais := make(map[string]map[string]string)

	niandais["nil"] = make(map[string]string)
	niandais["nil"]["key"] = ""
	niandais["nil"]["value"] = ""
	niandais["nil"]["sql"] = ""

	niandais["1950"] = make(map[string]string)
	niandais["1950"]["key"] = "1950年代"
	niandais["1950"]["value"] = "1950"
	niandais["1950"]["sql"] = "right(haoma,4) >= 1950 and right(haoma,4) < 1960"

	niandais["1960"] = make(map[string]string)
	niandais["1960"]["key"] = "1960年代"
	niandais["1960"]["value"] = "1960"
	niandais["1960"]["sql"] = "right(haoma,4) >= 1960 and right(haoma,4) < 1970"

	niandais["1970"] = make(map[string]string)
	niandais["1970"]["key"] = "1970年代"
	niandais["1970"]["value"] = "1970"
	niandais["1970"]["sql"] = "right(haoma,4) >= 1970 and right(haoma,4) < 1980"

	niandais["1980"] = make(map[string]string)
	niandais["1980"]["key"] = "1980年代"
	niandais["1980"]["value"] = "1980"
	niandais["1980"]["sql"] = "right(haoma,4) >= 1980 and right(haoma,4) < 1990"

	niandais["1990"] = make(map[string]string)
	niandais["1990"]["key"] = "1990年代"
	niandais["1990"]["value"] = "1990"
	niandais["1990"]["sql"] = "right(haoma,4) >= 1990 and right(haoma,4) < 2000"

	niandais["2000"] = make(map[string]string)
	niandais["2000"]["key"] = "2000年代"
	niandais["2000"]["value"] = "2000"
	niandais["2000"]["sql"] = "right(haoma,4) >= 2000 and right(haoma,4) < 2010"

	niandais["2010"] = make(map[string]string)
	niandais["2010"]["key"] = "2010年代"
	niandais["2010"]["value"] = "2010"
	niandais["2010"]["sql"] = "right(haoma,4) >= 2010 and right(haoma,4) < 2020"

	niandais["2020"] = make(map[string]string)
	niandais["2020"]["key"] = "2020年代"
	niandais["2020"]["value"] = "2020"
	niandais["2020"]["sql"] = "right(haoma,4) >= 2020 and right(haoma,4) < 2030"

	niandais["2030"] = make(map[string]string)
	niandais["2030"]["key"] = "2030年代"
	niandais["2030"]["value"] = "2030"
	niandais["2030"]["sql"] = "right(haoma,4) >= 2030 and right(haoma,4) < 2040"

	niandais["2040"] = make(map[string]string)
	niandais["2040"]["key"] = "2040年代"
	niandais["2040"]["value"] = "2040"
	niandais["2040"]["sql"] = "right(haoma,4) >= 2040 and right(haoma,4) < 2050"

	niandais["2050"] = make(map[string]string)
	niandais["2050"]["key"] = "2050年代"
	niandais["2050"]["value"] = "2050"
	niandais["2050"]["sql"] = "right(haoma,4) >= 2050 and right(haoma,4) < 2060"

	niandais["2060"] = make(map[string]string)
	niandais["2060"]["key"] = "2060年代"
	niandais["2060"]["value"] = "2060"
	niandais["2060"]["sql"] = "right(haoma,4) >= 2060 and right(haoma,4) < 2070"

	niandais["2070"] = make(map[string]string)
	niandais["2070"]["key"] = "2070年代"
	niandais["2070"]["value"] = "2070"
	niandais["2070"]["sql"] = "right(haoma,4) >= 2070 and right(haoma,4) < 2080"

	niandais["2080"] = make(map[string]string)
	niandais["2080"]["key"] = "2080年代"
	niandais["2080"]["value"] = "2080"
	niandais["2080"]["sql"] = "right(haoma,4) >= 2080 and right(haoma,4) < 2090"

	niandais["2090"] = make(map[string]string)
	niandais["2090"]["key"] = "2090年代"
	niandais["2090"]["value"] = "2090"
	niandais["2090"]["sql"] = "right(haoma,4) >= 2090 and right(haoma,4) < 2100"

	niandais["2100"] = make(map[string]string)
	niandais["2100"]["key"] = "2100年代"
	niandais["2100"]["value"] = "2100"
	niandais["2100"]["sql"] = "right(haoma,4) >= 2100 and right(haoma,4) < 2110"

	//判断name字段是否存在
	//v,ok := users[name]
	if niandais[niandai] != nil {
		//存在
		v = niandais[niandai]
	} else {
		//不存在
		v = niandais["nil"]
	}
	return
}

//价格
func SqlJiaGe(jiage string) (v map[string]string){
	jiages := make(map[string]map[string]string)

	jiages["nil"] = make(map[string]string)
	jiages["nil"]["key"] = ""
	jiages["nil"]["value"] = ""
	jiages["nil"]["sql"] = ""

	jiages["1"] = make(map[string]string)
	jiages["1"]["key"] = "待询价"
	jiages["1"]["value"] = "1"
	jiages["1"]["sql"] = "`jiage` = '0'"

	jiages["100"] = make(map[string]string)
	jiages["100"]["key"] = "100元以下"
	jiages["100"]["value"] = "100"
	jiages["100"]["sql"] = "`jiage` > '0' and `jiage` < '100'"

	jiages["200"] = make(map[string]string)
	jiages["200"]["key"] = "100-200元"
	jiages["200"]["value"] = "200"
	jiages["200"]["sql"] = "`jiage` >= '100' and `jiage` <= '200'"

	jiages["300"] = make(map[string]string)
	jiages["300"]["key"] = "200-300元"
	jiages["300"]["value"] = "300"
	jiages["300"]["sql"] = "`jiage` >= '200' and `jiage` <= '300'"

	jiages["500"] = make(map[string]string)
	jiages["500"]["key"] = "300-500元"
	jiages["500"]["value"] = "500"
	jiages["500"]["sql"] = "`jiage` >= '300' and `jiage` <= '500'"

	jiages["1000"] = make(map[string]string)
	jiages["1000"]["key"] = "500-1000元"
	jiages["1000"]["value"] = "1000"
	jiages["1000"]["sql"] = "`jiage` >= 500 and `jiage` <= 1000"

	jiages["2000"] = make(map[string]string)
	jiages["2000"]["key"] = "1000-2000元"
	jiages["2000"]["value"] = "2000"
	jiages["2000"]["sql"] = "`jiage` >= '1000' and `jiage` <= '2000'"

	jiages["5000"] = make(map[string]string)
	jiages["5000"]["key"] = "2000-5000元"
	jiages["5000"]["value"] = "5000"
	jiages["5000"]["sql"] = "`jiage` >= '2000' and `jiage` <= '5000'"

	jiages["10000"] = make(map[string]string)
	jiages["10000"]["key"] = "5000-10000元"
	jiages["10000"]["value"] = "10000"
	jiages["10000"]["sql"] = "`jiage` >= '5000' and `jiage` <= '10000'"

	jiages["20000"] = make(map[string]string)
	jiages["20000"]["key"] = "10000-20000元"
	jiages["20000"]["value"] = "20000"
	jiages["20000"]["sql"] = "`jiage` >= '10000' and `jiage` <= '200000'"

	jiages["30000"] = make(map[string]string)
	jiages["30000"]["key"] = "20000-30000元"
	jiages["30000"]["value"] = "30000"
	jiages["30000"]["sql"] = "`jiage` >= '20000' and `jiage` <= '30000'"

	jiages["50000"] = make(map[string]string)
	jiages["50000"]["key"] = "30000-50000元"
	jiages["50000"]["value"] = "50000"
	jiages["50000"]["sql"] = "`jiage` >= '30000' and `jiage` <= '50000'"

	jiages["100000"] = make(map[string]string)
	jiages["100000"]["key"] = "50000元以上"
	jiages["100000"]["value"] = "100000"
	jiages["100000"]["sql"] = "`jiage` >= '50000'"

	//判断name字段是否存在
	//v,ok := users[name]
	if jiages[jiage] != nil {
		//存在
		v = jiages[jiage]
	} else {
		//不存在
		v = jiages["nil"]
	}
	return
}

//吉利号码
func SqlJiLiHao(jilihao string) (v map[string]string){
	jilihaos := make(map[string]map[string]string)

	jilihaos["nil"] = make(map[string]string)
	jilihaos["nil"]["key"] = ""
	jilihaos["nil"]["value"] = ""
	jilihaos["nil"]["sql"] = ""

	jilihaos["888"] = make(map[string]string)
	jilihaos["888"]["key"] = "888"
	jilihaos["888"]["value"] = "888"
	jilihaos["888"]["sql"] = "instr(haoma,888) > 0"

	jilihaos["88"] = make(map[string]string)
	jilihaos["88"]["key"] = "88"
	jilihaos["88"]["value"] = "88"
	jilihaos["88"]["sql"] = "instr(haoma,88) > 0"

	jilihaos["666"] = make(map[string]string)
	jilihaos["666"]["key"] = "666"
	jilihaos["666"]["value"] = "666"
	jilihaos["666"]["sql"] = "instr(haoma,666) > 0"

	jilihaos["66"] = make(map[string]string)
	jilihaos["66"]["key"] = "66"
	jilihaos["66"]["value"] = "66"
	jilihaos["66"]["sql"] = "instr(haoma,66) > 0"

	jilihaos["1188"] = make(map[string]string)
	jilihaos["1188"]["key"] = "1188"
	jilihaos["1188"]["value"] = "1188"
	jilihaos["1188"]["sql"] = "instr(haoma,1188) > 0"

	jilihaos["168"] = make(map[string]string)
	jilihaos["168"]["key"] = "168"
	jilihaos["168"]["value"] = "168"
	jilihaos["168"]["sql"] = "instr(haoma,168) > 0"

	jilihaos["5588"] = make(map[string]string)
	jilihaos["5588"]["key"] = "5588"
	jilihaos["5588"]["value"] = "5588"
	jilihaos["5588"]["sql"] = "instr(haoma,5588) > 0"

	jilihaos["8199"] = make(map[string]string)
	jilihaos["8199"]["key"] = "8199"
	jilihaos["8199"]["value"] = "8199"
	jilihaos["8199"]["sql"] = "instr(haoma,8199) > 0"

	jilihaos["598"] = make(map[string]string)
	jilihaos["598"]["key"] = "598"
	jilihaos["598"]["value"] = "598"
	jilihaos["598"]["sql"] = "instr(haoma,598) > 0"

	jilihaos["999"] = make(map[string]string)
	jilihaos["999"]["key"] = "999"
	jilihaos["999"]["value"] = "999"
	jilihaos["999"]["sql"] = "instr(haoma,999) > 0"
	//判断name字段是否存在
	//v,ok := users[name]
	if jilihaos[jilihao] != nil {
		//存在
		v = jilihaos[jilihao]
	} else {
		//不存在
		v = jilihaos["nil"]
	}
	return
}

//手机号段 匹配前三个
func SqlHaoduanSJH(haoduan string) (v map[string]string){
	haoduans := make(map[string]map[string]string)

	haoduans["nil"] = make(map[string]string)
	haoduans["nil"]["key"] = ""
	haoduans["nil"]["value"] = ""
	haoduans["nil"]["sql"] = ""

	haoduans["166"] = make(map[string]string)
	haoduans["166"]["key"] = "166"
	haoduans["166"]["value"] = "166"
	haoduans["166"]["sql"] = "floor(LEFT(haoma,3)) = '166'"

	haoduans["147"] = make(map[string]string)
	haoduans["147"]["key"] = "147"
	haoduans["147"]["value"] = "147"
	haoduans["147"]["sql"] = "floor(LEFT(haoma,3)) = '147'"

	haoduans["173"] = make(map[string]string)
	haoduans["173"]["key"] = "173"
	haoduans["173"]["value"] = "173"
	haoduans["173"]["sql"] = "floor(LEFT(haoma,3)) = '173'"

	haoduans["181"] = make(map[string]string)
	haoduans["181"]["key"] = "181"
	haoduans["181"]["value"] = "181"
	haoduans["181"]["sql"] = "floor(LEFT(haoma,3)) = '181'"

	haoduans["176"] = make(map[string]string)
	haoduans["176"]["key"] = "176"
	haoduans["176"]["value"] = "176"
	haoduans["176"]["sql"] = "floor(LEFT(haoma,3)) = '176'"

	haoduans["183"] = make(map[string]string)
	haoduans["183"]["key"] = "183"
	haoduans["183"]["value"] = "183"
	haoduans["183"]["sql"] = "floor(LEFT(haoma,3)) = '183'"

	haoduans["178"] = make(map[string]string)
	haoduans["178"]["key"] = "178"
	haoduans["178"]["value"] = "178"
	haoduans["178"]["sql"] = "floor(LEFT(haoma,3)) = '178'"

	haoduans["182"] = make(map[string]string)
	haoduans["182"]["key"] = "182"
	haoduans["182"]["value"] = "182"
	haoduans["182"]["sql"] = "floor(LEFT(haoma,3)) = '182'"

	haoduans["177"] = make(map[string]string)
	haoduans["177"]["key"] = "177"
	haoduans["177"]["value"] = "177"
	haoduans["177"]["sql"] = "floor(LEFT(haoma,3)) = '177'"

	haoduans["189"] = make(map[string]string)
	haoduans["189"]["key"] = "189"
	haoduans["189"]["value"] = "189"
	haoduans["189"]["sql"] = "floor(LEFT(haoma,3)) = '189'"

	haoduans["180"] = make(map[string]string)
	haoduans["180"]["key"] = "180"
	haoduans["180"]["value"] = "180"
	haoduans["180"]["sql"] = "floor(LEFT(haoma,3)) = '180'"

	haoduans["153"] = make(map[string]string)
	haoduans["153"]["key"] = "153"
	haoduans["153"]["value"] = "153"
	haoduans["153"]["sql"] = "floor(LEFT(haoma,3)) = '153'"

	haoduans["185"] = make(map[string]string)
	haoduans["185"]["key"] = "185"
	haoduans["185"]["value"] = "185"
	haoduans["185"]["sql"] = "floor(LEFT(haoma,3)) = '185'"

	haoduans["186"] = make(map[string]string)
	haoduans["186"]["key"] = "186"
	haoduans["186"]["value"] = "186"
	haoduans["186"]["sql"] = "floor(LEFT(haoma,3)) = '186'"

	haoduans["156"] = make(map[string]string)
	haoduans["156"]["key"] = "156"
	haoduans["156"]["value"] = "156"
	haoduans["156"]["sql"] = "floor(LEFT(haoma,3)) = '156'"

	haoduans["155"] = make(map[string]string)
	haoduans["155"]["key"] = "155"
	haoduans["155"]["value"] = "155"
	haoduans["155"]["sql"] = "floor(LEFT(haoma,3)) = '155'"

	haoduans["130"] = make(map[string]string)
	haoduans["130"]["key"] = "130"
	haoduans["130"]["value"] = "130"
	haoduans["130"]["sql"] = "floor(LEFT(haoma,3)) = '130'"

	haoduans["187"] = make(map[string]string)
	haoduans["187"]["key"] = "187"
	haoduans["187"]["value"] = "187"
	haoduans["187"]["sql"] = "floor(LEFT(haoma,3)) = '187'"

	haoduans["188"] = make(map[string]string)
	haoduans["188"]["key"] = "188"
	haoduans["188"]["value"] = "188"
	haoduans["188"]["sql"] = "floor(LEFT(haoma,3)) = '188'"

	haoduans["157"] = make(map[string]string)
	haoduans["157"]["key"] = "157"
	haoduans["157"]["value"] = "157"
	haoduans["157"]["sql"] = "floor(LEFT(haoma,3)) = '157'"

	haoduans["159"] = make(map[string]string)
	haoduans["159"]["key"] = "159"
	haoduans["159"]["value"] = "159"
	haoduans["159"]["sql"] = "floor(LEFT(haoma,3)) = '159'"

	haoduans["158"] = make(map[string]string)
	haoduans["158"]["key"] = "158"
	haoduans["158"]["value"] = "158"
	haoduans["158"]["sql"] = "floor(LEFT(haoma,3)) = '158'"

	haoduans["152"] = make(map[string]string)
	haoduans["152"]["key"] = "152"
	haoduans["152"]["value"] = "152"
	haoduans["152"]["sql"] = "floor(LEFT(haoma,3)) = '152'"

	haoduans["151"] = make(map[string]string)
	haoduans["151"]["key"] = "151"
	haoduans["151"]["value"] = "151"
	haoduans["151"]["sql"] = "floor(LEFT(haoma,3)) = '151'"

	haoduans["150"] = make(map[string]string)
	haoduans["150"]["key"] = "150"
	haoduans["150"]["value"] = "150"
	haoduans["150"]["sql"] = "floor(LEFT(haoma,3)) = '150'"

	haoduans["139"] = make(map[string]string)
	haoduans["139"]["key"] = "139"
	haoduans["139"]["value"] = "139"
	haoduans["139"]["sql"] = "floor(LEFT(haoma,3)) = '139'"

	haoduans["138"] = make(map[string]string)
	haoduans["138"]["key"] = "138"
	haoduans["138"]["value"] = "138"
	haoduans["138"]["sql"] = "floor(LEFT(haoma,3)) = '138'"

	haoduans["137"] = make(map[string]string)
	haoduans["137"]["key"] = "137"
	haoduans["137"]["value"] = "137"
	haoduans["137"]["sql"] = "floor(LEFT(haoma,3)) = '137'"

	haoduans["136"] = make(map[string]string)
	haoduans["136"]["key"] = "136"
	haoduans["136"]["value"] = "136"
	haoduans["136"]["sql"] = "floor(LEFT(haoma,3)) = '136'"

	haoduans["135"] = make(map[string]string)
	haoduans["135"]["key"] = "135"
	haoduans["135"]["value"] = "135"
	haoduans["135"]["sql"] = "floor(LEFT(haoma,3)) = '135'"

	haoduans["134"] = make(map[string]string)
	haoduans["134"]["key"] = "134"
	haoduans["134"]["value"] = "134"
	haoduans["134"]["sql"] = "floor(LEFT(haoma,3)) = '134'"

	haoduans["133"] = make(map[string]string)
	haoduans["133"]["key"] = "133"
	haoduans["133"]["value"] = "133"
	haoduans["133"]["sql"] = "floor(LEFT(haoma,3)) = '133'"

	haoduans["132"] = make(map[string]string)
	haoduans["132"]["key"] = "132"
	haoduans["132"]["value"] = "132"
	haoduans["132"]["sql"] = "floor(LEFT(haoma,3)) = '132'"

	haoduans["131"] = make(map[string]string)
	haoduans["131"]["key"] = "131"
	haoduans["131"]["value"] = "131"
	haoduans["131"]["sql"] = "floor(LEFT(haoma,3)) = '131'"

	//判断name字段是否存在
	//v,ok := users[name]
	if haoduans[haoduan] != nil {
		//存在
		v = haoduans[haoduan]
	} else {
		//不存在
		v = haoduans["nil"]
	}
	return
}

//400号段
func SqlHaoduan400(haoduan string) (v map[string]string){
	haoduans := make(map[string]map[string]string)

	haoduans["nil"] = make(map[string]string)
	haoduans["nil"]["key"] = ""
	haoduans["nil"]["value"] = ""
	haoduans["nil"]["sql"] = ""

	haoduans["4009"] = make(map[string]string)
	haoduans["4009"]["key"] = "4009"
	haoduans["4009"]["value"] = "4009"
	haoduans["4009"]["sql"] = "floor(LEFT(haoma,4)) = '4009'"

	haoduans["4008"] = make(map[string]string)
	haoduans["4008"]["key"] = "4008"
	haoduans["4008"]["value"] = "4008"
	haoduans["4008"]["sql"] = "floor(LEFT(haoma,4)) = '4008'"

	haoduans["4007"] = make(map[string]string)
	haoduans["4007"]["key"] = "4007"
	haoduans["4007"]["value"] = "4007"
	haoduans["4007"]["sql"] = "floor(LEFT(haoma,4)) = '4007'"

	haoduans["4006"] = make(map[string]string)
	haoduans["4006"]["key"] = "4006"
	haoduans["4006"]["value"] = "4006"
	haoduans["4006"]["sql"] = "floor(LEFT(haoma,4)) = '4006'"

	haoduans["4000"] = make(map[string]string)
	haoduans["4000"]["key"] = "4000"
	haoduans["4000"]["value"] = "4000"
	haoduans["4000"]["sql"] = "floor(LEFT(haoma,4)) = '4000'"

	haoduans["4001"] = make(map[string]string)
	haoduans["4001"]["key"] = "4001"
	haoduans["4001"]["value"] = "4001"
	haoduans["4001"]["sql"] = "floor(LEFT(haoma,4)) = '4001'"

	//判断name字段是否存在
	//v,ok := users[name]
	if haoduans[haoduan] != nil {
		//存在
		v = haoduans[haoduan]
	} else {
		//不存在
		v = haoduans["nil"]
	}
	return
}

//号码特征 末尾匹配
func SqlTeZheng(tezheng string) (v map[string]string){
	tezhengs := make(map[string]map[string]string)

	tezhengs["nil"] = make(map[string]string)
	tezhengs["nil"]["key"] = ""
	tezhengs["nil"]["value"] = ""
	tezhengs["nil"]["sql"] = ""

	tezhengs["1"] = make(map[string]string)
	tezhengs["1"]["key"] = "AAAAA"
	tezhengs["1"]["value"] = "1"
	tezhengs["1"]["sql"] = "(" +
		"(floor(mid(haoma,2,5)) mod 11111 = 0 and " +
		"floor(mid(haoma,6,1))<>floor(mid(haoma,7,1)))" +
		"or ( floor(mid(haoma,3,5)) mod 11111 = 0 and " +
		"floor(mid(haoma,7,1))<>floor(mid(haoma,8,1)) and " +
		"floor(mid(haoma,3,1))<>floor(mid(haoma,2,1)) )" +
		"or ( floor(mid(haoma,4,5)) mod 11111 = 0 and " +
		"floor(mid(haoma,8,1))<>floor(mid(haoma,9,1)) " +
		"and floor(mid(haoma,4,1))<>floor(mid(haoma,3,1)) )" +
		"or ( floor(mid(haoma,5,5)) mod 11111 = 0 and " +
		"floor(mid(haoma,9,1))<>floor(mid(haoma,10,1)) and " +
		"floor(mid(haoma,5,1))<>floor(mid(haoma,4,1)) )" +
		"or ( floor(mid(haoma,6,5)) mod 11111 = 0 and " +
		"floor(mid(haoma,10,1))<>floor(mid(haoma,11,1)) and " +
		"floor(mid(haoma,6,1))<>floor(mid(haoma,5,1)) )" +
		"or ( floor(mid(haoma,7,5)) mod 11111 = 0 and " +
		"floor(mid(haoma,7,1))<>floor(mid(haoma,6,1)) )" +
		")"

	tezhengs["2"] = make(map[string]string)
	tezhengs["2"]["key"] = "AAAA"
	tezhengs["2"]["value"] = "2"
	tezhengs["2"]["sql"] = "(" +
		"( floor(mid(haoma,2,4)) mod 1111 = 0 and " +
		"floor(mid(haoma,5,1))<>floor(mid(haoma,6,1)) )" +
		"or ( floor(mid(haoma,3,4)) mod 1111 = 0 and " +
		"floor(mid(haoma,6,1))<>floor(mid(haoma,7,1)) and " +
		"floor(mid(haoma,3,1))<>floor(mid(haoma,2,1)) )" +
		"or ( floor(mid(haoma,4,4)) mod 1111 = 0 and " +
		"floor(mid(haoma,7,1))<>floor(mid(haoma,8,1)) and " +
		"floor(mid(haoma,4,1))<>floor(mid(haoma,3,1)) )" +
		"or ( floor(mid(haoma,5,4)) mod 1111 = 0 and " +
		"floor(mid(haoma,8,1))<>floor(mid(haoma,9,1)) and " +
		"floor(mid(haoma,5,1))<>floor(mid(haoma,4,1)) )" +
		"or ( floor(mid(haoma,6,4)) mod 1111 = 0 and " +
		"floor(mid(haoma,9,1))<>floor(mid(haoma,10,1)) and " +
		"floor(mid(haoma,6,1))<>floor(mid(haoma,5,1)) )" +
		"or ( floor(mid(haoma,7,4)) mod 1111 = 0 and " +
		"floor(mid(haoma,10,1))<>floor(mid(haoma,11,1)) and " +
		"floor(mid(haoma,7,1))<>floor(mid(haoma,6,1)) )" +
		"or ( floor(mid(haoma,8,4)) mod 1111 = 0 and " +
		"floor(mid(haoma,8,1))<>floor(mid(haoma,7,1)) )" +
		")"

	tezhengs["3"] = make(map[string]string)
	tezhengs["3"]["key"] = "AAA"
	tezhengs["3"]["value"] = "3"
	tezhengs["3"]["sql"] = "(" +
		"( floor(mid(haoma,2,3)) mod 111 = 0 and " +
		"floor(mid(haoma,4,1))<>floor(mid(haoma,5,1)) )" +
		"or ( floor(mid(haoma,3,3)) mod 111 = 0 and " +
		"floor(mid(haoma,5,1))<>floor(mid(haoma,6,1)) and " +
		"floor(mid(haoma,3,1))<>floor(mid(haoma,2,1)) )" +
		"or ( floor(mid(haoma,4,3)) mod 111 = 0 and " +
		"floor(mid(haoma,6,1))<>floor(mid(haoma,7,1)) and " +
		"floor(mid(haoma,4,1))<>floor(mid(haoma,3,1)) )" +
		"or ( floor(mid(haoma,5,3)) mod 111 = 0 and " +
		"floor(mid(haoma,7,1))<>floor(mid(haoma,8,1)) and " +
		"floor(mid(haoma,5,1))<>floor(mid(haoma,4,1)) )" +
		"or ( floor(mid(haoma,6,3)) mod 111 = 0 and " +
		"floor(mid(haoma,8,1))<>floor(mid(haoma,9,1)) and " +
		"floor(mid(haoma,6,1))<>floor(mid(haoma,5,1)) )" +
		"or ( floor(mid(haoma,7,3)) mod 111 = 0 and " +
		"floor(mid(haoma,9,1))<>floor(mid(haoma,10,1)) and " +
		"floor(mid(haoma,7,1))<>floor(mid(haoma,6,1)) )" +
		"or ( floor(mid(haoma,8,3)) mod 111 = 0 and " +
		"floor(mid(haoma,10,1))<>floor(mid(haoma,11,1)) and " +
		"floor(mid(haoma,8,1))<>floor(mid(haoma,7,1)) )" +
		"or ( floor(mid(haoma,9,3)) mod 111 = 0 and " +
		"floor(mid(haoma,9,1))<>floor(mid(haoma,8,1)) )" +
		")"

	tezhengs["4"] = make(map[string]string)
	tezhengs["4"]["key"] = "AA"
	tezhengs["4"]["value"] = "4"
	tezhengs["4"]["sql"] = "(" +
		"( floor(mid(haoma,2,2)) mod 11 = 0 and " +
		"floor(mid(haoma,3,1))<>floor(mid(haoma,4,1)) )" +
		"or ( floor(mid(haoma,3,2)) mod 11 = 0 and " +
		"floor(mid(haoma,4,1))<>floor(mid(haoma,5,1)) and " +
		"floor(mid(haoma,3,1))<>floor(mid(haoma,2,1)) )" +
		"or ( floor(mid(haoma,4,2)) mod 11 = 0 and " +
		"floor(mid(haoma,5,1))<>floor(mid(haoma,6,1)) and " +
		"floor(mid(haoma,4,1))<>floor(mid(haoma,3,1)) )" +
		"or ( floor(mid(haoma,5,2)) mod 11 = 0 and " +
		"floor(mid(haoma,6,1))<>floor(mid(haoma,7,1)) and " +
		"floor(mid(haoma,5,1))<>floor(mid(haoma,4,1)) )" +
		"or ( floor(mid(haoma,6,2)) mod 11 = 0 and " +
		"floor(mid(haoma,7,1))<>floor(mid(haoma,8,1)) and " +
		"floor(mid(haoma,6,1))<>floor(mid(haoma,5,1)) )" +
		"or ( floor(mid(haoma,7,2)) mod 11 = 0 and " +
		"floor(mid(haoma,8,1))<>floor(mid(haoma,9,1)) and " +
		"floor(mid(haoma,7,1))<>floor(mid(haoma,6,1)) )" +
		"or ( floor(mid(haoma,8,2)) mod 11 = 0 and " +
		"floor(mid(haoma,9,1))<>floor(mid(haoma,10,1)) and " +
		"floor(mid(haoma,8,1))<>floor(mid(haoma,7,1)) )" +
		"or ( floor(mid(haoma,9,2)) mod 11 = 0 and " +
		"floor(mid(haoma,10,1))<>floor(mid(haoma,11,1)) and " +
		"floor(mid(haoma,9,1))<>floor(mid(haoma,8,1)) )" +
		"or ( floor(mid(haoma,10,2)) mod 11 = 0 and " +
		"floor(mid(haoma,10,1))<>floor(mid(haoma,9,1)) )" +
		")"

	tezhengs["5"] = make(map[string]string)
	tezhengs["5"]["key"] = "AABB"
	tezhengs["5"]["value"] = "5"
	tezhengs["5"]["sql"] = "(" +
		"( floor(mid(haoma,2,2)) mod 11 = 0 and " +
		"floor(mid(haoma,4,2)) mod 11 =0 and " +
		"floor(mid(haoma,2,2))<>floor(mid(haoma,4,2)) and " +
		"floor(mid(haoma,5,1))<>floor(mid(haoma,6,1)) and " +
		"floor(mid(haoma,2,1))<>floor(mid(haoma,1,1)) )" +
		"or ( floor(mid(haoma,3,2)) mod 11 = 0 and " +
		"floor(mid(haoma,5,2)) mod 11 =0 and " +
		"floor(mid(haoma,3,2))<>floor(mid(haoma,5,2)) and " +
		"floor(mid(haoma,6,1))<>floor(mid(haoma,7,1)) and " +
		"floor(mid(haoma,3,1))<>floor(mid(haoma,2,1)) )" +
		"or ( floor(mid(haoma,4,2)) mod 11 = 0 and " +
		"floor(mid(haoma,6,2)) mod 11 =0 and " +
		"floor(mid(haoma,4,2))<>floor(mid(haoma,6,2)) and " +
		"floor(mid(haoma,7,1))<>floor(mid(haoma,8,1)) and " +
		"floor(mid(haoma,4,1))<>floor(mid(haoma,3,1)) )" +
		"or ( floor(mid(haoma,5,2)) mod 11 = 0 and " +
		"floor(mid(haoma,7,2)) mod 11 =0 and " +
		"floor(mid(haoma,5,2))<>floor(mid(haoma,7,2)) and " +
		"floor(mid(haoma,8,1))<>floor(mid(haoma,9,1)) and " +
		"floor(mid(haoma,5,1))<>floor(mid(haoma,4,1)) )" +
		"or ( floor(mid(haoma,6,2)) mod 11 = 0 and " +
		"floor(mid(haoma,8,2)) mod 11 =0 and " +
		"floor(mid(haoma,6,2))<>floor(mid(haoma,8,2)) and " +
		"floor(mid(haoma,9,1))<>floor(mid(haoma,10,1)) and " +
		"floor(mid(haoma,6,1))<>floor(mid(haoma,5,1)) )" +
		"or ( floor(mid(haoma,7,2)) mod 11 = 0 and " +
		"floor(mid(haoma,9,2)) mod 11 =0 and " +
		"floor(mid(haoma,7,2))<>floor(mid(haoma,9,2)) and " +
		"floor(mid(haoma,10,1))<>floor(mid(haoma,11,1)) and " +
		"floor(mid(haoma,7,1))<>floor(mid(haoma,6,1)) )" +
		"or ( floor(mid(haoma,8,2)) mod 11 = 0 and " +
		"floor(mid(haoma,10,2)) mod 11 =0 and " +
		"floor(mid(haoma,8,2))<>floor(mid(haoma,10,2)) and " +
		"floor(mid(haoma,8,1))<>floor(mid(haoma,7,1)) )" +
		")"

	tezhengs["6"] = make(map[string]string)
	tezhengs["6"]["key"] = "AABBB"
	tezhengs["6"]["value"] = "6"
	tezhengs["6"]["sql"] = "(" +
		"floor(right(haoma,3)) mod 111 = 0 and " +
		"floor(mid(haoma,length(haoma)-4,2)) mod 11 =0 and " +
		"floor(right(haoma,3))<>floor(mid(haoma,length(haoma)-4,2))  and " +
		"floor(mid(right(haoma,4),1,1))<>floor(mid(right(haoma,3),1,1))" +
		")"

	tezhengs["7"] = make(map[string]string)
	tezhengs["7"]["key"] = "AABBCC"
	tezhengs["7"]["value"] = "7"
	tezhengs["7"]["sql"] = "(" +
		"floor(right(haoma,2)) mod 11 = 0 and " +
		"floor(mid(haoma,length(haoma)-3,2)) mod 11 =0 and " +
		"floor(mid(haoma,length(haoma)-5,2)) mod 11 =0 and " +
		"floor(right(haoma,2))<>floor(mid(haoma,length(haoma)-3,2))" +
		"<>floor(mid(haoma,length(haoma)-5,2))" +
		")"

	tezhengs["8"] = make(map[string]string)
	tezhengs["8"]["key"] = "ABAB"
	tezhengs["8"]["value"] = "8"
	tezhengs["8"]["sql"] = "(" +
		"mid(haoma,length(haoma)-3,1)=mid(haoma,length(haoma)-1,1) and " +
		"mid(haoma,length(haoma)-2,1)=right(haoma,1) and " +
		"floor(right(haoma,4)) mod 1111<>0  and " +
		"floor(mid(right(haoma,6),1,1))<>floor(mid(right(haoma,4),1,1))" +
		")"

	tezhengs["9"] = make(map[string]string)
	tezhengs["9"]["key"] = "ABABAB"
	tezhengs["9"]["value"] = "9"
	tezhengs["9"]["sql"] = "(" +
		"mid(haoma,length(haoma)-3,1)=mid(haoma,length(haoma)-1,1) and " +
		"mid(haoma,length(haoma)-1,1)=mid(haoma,length(haoma)-5,1) and " +
		"mid(haoma,length(haoma)-2,1)=right(haoma,1) and " +
		"mid(haoma,length(haoma)-4,1)=right(haoma,1)" +
		")"

	tezhengs["10"] = make(map[string]string)
	tezhengs["10"]["key"] = "ABC"
	tezhengs["10"]["value"] = "10"
	tezhengs["10"]["sql"] = "(" +
		"(mid(haoma,length(haoma)-1,1)-1 = floor(right(haoma,1)) and " +
		"mid(haoma,length(haoma)-2,1)-2 = floor(right(haoma,1))) " +
		"or  (mid(haoma,length(haoma)-1,1)+1 = floor(right(haoma,1)) and " +
		"mid(haoma,length(haoma)-2,1)+2 = floor(right(haoma,1)))" +
		")"

	tezhengs["11"] = make(map[string]string)
	tezhengs["11"]["key"] = "ABCD"
	tezhengs["11"]["value"] = "11"
	tezhengs["11"]["sql"] = "(" +
		"(mid(haoma,length(haoma)-1,1)-1 = floor(right(haoma,1)) and " +
		"mid(haoma,length(haoma)-2,1)-2 = floor(right(haoma,1)) and " +
		"mid(haoma,length(haoma)-3,1)-3 = floor(right(haoma,1))) " +
		"or (mid(haoma,length(haoma)-1,1)+1 = floor(right(haoma,1)) and " +
		"mid(haoma,length(haoma)-2,1)+2 = floor(right(haoma,1)) and " +
		"mid(haoma,length(haoma)-3,1)+3 = floor(right(haoma,1)))" +
		")"

	tezhengs["12"] = make(map[string]string)
	tezhengs["12"]["key"] = "ABCABC"
	tezhengs["12"]["value"] = "12"
	tezhengs["12"]["sql"] = "(" +
		"floor(mid(right(haoma,6),4,1))=floor(mid(right(haoma,6),1,1)) and " +
		"floor(mid(right(haoma,6),5,1))= floor(mid(right(haoma,6),2,1)) and " +
		"floor(mid(right(haoma,6),6,1))=floor(mid(right(haoma,6),3,1))" +
		")"

	tezhengs["13"] = make(map[string]string)
	tezhengs["13"]["key"] = "ABCDABCD"
	tezhengs["13"]["value"] = "13"
	tezhengs["13"]["sql"] = "(" +
		"floor(mid(right(haoma,8),5,1))=floor(mid(right(haoma,8),1,1)) and " +
		"floor(mid(right(haoma,8),6,1))= floor(mid(right(haoma,8),2,1)) and " +
		"floor(mid(right(haoma,8),7,1))=floor(mid(right(haoma,8),3,1)) and " +
		"floor(mid(right(haoma,8),8,1))= floor(mid(right(haoma,8),4,1))" +
		")"

	tezhengs["14"] = make(map[string]string)
	tezhengs["14"]["key"] = "AABCC"
	tezhengs["14"]["value"] = "14"
	tezhengs["14"]["sql"] = "(" +
		"floor(right(haoma,2)) mod 11 = 0 and " +
		"mid(haoma,length(haoma)-4,2) mod 11 = 0  and mid(haoma,length(haoma)-3,1) and " +
		"floor(right(haoma,3)) mod 111 <> 0" +
		")"

	tezhengs["15"] = make(map[string]string)
	tezhengs["15"]["key"] = "AAAB"
	tezhengs["15"]["value"] = "15"
	tezhengs["15"]["sql"] = "(" +
		"mid(haoma,length(haoma)-3,3) mod 111 = 0 and " +
		"floor(right(haoma,1)) <> mid(haoma,length(haoma)-2,1)" +
		")"

	tezhengs["16"] = make(map[string]string)
	tezhengs["16"]["key"] = "ABAC"
	tezhengs["16"]["value"] = "16"
	tezhengs["16"]["sql"] = "(" +
		"mid(haoma,length(haoma)-1,1) = mid(haoma,length(haoma)-3,1) and " +
		"mid(haoma,length(haoma)-1,1) <> mid(haoma,length(haoma)-2,1)  and " +
		"mid(haoma,length(haoma),1) <> mid(haoma,length(haoma)-2,1)" +
		")"

	tezhengs["17"] = make(map[string]string)
	tezhengs["17"]["key"] = "BACA"
	tezhengs["17"]["value"] = "17"
	tezhengs["17"]["sql"] = "(" +
		"mid(haoma,length(haoma),1) = mid(haoma,length(haoma)-2,1) and " +
		"mid(haoma,length(haoma)-1,1) <> mid(haoma,length(haoma)-2,1) and " +
		"mid(haoma,length(haoma)-1,1) <> mid(haoma,length(haoma)-3,1)" +
		")"

	tezhengs["18"] = make(map[string]string)
	tezhengs["18"]["key"] = "ABBA"
	tezhengs["18"]["value"] = "18"
	tezhengs["18"]["sql"] = "(" +
		"floor(mid(haoma,length(haoma)-2,2)) mod 11 =0 and " +
		"floor(right(haoma,1)) = mid(haoma,length(haoma)-3,1) and " +
		"floor(right(haoma,1)) <> floor(mid(haoma,length(haoma)-2,1))" +
		")"

	tezhengs["19"] = make(map[string]string)
	tezhengs["19"]["key"] = "ABCAB"
	tezhengs["19"]["value"] = "19"
	tezhengs["19"]["sql"] = "(" +
		"floor(mid(right(haoma,4),1,1))=floor(mid(right(haoma,1),1,1)) and " +
		"floor(mid(right(haoma,5),1,1))= floor(mid(right(haoma,2),1,1)) and " +
		"floor(mid(right(haoma,3),1,1))<> floor(mid(right(haoma,2),1,1))" +
		")"

	tezhengs["20"] = make(map[string]string)
	tezhengs["20"]["key"] = "ABBABB"
	tezhengs["20"]["value"] = "20"
	tezhengs["20"]["sql"] = "(" +
		"floor(mid(right(haoma,2),1,1))=floor(mid(right(haoma,1),1,1)) and " +
		"floor(mid(right(haoma,5),1,1))= floor(mid(right(haoma,4),1,1)) and " +
		"floor(mid(right(haoma,3),1,1))<> floor(mid(right(haoma,4),1,1)) and " +
		"floor(mid(right(haoma,5),1,1))<> floor(mid(right(haoma,6),1,1)) and " +
		"floor(mid(right(haoma,3),1,1))<> floor(mid(right(haoma,2),1,1)) and " +
		"floor(mid(right(haoma,6),1,1))=floor(mid(right(haoma,3),1,1))" +
		")"

	tezhengs["21"] = make(map[string]string)
	tezhengs["21"]["key"] = "*A*A*A*A"
	tezhengs["21"]["value"] = "21"
	tezhengs["21"]["sql"] = "(" +
		"floor(mid(right(haoma,3),1,1))=floor(mid(right(haoma,1),1,1)) and " +
		"floor(mid(right(haoma,5),1,1))=floor(mid(right(haoma,3),1,1)) and " +
		"floor(mid(right(haoma,7),1,1))=floor(mid(right(haoma,5),1,1)) and " +
		"floor(mid(right(haoma,2),1,1))<>floor(mid(right(haoma,1),1,1))" +
		")"

	tezhengs["22"] = make(map[string]string)
	tezhengs["22"]["key"] = "**AA**AA"
	tezhengs["22"]["value"] = "22"
	tezhengs["22"]["sql"] = "(" +
		"floor(mid(right(haoma,2),1,1))=floor(mid(right(haoma,1),1,1)) and " +
		"floor(mid(right(haoma,6),1,1))= floor(mid(right(haoma,5),1,1))  and " +
		"floor(mid(right(haoma,4),1,1))<>floor(mid(right(haoma,3),1,1))" +
		")"

	tezhengs["23"] = make(map[string]string)
	tezhengs["23"]["key"] = "**AB**AB"
	tezhengs["23"]["value"] = "23"
	tezhengs["23"]["sql"] = "(" +
		"floor(mid(right(haoma,5),1,1))=floor(mid(right(haoma,1),1,1)) and " +
		"floor(mid(right(haoma,6),1,1))= floor(mid(right(haoma,2),1,1))  and " +
		"floor(mid(right(haoma,4),1,1))<>floor(mid(right(haoma,2),1,1))" +
		")"

	//判断name字段是否存在
	//v,ok := users[name]
	if tezhengs[tezheng] != nil {
		//存在
		v = tezhengs[tezheng]
	} else {
		//不存在
		v = tezhengs["nil"]
	}
	return
}

