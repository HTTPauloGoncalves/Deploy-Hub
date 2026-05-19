package internal

import (
	"github.com/HTTPauloGoncalves/Deploy-Hub/internal/config"
)

func NewServer(name string, server config.Server) error {
	err := config.AddServer(config.DefaultPath, name, server)
	if err != nil {
		return err
	}
	return nil
}

func UpdateServer(name string, server config.Server) error {
	err := config.UpdateServer(config.DefaultPath, name, server)
	if err != nil {
		return err
	}
	return nil
}

func GetServer(name string) (config.Server, error) {
	server, err := config.GetServer(config.DefaultPath, name)
	if err != nil {
		return config.Server{}, err
	}
	return server, nil
}

func DeleteServer(name string) error {
	err := config.RemoveServer(config.DefaultPath, name)
	if err != nil {
		return err
	}
	return nil
}
