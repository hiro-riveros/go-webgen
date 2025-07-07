package scaffolder

import (
	"path/filepath"
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
