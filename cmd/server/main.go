package main

import (
	"fmt"

	"github.com/gabrielAnFran/api-go/configs"
)

func main() {
	fmt.Println("hei")
	// Calls the LoadConfig func passing the path to the config file
	// that is in the same directory.
	configs, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	fmt.Println(configs.DBName)
}
