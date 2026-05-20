package internal

import (
	"fmt"

	"github.com/HTTPauloGoncalves/Deploy-Hub/internal/config"
)

func Deploy(name string) error {
	cfg, err := config.LoadConfig(config.DefaultPath)
	if err != nil {
		return err
	}

	service, exists := cfg.Services[name]
	if !exists {
		return fmt.Errorf("serviço %s não encontrado", name)
	}

	server, exists := cfg.Servers[service.Server]
	if !exists {
		return fmt.Errorf("servidor %s não encontrado", service.Server)
	}

	client, err := ConnectWithPassword(server.User, server.Host, server.Port, server.Password)
	if err != nil {
		return err
	}
	defer client.Close()

	for _, command := range service.Commands {
		fullCommand := fmt.Sprintf("cd %s && %s", service.Path, command)

		_, out, err := client.Run(fullCommand)
		if err != nil {
			return fmt.Errorf("erro ao executar comando %q: %w", command, err)
		}

		fmt.Printf("Saída do comando %q:\n%s\n", command, out)
	}

	return nil
}
