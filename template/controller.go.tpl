package controller

import (
	"{{.ProjectName}}/model"
	"{{.ProjectName}}/services"
	"{{.ProjectName}}/utils"

	"github.com/gin-gonic/gin"
)

func Add{{.UpperCaseName}}(c *gin.Context) {
	var {{.LowerCaseName}} model.{{.UpperCaseName}}
	c.Bind(&{{.LowerCaseName}})
	if {{.LowerCaseName}}Id, err := services.Add{{.UpperCaseName}}(&{{.LowerCaseName}}); err != nil {
		c.JSON(200, utils.FormatResponse[any](500, nil, err.Error()))
	} else {
		c.JSON(200, utils.FormatResponse(0, {{.LowerCaseName}}Id, ""))
	}
}

func Get{{.UpperCaseName}}List(c *gin.Context) {
	var params model.{{.UpperCaseName}}Params
	c.Bind(&params)
	{{.LowerCaseName}}s, count, err := services.Get{{.UpperCaseName}}List(&params)
	if err != nil {
		c.JSON(200, utils.FormatResponse[any](500, nil, err.Error()))
	} else {
		c.JSON(200, utils.FormatPageResponse({{.LowerCaseName}}s, count))
	}
}
