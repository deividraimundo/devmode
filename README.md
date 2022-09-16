# DevMode

## Configurando ambiente de desenvolvimento

### 1º Passo: Criar uma pasta que colocará todo seu projeto.

### 2º Passo: Iniciar o air com o comando air init

### 3º Passo: Criar uma pasta cmd e dentro dela cdm/nomeDaPasta

### 4º Passo: Criar um arquivo main.go e colocar o seguinte código dentro dele:
```
package main

import (
	"fmt"

	"github.com/deividraimundo/devmode/config"
)

var devMode = "false"

func main() {
	cfg := config.New(devMode)
	fmt.Println(cfg)
}
```
### 5º Passo: Criar uma pasta config e dentro dela um arquivo config.go, com o seguinte código:
```
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
```
### 6º Passo: Configurar o cmd do seu arquivo .air.toml da seguinte maneira:

`cmd = "go build -o ./tmp/main -ldflags=\"-X 'main.devMode=true'\" ./cmd/devmode"`
