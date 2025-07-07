/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/hiro-riveros/webgen/internal/scaffolder"
	"github.com/spf13/cobra"
)

var withDocker bool

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate [project name]",
	Short: "Genera un nuevo proyecto web",
	Args:  cobra.ExactArgs(1),
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		projectName := args[0]
		projectPath := filepath.Join(".", projectName)

		err := os.MkdirAll(projectPath, 0755)
		if err != nil {
			fmt.Println("❌ Error al crear proyecto: ", err)
			return
		}

		if withDocker {
			err := scaffolder.GenerateDockerFiles(projectPath, projectName)
			if err != nil {
				fmt.Println("❌ Error al generar archivos Docker:", err)
				return
			}

			fmt.Println("✅ Archivos Docker generados")
		} else {
			fmt.Println("ℹ️ Docker no fue incluido")
		}

	},
}

func init() {
	generateCmd.Flags().BoolVarP(&withDocker, "docker", "d", false, "Incluir Docker y docker-compose")
}
