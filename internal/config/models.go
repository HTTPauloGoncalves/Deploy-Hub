package config

type Config struct {
	Project  string             `yaml:"project"`
	Servers  map[string]Server  `yaml:"servers"`
	Services map[string]Service `yaml:"services"`
}

type Server struct {
	Host string `yaml:"host"`
	User string `yaml:"user"`
	Auth string `yaml:"auth"`
}

type Service struct {
	Server   string   `yaml:"server"`
	Path     string   `yaml:"path"`
	Commands []string `yaml:"command"`
}
