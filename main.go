package main

import (
	//"fmt"
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
		// database credentials
		"127.0.0.1",
		"4567",
		"dbuser",
		"dbpass",
		"dbname",
	}

	serverInstance = server.New()
	serverInstance.Initialize()
}
