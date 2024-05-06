//go:build script

package main

import (
	"gopkg.in/yaml.v3"
	"os"
)

func main() {
	_, err := os.Stat("config.yaml")
	if err != nil && os.IsNotExist(err) {
		println("Config file not found, skipping...")
		return
	}

	file, err := os.ReadFile("config.yaml")
	if err != nil {
		panic(err)
	}

	obj := make(map[string]interface{})

	err = yaml.Unmarshal(file, &obj)

	if err != nil {
		panic(err)
	}

	obj["environment"] = "development"

	file, err = yaml.Marshal(obj)

	if err != nil {
		panic(err)
	}

	err = os.WriteFile("config.yaml", file, 0644)
	println("Changed config environment to development")
}
