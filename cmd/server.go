/*
Copyright (c) 2026
*/
package cmd

import (
	"fmt"
	"sort"
	"strconv"

	"github.com/AlecAivazis/survey/v2"
	"github.com/HTTPauloGoncalves/Deploy-Hub/internal"
	"github.com/HTTPauloGoncalves/Deploy-Hub/internal/config"
	"github.com/spf13/cobra"
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Gerencia servidores",
}

var serverAddCmd = &cobra.Command{
	Use:   "add",
	Short: "Adiciona um novo servidor",
	Run: func(cmd *cobra.Command, args []string) {
		var name string
		var host string
		var user string
		var password string
		var auth string
		var portInput string

		survey.AskOne(&survey.Input{
			Message: "Nome do servidor:",
		}, &name)

		survey.AskOne(&survey.Input{
			Message: "Host:",
		}, &host)

		survey.AskOne(&survey.Input{
			Message: "Usuario:",
		}, &user)

		survey.AskOne(&survey.Select{
			Message: "Autenticacao:",
			Options: []string{"password", "key"},
			Default: "password",
		}, &auth)

		if auth == "password" {
			survey.AskOne(&survey.Password{
				Message: "Senha:",
			}, &password)
		}

		survey.AskOne(&survey.Input{
			Message: "Porta:",
			Default: "22",
		}, &portInput)

		port, err := strconv.Atoi(portInput)
		if err != nil {
			fmt.Println("Erro: porta invalida")
			return
		}

		server := config.Server{
			Host:     host,
			User:     user,
			Password: password,
			Auth:     auth,
			Port:     port,
		}

		err = internal.NewServer(name, server)
		if err != nil {
			fmt.Println("Erro:", err)
			return
		}

		fmt.Println("Servidor adicionado com sucesso:", name)
	},
}

var serverListCmd = &cobra.Command{
	Use:   "list",
	Short: "Lista os servidores cadastrados",
	Run: func(cmd *cobra.Command, args []string) {
		servers, err := internal.ListServers()
		if err != nil {
			fmt.Println("Erro:", err)
			return
		}

		if len(servers) == 0 {
			fmt.Println("Nenhum servidor cadastrado.")
			return
		}

		fmt.Println("Servidores cadastrados:")
		for name, server := range servers {
			fmt.Printf("- %s: %s@%s:%d (auth: %s)\n", name, server.User, server.Host, server.Port, server.Auth)
		}
	},
}

var serverDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Deleta um servidor cadastrado",
	Run: func(cmd *cobra.Command, args []string) {
		var name string

		servers, err := internal.ListServers()
		if err != nil {
			fmt.Println("Erro:", err)
			return
		}

		if len(servers) == 0 {
			fmt.Println("Nenhum servidor cadastrado.")
			return
		}

		serversOptions := make([]string, 0, len(servers))
		for serverName := range servers {
			serversOptions = append(serversOptions, serverName)
		}

		sort.Strings(serversOptions)

		survey.AskOne(&survey.Select{
			Message: "Selecione o servidor para deletar:",
			Options: serversOptions,
		}, &name)

		err = internal.DeleteServer(name)
	},
}

var serviceUpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "Atualiza um serviço de deploy",
	Run: func(cmd *cobra.Command, args []string) {
		var name string

		servers, err := internal.ListServers()
		if err != nil {
			fmt.Println("Erro:", err)
			return
		}

		if len(servers) == 0 {
			fmt.Println("Nenhum servidor cadastrado.")
			return
		}

		serversOptions := make([]string, 0, len(servers))

		for serverName := range servers {
			serversOptions = append(serversOptions, serverName)
		}

		sort.Strings(serversOptions)

		survey.AskOne(&survey.Select{
			Message: "Selecione o servidor para atualizar:",
			Options: serversOptions,
		}, &name)

		server, err := internal.GetServer(name)
		if err != nil {
			fmt.Println("Erro:", err)
			return
		}

		survey.AskOne(&survey.Input{
			Message: "Host:",
			Default: server.Host,
		}, &server.Host)

		survey.AskOne(&survey.Input{
			Message: "Usuario:",
			Default: server.User,
		}, &server.User)

		survey.AskOne(&survey.Select{
			Message: "Autenticacao:",
			Options: []string{"password", "key"},
			Default: server.Auth,
		}, &server.Auth)

		if server.Auth == "password" {
			survey.AskOne(&survey.Password{
				Message: "Senha: (deixe em branco para manter a senha atual)",
			}, &server.Password)
		}

		if server.Password == "" {
			server.Password = servers[name].Password
		}

		survey.AskOne(&survey.Input{
			Message: "Porta:",
			Default: strconv.Itoa(server.Port),
		}, &server.Port)

		err = internal.UpdateServer(name, server)
		if err != nil {
			fmt.Println("Erro:", err)
			return
		}

		fmt.Println("Servidor atualizado com sucesso:", name)
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)
	serverCmd.AddCommand(serverAddCmd)
	serverCmd.AddCommand(serverListCmd)
	serverCmd.AddCommand(serverDeleteCmd)
	serverCmd.AddCommand(serviceUpdateCmd)
}
