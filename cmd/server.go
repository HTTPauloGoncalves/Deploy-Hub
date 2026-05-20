/*
Copyright (c) 2026
*/
package cmd

import (
	"fmt"
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

func init() {
	rootCmd.AddCommand(serverCmd)
	serverCmd.AddCommand(serverAddCmd)
}
