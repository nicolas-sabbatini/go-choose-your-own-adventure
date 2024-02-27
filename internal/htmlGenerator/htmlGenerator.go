package htmlGenerator

import (
	"html/template"
)

var ChapterTemplate *template.Template
var IndexTemplate *template.Template

func LoadTemplates(chapterTemplatePath string, indexTemplatePath string) error {
	var err error

	ChapterTemplate, err = template.ParseFiles(chapterTemplatePath)
	if err != nil {
		return err
	}

	IndexTemplate, err = template.ParseFiles(indexTemplatePath)
	if err != nil {
		return err
	}

	return nil
}
