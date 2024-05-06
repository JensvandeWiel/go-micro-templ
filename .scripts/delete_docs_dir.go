//go:build script

package main

import "os"

func main() {
	_, err := os.Stat("docs")
	if err != nil && os.IsNotExist(err) {
		println("Docs directory not found, skipping...")
		return
	}

	err = os.RemoveAll("docs")
	if err != nil {
		panic(err)
	}

	println("Deleted docs directory")
}
