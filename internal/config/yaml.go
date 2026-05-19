package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

const DefaultPath = "deploy.yaml"

func NewConfig(projectName string) Config {
	return Config{
		Project:  projectName,
		Servers:  make(map[string]Server),
		Services: make(map[string]Service),
	}
}

func CreateConfigFile(projectName string, path string) error {
	if path == "" {
		path = DefaultPath
	}

	if FileExists(path) {
		return fmt.Errorf("Arquivo já %s existe", path)
	}

	cfg := NewConfig(projectName)

	return saveConfig(path, cfg)
}

func saveConfig(path string, cfg Config) error {
	if path == "" {
		path = DefaultPath
	}

	data, err := yaml.Marshal(&cfg)

	if err != nil {
		return err
	}

	return os.WriteFile(path, data, 0644)
}

func LoadConfig(path string) (Config, error) {
	if path == "" {
		path = DefaultPath
	}

	var cfg Config

	data, err := os.ReadFile(path)

	if err != nil {
		return cfg, err
	}

	err = yaml.Unmarshal(data, &cfg)

	if err != nil {
		return cfg, err
	}

	if cfg.Servers == nil {
		cfg.Servers = make(map[string]Server)
	}

	if cfg.Services == nil {
		cfg.Services = make(map[string]Service)
	}

	return cfg, nil
}

func AddServer(path string, name string, server Server) error {
	cfg, err := LoadConfig(path)

	if err != nil {
		return fmt.Errorf("Erro ao carregar as configurações. %s", err)
	}

	if name == "" {
		return fmt.Errorf("O nome não pode estar vazio")
	}

	if _, exist := cfg.Servers[name]; exist {
		return fmt.Errorf("Já existe um servidor com esse nome")
	}

	cfg.Servers[name] = server

	return saveConfig(path, cfg)
}

func AddService(path string, name string, service Service) error {
	cfg, err := LoadConfig(path)

	if err != nil {
		return fmt.Errorf("Erro ao carregar as configurações. %s", err)
	}

	if name == "" {
		return fmt.Errorf("O nome não pode estar vazio")
	}

	if _, exist := cfg.Services[name]; exist {
		return fmt.Errorf("Já existe um serviço de deploy com esse nome")
	}

	cfg.Services[name] = service

	return saveConfig(path, cfg)
}

func GetServer(path string, name string) (Server, error) {
	cfg, err := LoadConfig(path)

	if err != nil {
		return Server{}, err
	}

	server, exist := cfg.Servers[name]

	if !exist {
		return Server{}, fmt.Errorf("O servidor %s não foi encontrado", name)
	}

	return server, nil
}

func GetService(path string, name string) (Service, error) {
	cfg, err := LoadConfig(path)

	if err != nil {
		return Service{}, fmt.Errorf("Erro ao carregar as configurações. %s", err)
	}

	service, exist := cfg.Services[name]

	if !exist {
		return Service{}, fmt.Errorf("O deploy %s não foi encontrado", name)
	}

	return service, nil
}

func ListServers(path string) (map[string]Server, error) {
	cfg, err := LoadConfig(path)

	if err != nil {
		return nil, err
	}

	return cfg.Servers, nil
}

func ListServices(path string) (map[string]Service, error) {
	cfg, err := LoadConfig(path)

	if err != nil {
		return nil, err
	}

	return cfg.Services, nil
}

func RemoveService(path string, name string) error {
	cfg, err := LoadConfig(path)

	if err != nil {
		return err
	}

	delete(cfg.Services, name)

	saveConfig(path, cfg)

	return nil
}

func RemoveServer(path string, name string) error {
	cfg, err := LoadConfig(path)

	if err != nil {
		return err
	}

	delete(cfg.Servers, name)

	saveConfig(path, cfg)

	return nil
}

func UpdateService(path string, name string, service Service) error {
	cfg, err := LoadConfig(path)

	if err != nil {
		return err
	}

	if _, exist := cfg.Services[name]; !exist {
		return fmt.Errorf("O serviço de deploy %s não existe", name)
	}

	cfg.Services[name] = service

	saveConfig(path, cfg)

	return nil
}

func UpdateServer(path string, name string, service Server) error {
	cfg, err := LoadConfig(path)

	if err != nil {
		return err
	}

	if _, exist := cfg.Servers[name]; !exist {
		return fmt.Errorf("O servidor %s não existe", name)
	}

	cfg.Servers[name] = service

	saveConfig(path, cfg)

	return nil
}

func FileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}
