//go:build script

package main

import (
	"gopkg.in/yaml.v3"
	"os"
)

func main() {
	file, err := os.ReadFile("config.yaml")
	if err != nil {
		panic(err)
	}

	obj := make(map[string]interface{})

	err = yaml.Unmarshal(file, &obj)

	if err != nil {
		panic(err)
	}

	obj["environment"] = "production"

	file, err = yaml.Marshal(obj)

	if err != nil {
		panic(err)
	}

	err = os.WriteFile("config.yaml", file, 0644)
}
