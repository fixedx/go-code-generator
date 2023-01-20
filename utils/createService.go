package utils

import (
	"os"
	"path"

	"github.com/spf13/viper"
)

type ServiceTemplateData struct {
	LowerCaseName string
	UpperCaseName string
	ProjectName   string
}

func CreateService() {
	var base, _ = os.Getwd()
	lowerCaseName, upperCaseName := GetLowerAndUpperName(os.Args[1])
	var projectName = viper.GetString("project.name")
	var templateFile = path.Join(base, viper.GetString("service.template"))
	var targetFile = path.Join(base, viper.GetString("service.target"))
	targetFile = ReplaceVariable(targetFile, map[string]string{
		"lowerCaseName": lowerCaseName,
	})
	RenderTemplate(templateFile, targetFile, &ServiceTemplateData{
		LowerCaseName: lowerCaseName,
		UpperCaseName: upperCaseName,
		ProjectName:   projectName,
	})

}
