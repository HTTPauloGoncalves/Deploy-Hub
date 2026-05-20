package config

type Config struct {
	Project  string             `yaml:"project"`
	Servers  map[string]Server  `yaml:"servers"`
	Services map[string]Service `yaml:"services"`
}

type Server struct {
	Host     string `yaml:"host"`
	User     string `yaml:"user"`
	Password string `yaml:"password,omitempty"`
	Auth     string `yaml:"auth"`
	Port     int    `yaml:"port,omitempty"`
}

type Service struct {
	Server   string   `yaml:"server"`
	Path     string   `yaml:"path"`
	Commands []string `yaml:"command"`
}
