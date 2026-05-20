package internal

import (
	"github.com/HTTPauloGoncalves/Deploy-Hub/internal/config"
)

func NewService(name string, service config.Service) error {
	err := config.AddService(config.DefaultPath, name, service)
	if err != nil {
		return err
	}
	return nil
}

func UpdateService(name string, service config.Service) error {
	err := config.UpdateService(config.DefaultPath, name, service)
	if err != nil {
		return err
	}
	return nil
}

func GetService(name string) (config.Service, error) {
	service, err := config.GetService(config.DefaultPath, name)
	if err != nil {
		return config.Service{}, err
	}
	return service, nil
}

func ListServices() (map[string]config.Service, error) {
	services, err := config.ListServices(config.DefaultPath)
	if err != nil {
		return nil, err
	}
	return services, nil
}

func DeleteService(name string) error {
	err := config.RemoveService(config.DefaultPath, name)
	if err != nil {
		return err
	}
	return nil
}
