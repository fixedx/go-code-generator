package utils

import (
	"os"
	"path"

	"github.com/spf13/viper"
)

type RouterTemplateData struct {
	Routes        []string
	LowerCaseName string
	UpperCaseName string
	ProjectName   string
}

func CreateRouter() {
	var base, _ = os.Getwd()
	name := os.Args[1]
	lowerCaseName, upperCaseName := GetLowerAndUpperName(name)
	var projectName = viper.GetString("project.name")
	var templateFile = path.Join(base, viper.GetString("router.template"))
	var targetFile = path.Join(base, viper.GetString("router.target"))
	targetFile = ReplaceVariable(targetFile, map[string]string{
		"lowerCaseName": lowerCaseName,
	})
	RenderTemplate(templateFile, targetFile, &RouterTemplateData{
		LowerCaseName: lowerCaseName,
		UpperCaseName: upperCaseName,
		ProjectName:   projectName,
	})
	// // generate the router index
	// routes := strings.Split(viper.GetString("routes"), ",")
	// templateFile = path.Join(base, viper.GetString("router.index.template"))
	// targetFile = path.Join(base, viper.GetString("router.index.target"))
	// RenderTemplate(templateFile, targetFile, &RouterTemplateData{Routes: routes})
}
