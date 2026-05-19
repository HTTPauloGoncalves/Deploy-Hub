package internal

import "github.com/HTTPauloGoncalves/Deploy-Hub/internal/config"

const constPath = "deploy.yaml"

func Init(name string) {
	if config.FileExists(constPath) {

		cfg, err := config.LoadConfig(constPath)
		if err != nil {
			panic(err)
		}

		_ = cfg

		return
	}

	err := config.CreateConfigFile(name, constPath)

	if err != nil {
		panic(err)
	}
}
