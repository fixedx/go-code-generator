package utils

import (
	"os"
	"path"

	"github.com/spf13/viper"
)

type ControllerTemplateData struct {
	LowerCaseName string
	UpperCaseName string
	ProjectName   string
}

func CreateController() {
	var base, _ = os.Getwd()

	lowerCaseName, upperCaseName := GetLowerAndUpperName(os.Args[1])
	var projectName = viper.GetString("project.name")
	var templateFile = path.Join(base, viper.GetString("controller.template"))
	var targetFile = path.Join(base, viper.GetString("controller.target"))
	targetFile = ReplaceVariable(targetFile, map[string]string{
		"lowerCaseName": lowerCaseName,
	})
	RenderTemplate(templateFile, targetFile, &ControllerTemplateData{
		LowerCaseName: lowerCaseName,
		UpperCaseName: upperCaseName,
		ProjectName:   projectName,
	})

}
