package dao

import (
	"fmt"
	"github.com/waomao/hubula/common"
	"github.com/waomao/hubula/datasource"
	"github.com/waomao/hubula/models"
	"github.com/xormplus/xorm"
	"html/template"
)

type ADemoDao struct {
	//数据库相关的操作 xorm引擎
	engine *xorm.Engine
}

//New Dao 实例化公共方法
func NewADemoDao(engine *xorm.Engine) *ADemoDao {
	return &ADemoDao{
		engine: engine,
	}
}

// NewAdmin 初始化
func (d *ADemoDao) newModel() *models.ADemo {
	return new(models.ADemo)
}

// newMakeDataArr 初始化列表
func (d *ADemoDao) newMakeDataArr() []models.ADemo {
	return make([]models.ADemo, 0)
}

// GetAll 列表查询
//条件 fields字段常和更新一起使用为0查询或更新所有字段 排序 页数 每页条数 返回分页内容 err
func (d *ADemoDao) GetAll(sqlwhere *common.SqlWhere) (*common.SqlReturn, error) {
	//获取符合条件的数据总数
	sessionCount := datasource.Filter(sqlwhere.Conditions)
	defer sessionCount.Close()
	count, err := sessionCount.Count(&models.ADemo{})
	if err != nil {
		fmt.Println(err)
		return nil, common.NewError(err.Error())
	}
	if count == 0 {
		return nil, nil
	}

	//返回 总页数
	po := common.GetPages(sqlwhere,count)
	str := common.H(po)

	//获取数据
	session := datasource.Filter(sqlwhere.Conditions)
	defer session.Close()
	if sqlwhere.OrderBy != "" {
		session.OrderBy(sqlwhere.OrderBy)
	}
	session.Limit(int(po.PageSize), int((po.Currentpage - 1) * sqlwhere.PageSize))
	if len(sqlwhere.Fields) == 0 {
		//更新所有字段
		session.AllCols()
	}

	data := d.newMakeDataArr()

	err = session.Find(&data)
	if err != nil {
		fmt.Println(err)
		return nil, common.NewError(err.Error())
	}

	//整合需要输出的内容
	sqlR := new(common.SqlReturn)

	sqlR.Data = make([]interface{}, len(data))
	for y, x := range data {
		sqlR.Data[y] = x
	}

	sqlR.Str = template.HTML(str)
	sqlR.Page = po.Currentpage
	sqlR.PageSize = po.PageSize
	sqlR.TotalCount = count
	sqlR.TotalPage = po.TotalPage
	sqlR.Href = po.Href
	return sqlR, nil
}

// GetById 获取单条记录
func (d *ADemoDao) GetById(id int64) (*models.ADemo, error) {
	m := d.newModel()
	//fmt.Println(id)
	m.Id = id

	s, err := d.engine.Get(m)
	if err == nil {
		if s {
			return m,nil
		}
		return nil, common.NewError("不存在")
	}
	return nil, err
}

// CountAll 统计
func (d *ADemoDao) CountAll() int64 {
	m := d.newModel()
	num, err := d.engine.Count(m)
	if err != nil {
		return 0
	} else {
		return num
	}
}

// Create 添加单条记录
func (d *ADemoDao) Create(data *models.ADemo) (int64,error) {
	num, err := d.engine.InsertOne(data)
	return num,err
}

// Update 修改单条记录
func (d *ADemoDao) Update(data *models.ADemo, columns []string) (int64,error) {
	num, err := d.engine.ID(data.Id).MustCols(columns...).Update(data)
	return num,err
}

// RuanDelete 软删除单条记录
func (d *ADemoDao) RuanDelete(id int64) (int64, error) {
	m := d.newModel()
	m.Id = id
	m.IsDel = 1

	num, err := d.engine.ID(m.Id).Update(m)
	if err == nil {
		return num, nil
	}
	return num, err
}

// Delete 删除单条记录
func (d *ADemoDao) Delete(id int64) (int64, error) {
	m := d.newModel()
	m.Id = id

	num, err := d.engine.Delete(m)
	if err == nil {
		return num, nil
	}
	return num, err
}

// GetWhere Sql语句
//ll := "12"
//sql := "SELECT * FROM `lianxiren` WHERE `tel` LIKE '%" + ll + "%' "
//where := c.ServiceADemo.GetWhere(sql)
func (d *ADemoDao) GetWhere(sql string) []models.ADemo {
	datalist := d.newMakeDataArr()
	err := d.engine.SQL(sql).Find(&datalist)
	if err != nil {
		return datalist
	} else {
		return datalist
	}
}

/**
递归获取树形菜单
*/

func (d *ADemoDao) GetMenu(pid int) []*models.ADemo {
	menu := make([]models.ADemo, 0)
	err := d.engine.Where("superior_id =?", pid).Find(&menu)
	if err != nil || len(menu) < 1 {
		return nil
	}

	WebsiteMap := make([]*models.ADemo, 0)
	for _, v := range menu {
		fmt.Println(v.Id)

		child := d.GetMenu(int(v.Id))
		node := &models.ADemo{
			Id:            v.Id,
			ToId:          v.ToId,
			LoginName:     v.LoginName,
			Password:      v.Password,
			Vsername:      v.Vsername,
			Mobile:        v.Mobile,
			Email:         v.Email,
			GenTime:       v.GenTime,
			LoginTime:     v.LoginTime,
			LastLoginTime: v.LastLoginTime,
			Count:         v.Count,
			IsDel:         v.IsDel,
			SuperiorId:    v.SuperiorId,
			Power:         v.Power,
		}
		node.Children = child
		WebsiteMap = append(WebsiteMap, node)
	}
	return WebsiteMap
}