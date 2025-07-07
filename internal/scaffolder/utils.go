package scaffolder

import (
	"embed"
	"fmt"
	"os"
	"path/filepath"
	"text/template"
)

func renderTemplateToFileEmbedded(fs embed.FS, templatePath, outputPath string, data interface{}) error {
	tmplContent, err := fs.ReadFile(templatePath)
	if err != nil {
		return fmt.Errorf("read template error: %w", err)
	}

	tmpl, err := template.New(filepath.Base(templatePath)).Parse(string(tmplContent))
	if err != nil {
		return fmt.Errorf("parse template error: %w", err)
	}

	f, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("create file error: %w", err)
	}
	defer f.Close()

	return tmpl.Execute(f, data)
}
