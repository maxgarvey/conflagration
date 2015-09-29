package main

type Configuration struct {
	ListenAddress string `json:"listenAddress"`
	ListenPort    string `json:"listenPort"`

	// this directory is where all the static files for the
	// UI are located
	StaticDirectory string `json:"staticDirectory"`

	// here's the information about the PG database that we're connecting to
	DBHost string `json:"dbHost"`
	DBPort string `json:"dbPort"`
	DBUser string `json:"dbUser"`
	DBPass string `json:"dbPass"`
	DBName string `json:"dbName"`
}
