package services

import (
	"fmt"
	"{{.ProjectName}}/model"
	"{{.ProjectName}}/utils"
)

func init() {
	fmt.Println("init {{.LowerCaseName}} service")
	utils.Db.AutoMigrate(&model.{{.UpperCaseName}}{})
}

func Add{{.UpperCaseName}}({{.LowerCaseName}} *model.{{.UpperCaseName}}) (int64, error) {
	result := utils.Db.Create({{.LowerCaseName}})
	return {{.LowerCaseName}}.ID, result.Error
}

func Get{{.UpperCaseName}}List(params *model.{{.UpperCaseName}}Query) ({{.LowerCaseName}}s []model.{{.UpperCaseName}}, count int64, err error) {
	var sql = utils.Db.Model(&model.{{.UpperCaseName}}{})
	if params.PageNum == 0 {
		params.PageNum = 1
	}
	if params.PageSize == 0 {
		params.PageSize = 15
	}

	if params.Title != "" {
		sql = sql.Where("title like ?", "%"+params.Title+"%")
	}

	result := sql.Count(&count)
	if result.Error != nil {
		return nil, 0, result.Error
	}

	result = sql.Find(&{{.LowerCaseName}}s)

	return {{.LowerCaseName}}s, count, result.Error
}
