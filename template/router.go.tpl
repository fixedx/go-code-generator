package router

import (
	"{{ .ProjectName }}/controller"

	"github.com/gin-gonic/gin"
)

func {{.UpperCaseName}}Router(r *gin.Engine) {
	r.POST("/{{.LowerCaseName}}", controller.Add{{.UpperCaseName}})
	r.GET("/{{.LowerCaseName}}s", controller.Get{{.UpperCaseName}}List)
}
