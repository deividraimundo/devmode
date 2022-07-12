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
