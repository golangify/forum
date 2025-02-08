package templatesutils

import (
	"fmt"
	"forum/config"
	"html/template"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func Init(config *config.Config, engine *gin.Engine) error {
	var templateFiles []string
	err := filepath.Walk("templates/", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			fmt.Println(path)
			templateFiles = append(templateFiles, path)
		}
		return nil
	})
	if err != nil {
		return err
	}

	engine.SetFuncMap(template.FuncMap{
		"timeAgo": timeAgo,
	})
	engine.LoadHTMLFiles(templateFiles...)

	return nil
}
