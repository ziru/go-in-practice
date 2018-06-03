package main

import (
	"fmt"

	"github.com/kylelemons/go-gypsy/yaml"
)

func main() {
	config, err := yaml.ReadFile("conf.yaml")
	if err != nil {
		panic(err)
	}

	path, _ := config.Get("path")
	fmt.Println("path:", path)
	enabled, _ := config.GetBool("enabled")
	fmt.Println("enabled:", enabled)
}
