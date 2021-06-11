package services

/**
 * 数据处理（包括数据库，也包括缓存等其他形式数据）
 */
import (
	"github.com/waomao/hubula/dao"
	"github.com/waomao/hubula/datasource"
	"github.com/waomao/hubula/models"
)

//UserService 接口
type UserService interface {
	Get(id int) *models.User
}

//私有 实现接口
type userService struct {
	dao *dao.UserDao
}

//NewUserService 返回接口
func NewUserService() UserService {
	return &userService{
		dao: dao.NewUserDao(datasource.InstanceDbMaster()),
	}
}

func (s *userService) Get(id int) *models.User {

	// 直接读取数据库的方式
	return s.dao.Get(id)

}