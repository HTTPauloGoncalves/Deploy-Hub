package internal

import "github.com/HTTPauloGoncalves/Deploy-Hub/internal/config"

func Init(name string) {
	if config.FileExists(config.DefaultPath) {

		cfg, err := config.LoadConfig(config.DefaultPath)
		if err != nil {
			panic(err)
		}

		_ = cfg

		return
	}

	err := config.CreateConfigFile(name, config.DefaultPath)

	if err != nil {
		panic(err)
	}
}
