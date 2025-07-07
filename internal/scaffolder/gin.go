package scaffolder

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

type GinConfig struct {
	ProjectName   string
	ProjectModule string
}

func GenerateGinApp(projectPath string, config GinConfig) error {
	folders := []string{
		filepath.Join(projectPath, "internal", "api", "dto"),
		filepath.Join(projectPath, "internal", "api", "handlers"),
		filepath.Join(projectPath, "internal", "api", "middlewares"),
		filepath.Join(projectPath, "internal", "api", "routes"),
		filepath.Join(projectPath, "models"),
		filepath.Join(projectPath, "services"),
		filepath.Join(projectPath, "config"),
	}

	for _, folder := range folders {
		if err := os.MkdirAll(folder, 0755); err != nil {
			return err
		}
	}

	files := map[string]string{
		"gin/README.md.tmpl":                                "README.md",
		"gin/go.mod.tmpl":                                   "go.mod",
		"gin/main.go.tmpl":                                  "main.go",
		"gin/config/config.go.tmpl":                         "config/config.go",
		"gin/api/routes/router.go.tmpl":                     "internal/api/routes/router.go",
		"gin/api/handlers/ping.go.tmpl":                     "internal/api/handlers/ping.go",
		"gin/api/middlewares/cors_middleware.go.tmpl":       "internal/api/middlewares/cors_middleware.go",
		"gin/api/middlewares/rate_limit_middleware.go.tmpl": "internal/api/middlewares/rate_limit_middleware.go",
		"gin/api/models/setup.go.tmpl":                      "models/setup.go",
	}

	for tmplName, outputName := range files {
		err := renderTemplateToFileEmbedded(
			ginTemplates,
			"templates/"+tmplName,
			filepath.Join(projectPath, outputName),
			config,
		)
		if err != nil {
			return err
		}
	}

	cmd := exec.Command("go", "mod", "tidy")
	cmd.Dir = projectPath
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("❌ Error al ejectar 'go mod tidy': %s\n", string(output))
		return err
	}

	fmt.Println("✅ Dependencias instaladas correctamente (go mod tidy)")
	return nil
}
