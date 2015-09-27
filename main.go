package main

import (
	"fmt"
	"github.com/maxgarvey/conflagration/server"
)

var (
	configuration  Configuration
	serverInstance server.Instance
)

func main() {
	//fmt.Printf("in func main\n") // debug
	config()

	serverInstance.Run()
}

func config() {
	//fmt.Printf("in func config\n") // debug

	// TODO: importation from config file
	configuration = Configuration{
		"0.0.0.0",
		"8000",
		"./static",
	}

	serverInstance = server.New()
	serverInstance.Initialize()
}
