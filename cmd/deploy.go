/*
Copyright (c) 2026
*/
package cmd

import (
	"fmt"
	"sort"

	"github.com/AlecAivazis/survey/v2"
	"github.com/HTTPauloGoncalves/Deploy-Hub/internal"
	"github.com/spf13/cobra"
)

var deployCmd = &cobra.Command{
	Use:   "deploy [servico]",
	Short: "Executa o deploy de um servico",
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		serviceName := ""
		var confirm bool

		if len(args) > 0 {
			serviceName = args[0]
		} else {
			services, err := internal.ListServices()
			if err != nil {
				fmt.Println("Erro:", err)
				return
			}

			if len(services) == 0 {
				fmt.Println("Nenhum servico cadastrado. Cadastre um servico antes de fazer deploy.")
				return
			}

			serviceOptions := make([]string, 0, len(services))
			for name := range services {
				serviceOptions = append(serviceOptions, name)
			}
			sort.Strings(serviceOptions)

			err = survey.AskOne(&survey.Select{
				Message: "Servico:",
				Options: serviceOptions,
			}, &serviceName)
			if err != nil {
				fmt.Println("Erro:", err)
				return
			}
		}

		fmt.Println("Iniciando deploy:", serviceName)

		err := internal.Deploy(serviceName)
		if err != nil {
			fmt.Println("Erro:", err)
			return
		}
		survey.AskOne(&survey.Confirm{
			Message: "O deploy foi executado com sucesso. Deseja marcar o deploy como finalizado?",
			Default: true,
		}, &confirm)
		if err != nil {
			fmt.Println("Erro:", err)
			return
		}

		if !confirm {
			fmt.Println("O deploy não foi marcado como finalizado.")
			return
		}

		fmt.Println("Deploy finalizado com sucesso:", serviceName)
	},
}

func init() {
	rootCmd.AddCommand(deployCmd)
}
