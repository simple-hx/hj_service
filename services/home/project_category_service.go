package home_service

import (
	home_dao "hj_service/dao/home"
	"hj_service/datasource"
	"hj_service/models/home"
)

// 类
type projectCategoryService struct {
	dao *home_dao.ProjectCategoryDao
}

func (p projectCategoryService) GetProjectCategoryList() []home_model.ProjectCategory {
	return p.dao.GetProjectCategoryList()
}

// 接口
type ProjectCategoryService interface {
	GetProjectCategoryList() []home_model.ProjectCategory
}

// 构造函数
func NewProjectCategoryService() ProjectCategoryService {
	return &projectCategoryService{
		dao: home_dao.NewProjectCategoryDao(datasource.GetDb()),
	}
}