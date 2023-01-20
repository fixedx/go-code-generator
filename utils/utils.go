package utils

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"text/template"
)

func FileExists(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		return os.IsExist(err)
	}
	return true
}

func ReplaceVariable(s string, data map[string]string) string {
	reg := regexp.MustCompile(`\{(.*?)\}`)
	return reg.ReplaceAllStringFunc(s, func(a string) string {
		var arr = strings.Split(a, "")
		prop := strings.Join(arr[1:len(arr)-1], "")
		return data[prop]
	})
}

func GetLowerAndUpperName(name string) (lowerCaseName string, upperCaseName string) {
	lowerCaseName = strings.ToLower(name[0:1]) + name[1:]
	upperCaseName = strings.ToUpper(name[0:1]) + name[1:]
	return
}

func RenderTemplate(templateFile string, target string, data interface{}) {
	t, err := template.ParseFiles(templateFile)
	if err != nil {
		fmt.Println(err)
	}
	var buf bytes.Buffer
	err = t.Execute(&buf, data)
	if err != nil {
		fmt.Println(err)
		return
	}
	if !FileExists(filepath.Dir(target)) {
		err := os.MkdirAll(filepath.Dir(target), 0777)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	err = os.WriteFile(target, buf.Bytes(), 0777)
	if err != nil {
		fmt.Println(err)
		return
	}
}
