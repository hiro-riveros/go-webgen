package scaffolder

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

type ProjectConfig struct {
	ProjectName   string
	ProjectModule string
}

func GenerateAsynqApp(projectPath string, config ProjectConfig) error {
	folders := []string{
		filepath.Join(config.ProjectModule, "jobs"),
		filepath.Join(config.ProjectModule, "models"),
		filepath.Join(config.ProjectModule, "services"),
		filepath.Join(config.ProjectModule, "config"),
	}

	for _, folder := range folders {
		if err := os.MkdirAll(folder, 0755); err != nil {
			return err
		}
	}

	files := map[string]string{
		"asynq/README.md.tmpl":        "README.md",
		"asynq/main.go.tmpl":          "main.go",
		"asynq/go.mod.tmpl":           "go.mod",
		"asynq/config/config.go.tmpl": "config/config.go",
		"asynq/models/setup.go.tmpl":  "models/setup.go",
		"asynq/jobs/server.go.tmpl":   "jobs/server.go",
		"asynq/jobs/tasks.go.tmpl":    "jobs/tasks.go",
	}

	for tmplName, outputName := range files {
		err := renderTemplateToFileEmbedded(
			asynqTemplates,
			"templates/"+tmplName,
			filepath.Join(projectPath, outputName),
			config,
		)
		if err != nil {
			return err
		}
	}

	cmd := exec.Command("go", "mod", "tidy")
	cmd.Dir = config.ProjectModule
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("❌ Error al ejectar 'go mod tidy': %s\n", string(output))
		return err
	}

	fmt.Println("✅ Dependencias instaladas correctamente (go mod tidy)")
	return nil
}
