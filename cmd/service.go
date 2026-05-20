/*
Copyright (c) 2026
*/
package cmd

import (
	"fmt"
	"sort"
	"strings"

	"github.com/AlecAivazis/survey/v2"
	"github.com/HTTPauloGoncalves/Deploy-Hub/internal"
	"github.com/HTTPauloGoncalves/Deploy-Hub/internal/config"
	"github.com/spf13/cobra"
)

var serviceCmd = &cobra.Command{
	Use:   "service",
	Short: "Gerencia servicos de deploy",
}

var serviceAddCmd = &cobra.Command{
	Use:   "add",
	Short: "Adiciona um novo servico de deploy",
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
			fmt.Println("Nenhum servidor cadastrado. Cadastre um servidor antes de adicionar um servico.")
			return
		}

		serverOptions := sortedKeys(servers)

		survey.AskOne(&survey.Input{
			Message: "Nome do servico:",
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

			if strings.EqualFold(command, "done") {
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

		fmt.Println("Servico adicionado com sucesso:", name)
	},
}

var serviceDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Deleta um servico de deploy",
	Run: func(cmd *cobra.Command, args []string) {
		var name string

		services, err := internal.ListServices()
		if err != nil {
			fmt.Println("Erro:", err)
			return
		}

		if len(services) == 0 {
			fmt.Println("Nenhum servico cadastrado.")
			return
		}

		serviceOptions := sortedKeys(services)

		err = survey.AskOne(&survey.Select{
			Message: "Selecione o servico para deletar:",
			Options: serviceOptions,
		}, &name)
		if err != nil {
			fmt.Println("Erro:", err)
			return
		}

		err = internal.DeleteService(name)
		if err != nil {
			fmt.Println("Erro:", err)
			return
		}

		fmt.Println("Servico deletado com sucesso:", name)
	},
}

var serviceListCmd = &cobra.Command{
	Use:   "list",
	Short: "Lista os servicos de deploy cadastrados",
	Run: func(cmd *cobra.Command, args []string) {
		services, err := internal.ListServices()
		if err != nil {
			fmt.Println("Erro:", err)
			return
		}

		if len(services) == 0 {
			fmt.Println("Nenhum servico cadastrado.")
			return
		}

		fmt.Println("Servicos de Deploy Cadastrados:")
		serviceNames := sortedKeys(services)
		for _, name := range serviceNames {
			service := services[name]
			fmt.Printf("- %s (Servidor: %s, Path: %s)\n", name, service.Server, service.Path)
			for _, cmd := range service.Commands {
				fmt.Printf("  - Comando: %s\n", cmd)
			}
		}
	},
}

var serviceUpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "Atualiza um servico de deploy",
	Run: func(cmd *cobra.Command, args []string) {
		var name string
		var server string
		var path string
		var commands []string
		keepCommands := true

		services, err := internal.ListServices()
		if err != nil {
			fmt.Println("Erro:", err)
			return
		}

		if len(services) == 0 {
			fmt.Println("Nenhum servico cadastrado.")
			return
		}

		servers, err := internal.ListServers()
		if err != nil {
			fmt.Println("Erro:", err)
			return
		}

		if len(servers) == 0 {
			fmt.Println("Nenhum servidor cadastrado. Cadastre um servidor antes de atualizar um servico.")
			return
		}

		err = survey.AskOne(&survey.Select{
			Message: "Selecione o servico para atualizar:",
			Options: sortedKeys(services),
		}, &name)
		if err != nil {
			fmt.Println("Erro:", err)
			return
		}

		currentService := services[name]

		err = survey.AskOne(&survey.Select{
			Message: "Servidor:",
			Options: sortedKeys(servers),
			Default: currentService.Server,
		}, &server)
		if err != nil {
			fmt.Println("Erro:", err)
			return
		}

		err = survey.AskOne(&survey.Input{
			Message: "Path na VPS:",
			Default: currentService.Path,
		}, &path)
		if err != nil {
			fmt.Println("Erro:", err)
			return
		}

		if len(currentService.Commands) > 0 {
			fmt.Println("Comandos atuais:")
			for _, command := range currentService.Commands {
				fmt.Printf("- %s\n", command)
			}

			err = survey.AskOne(&survey.Confirm{
				Message: "Manter comandos atuais?",
				Default: true,
			}, &keepCommands)
			if err != nil {
				fmt.Println("Erro:", err)
				return
			}
		} else {
			keepCommands = false
		}

		if keepCommands {
			commands = currentService.Commands
		} else {
			commands, err = askServiceCommands()
			if err != nil {
				fmt.Println("Erro:", err)
				return
			}
		}

		service := config.Service{
			Server:   server,
			Path:     path,
			Commands: commands,
		}

		err = internal.UpdateService(name, service)
		if err != nil {
			fmt.Println("Erro:", err)
			return
		}

		fmt.Println("Servico atualizado com sucesso:", name)
	},
}

func init() {
	rootCmd.AddCommand(serviceCmd)
	serviceCmd.AddCommand(serviceAddCmd)
	serviceCmd.AddCommand(serviceListCmd)
	serviceCmd.AddCommand(serviceUpdateCmd)
	serviceCmd.AddCommand(serviceDeleteCmd)
}

func sortedKeys[T any](items map[string]T) []string {
	keys := make([]string, 0, len(items))
	for key := range items {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	return keys
}

func askServiceCommands() ([]string, error) {
	commands := []string{}

	for {
		var command string

		err := survey.AskOne(&survey.Input{
			Message: "Comando de deploy (digite 'done' para finalizar):",
		}, &command)
		if err != nil {
			return nil, err
		}

		if strings.EqualFold(command, "done") {
			break
		}

		if command != "" {
			commands = append(commands, command)
		}
	}

	return commands, nil
}
