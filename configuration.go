package main

type Configuration struct {
	ListenAddress string `json:"listenAddress"`
	ListenPort    string `json:"listenPort"`

	StaticDirectory string `json:"staticDirectory"`
}
