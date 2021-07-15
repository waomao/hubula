package common

import (
	"encoding/json"
	"fmt"
	"github.com/waomao/hubula/models"
)

/**
递归获取树形菜单
*/

//树形结构
func Tree(list []*models.ADemo) string {
	fmt.Println(list)
	data := buildData(list)
	result := makeTreeCore(0, data)

	body, err := json.Marshal(result)
	if err != nil {
		fmt.Println(err)
	}
	return string(body)
}

//以下为用到的方法
//编译数据
func buildData(list []*models.ADemo) map[int]map[int]*models.ADemo {
	var data map[int]map[int]*models.ADemo = make(map[int]map[int]*models.ADemo)
	for _, v := range list {
		id := int(v.Id)
		fid := v.SuperiorId
		if _, ok := data[fid]; !ok {
			data[fid] = make(map[int]*models.ADemo)
		}
		data[fid][id] = v
	}
	return data
}

//生成树
func makeTreeCore(index int, data map[int]map[int]*models.ADemo) []*models.ADemo {
	tmp := make([]*models.ADemo, 0)
	for id, item := range data[index] {
		if data[id] != nil {
			item.Children = makeTreeCore(id, data)
		}
		tmp = append(tmp, item)
	}
	return tmp
}

