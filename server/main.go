package server

import (
	"fmt"
	"io"
	"net/http"

	"github.com/gorilla/mux"
)

type Instance struct {
	ListenAddress string
	ListenPort    string

	StaticFilesDirectory string

	Router *mux.Router
}

func New() Instance {
	fmt.Printf("in server New()\n")

	var instance Instance
	return instance
}

func (instance *Instance) Initialize() {
	router := mux.NewRouter()

	router.HandleFunc("/status", StatusHandler)
	router.HandleFunc("/api", ApiHandler)
	router.PathPrefix("/static").Handler(
		http.StripPrefix("/static",
			http.FileServer(
				http.Dir(
					// TODO: replace "static" with instance.StaticFilesDirectory
					"./static")),
		),
	)

	instance.Router = router
}

func (instance *Instance) Run() {
	fmt.Printf("in server.Run().\n")
	http.ListenAndServe(
		// TODO: replace the "8080" with instance.ListenPort
		fmt.Sprintf(":%d", 8080),
		instance.Router,
	)
}

func StatusHandler(respW http.ResponseWriter, req *http.Request) {
	fmt.Printf("in status handler.\n")

	respW.WriteHeader(http.StatusOK)
	io.WriteString(respW, "status OK\n")
}

func ApiHandler(respW http.ResponseWriter, req *http.Request) {
	fmt.Printf("in api handler.\n")

	respW.WriteHeader(http.StatusOK)
	io.WriteString(respW, "api response here\n")
}
