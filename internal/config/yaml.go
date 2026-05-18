package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

const defaultPath = "deploy.yaml"

func NewConfig(projectName string) Config {
	return Config{
		Project:  projectName,
		Servers:  make(map[string]Server),
		Services: make(map[string]Service),
	}
}

func CreateConfigFile(projectName string, path string) error {
	if path == "" {
		path = defaultPath
	}

	if FileExists(path) {
		return fmt.Errorf("Arquivo já %s existe", path)
	}

	cfg := NewConfig(projectName)

	return SaveConfig(path, cfg)
}

func SaveConfig(path string, cfg Config) error {
	if path == "" {
		path = defaultPath
	}

	data, err := yaml.Marshal(&cfg)

	if err != nil {
		return err
	}

	return os.WriteFile(path, data, 0644)
}

func LoadConfig(path string) (Config, error) {
	if path == "" {
		path = defaultPath
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
		return fmt.Errorf("Erro ao carregar as configurações")
	}

	if name == "" {
		return fmt.Errorf("O nome não pode estar vazio")
	}

	if _, exist := cfg.Servers[name]; exist {
		return fmt.Errorf("Já existe um servidor com esse nome")
	}

	cfg.Servers[name] = server

	return SaveConfig(path, cfg)
}

func AddService(path string, name string, service Service) error {
	cfg, err := LoadConfig(path)

	if err != nil {
		return fmt.Errorf("Erro ao carregar as configurações")
	}

	if name == "" {
		return fmt.Errorf("O nome não pode estar vazio")
	}

	if _, exist := cfg.Servers[name]; exist {
		return fmt.Errorf("Já existe um serviço de deploy com esse nome")
	}

	cfg.Services[name] = service

	return SaveConfig(path, cfg)
}

func FileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}
