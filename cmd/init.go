/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/HTTPauloGoncalves/Deploy-Hub/internal"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init [project-name]",
	Short: "Inicializa um novo projeto de deploy",
	Long:  `Inicializa um novo projeto de deploy. Isso criará um arquivo de configuração padrão para o projeto, onde você poderá definir seus servidores e serviços.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("Por favor, especifique o nome do projeto.")
			return
		}
		projectName := args[0]
		fmt.Printf("Inicializando projeto %s\n", projectName)

		internal.Init(projectName)

		fmt.Printf("Projeto %s inicializado com sucesso! Edite o arquivo deploy.yaml para configurar seus servidores e serviços.\n", projectName)
	},
}

func init() {
	rootCmd.AddCommand(initCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
