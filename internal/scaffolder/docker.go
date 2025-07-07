package scaffolder

import (
	"embed"
	"fmt"
	"os"
	"path/filepath"
	"text/template"
)

func GenerateDockerFiles(projectPath string, projectName string) error {
	files := map[string]string{
		"Dockerfile.tmpl":             "Dockerfile",
		"Dockerfile.dev.tmpl":         "Dockerfile.dev",
		"docker-compose.yml.tmpl":     "docker-compose.yml",
		"docker-compose.dev.yml.tmpl": "docker-compose.dev.yml",
	}

	for tmplName, outputName := range files {
		err := renderTemplateToFileEmbedded(
			dockerTemplates,
			"templates/docker/"+tmplName,
			filepath.Join(projectPath, outputName),
			map[string]string{"ProjectName": projectName},
		)
		if err != nil {
			return err
		}
	}

	return nil
}

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
