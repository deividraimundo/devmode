package config

import (
	"flag"
	"fmt"
)

type Config struct {
	DevMode bool
}

func New(devModeGoBuild string) *Config {
	devMode := false

	if devModeGoBuild == "true" {
		devMode = true
	} else {
		flag.BoolVar(&devMode, "devmode", false, "Adicionar esta flag para modo desenvolvimento.")
		flag.Parse()
	}

	if devMode {
		fmt.Println("#### Ambiente de desenvolvimento ####")
	}

	return &Config{
		DevMode: devMode,
	}
}
