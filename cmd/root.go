package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "go-webgen",
	Short: "Generador de proyectos web con Go",
	Long:  `Un generador de boilerplates para proyectos web, con opciones de frameworks, docker, etc.`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(generateCmd)
}
