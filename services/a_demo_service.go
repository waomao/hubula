package services

//数据处理（包括数据库，也包括缓存等其他形式数据）

import (
	"github.com/waomao/hubula/common"
	"github.com/waomao/hubula/dao"
	"github.com/waomao/hubula/datasource"
	"github.com/waomao/hubula/models"
)

type ADemoService interface {
	// GetAll 列表查询
	GetAll(sqlwhere *common.SqlWhere) (*common.SqlReturn, error)
	// GetById 获取单条记录
	GetById(id int64) (*models.ADemo, error)
	// CountAll 统计
	CountAll() int64

	// Create 添加单条记录
	Create(data *models.ADemo) (int64,error)
	// Update 修改单条记录
	Update(data *models.ADemo, columns []string) (int64,error)
	// RuanDelete 软删除单条记录
	RuanDelete(id int64) (int64, error)
	// Delete 删除单条记录
	Delete(id int64) (int64, error)

	// GetWhere Sql语句
	GetWhere(sql string) []models.ADemo
}

type aDemoService struct {
	dao *dao.ADemoDao
}

func NewADemoService() ADemoService {
	return &aDemoService{
		dao: dao.NewADemoDao(datasource.MasterEngine()),
	}
}

func (s *aDemoService) GetAll (sqlwhere *common.SqlWhere) (*common.SqlReturn, error)  {
	return s.dao.GetAll(sqlwhere)
}

// GetById 获取单条记录
func (s *aDemoService) GetById(id int64) (*models.ADemo, error) {
	return s.dao.GetById(id)
}

// CountAll 统计
func (s *aDemoService) CountAll() int64 {
	return s.dao.CountAll()
}

// Create 添加单条记录
func (s *aDemoService) Create(data *models.ADemo) (int64,error) {
	// 先更新缓存
	//s.updateByCache(data, nil)
	// 再更新数据库
	return s.dao.Create(data)
}

// Update 修改单条记录
func (s *aDemoService) Update(data *models.ADemo, columns []string) (int64,error) {
	// 先更新缓存
	//s.updateByCache(data, columns)
	// 再更新数据库
	return s.dao.Update(data, columns)
}

// RuanDelete 软删除单条记录
func (s *aDemoService) RuanDelete(id int64) (int64, error) {
	// 先更新缓存
	//data := &models.Users{Id: id}
	//s.updateByCache(data, nil)
	// 再更新数据库
	return s.dao.RuanDelete(id)
}

// Delete 删除单条记录
func (s *aDemoService) Delete(id int64) (int64, error) {
	// 先更新缓存
	//data := &models.Users{Id: id}
	//s.updateByCache(data, nil)
	// 再更新数据库
	return s.dao.Delete(id)
}

// GetWhere Sql语句
func (s *aDemoService) GetWhere(sql string) []models.ADemo {
	return s.dao.GetWhere(sql)
}
