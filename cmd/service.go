/*
Copyright (c) 2026
*/
package cmd

import (
	"fmt"
	"sort"

	"github.com/AlecAivazis/survey/v2"
	"github.com/HTTPauloGoncalves/Deploy-Hub/internal"
	"github.com/HTTPauloGoncalves/Deploy-Hub/internal/config"
	"github.com/spf13/cobra"
)

var serviceCmd = &cobra.Command{
	Use:   "service",
	Short: "Gerencia serviços de deploy",
}

var serviceAddCmd = &cobra.Command{
	Use:   "add",
	Short: "Adiciona um novo serviço de deploy",
	Run: func(cmd *cobra.Command, args []string) {
		var name string
		var server string
		var path string
		var commands []string

		servers, err := internal.ListServers()
		if err != nil {
			fmt.Println("Erro:", err)
			return
		}

		if len(servers) == 0 {
			fmt.Println("Nenhum servidor cadastrado. Cadastre um servidor antes de adicionar um serviço.")
			return
		}

		serverOptions := make([]string, 0, len(servers))
		for serverName := range servers {
			serverOptions = append(serverOptions, serverName)
		}
		sort.Strings(serverOptions)

		survey.AskOne(&survey.Input{
			Message: "Nome do serviço:",
		}, &name)

		survey.AskOne(&survey.Select{
			Message: "Servidor:",
			Options: serverOptions,
		}, &server)

		survey.AskOne(&survey.Input{
			Message: "Path na VPS:",
		}, &path)

		for {
			var command string

			survey.AskOne(&survey.Input{
				Message: "Comando de deploy (digite 'done' para finalizar):",
			}, &command)

			if command == "done" {
				break
			}

			if command != "" {
				commands = append(commands, command)
			}
		}

		service := config.Service{
			Server:   server,
			Path:     path,
			Commands: commands,
		}

		err = internal.NewService(name, service)
		if err != nil {
			fmt.Println("Erro:", err)
			return
		}

		fmt.Println("Serviço adicionado com sucesso:", name)
	},
}

func init() {
	rootCmd.AddCommand(serviceCmd)
	serviceCmd.AddCommand(serviceAddCmd)
}
