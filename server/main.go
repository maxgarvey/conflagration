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

	apiRouter := router.PathPrefix("/api").Subrouter()
	apiRouter.PathPrefix("/configurations").HandlerFunc(
		ConfigurationsHandler)
	apiRouter.PathPrefix("/configuration/{configurationId}").HandlerFunc(
		GetPutConfigurationHandler)
	apiRouter.PathPrefix("/configuration").HandlerFunc(
		NewConfigurationHandler)

	router.PathPrefix("/status").HandlerFunc(StatusHandler)
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
	fmt.Printf("in server.Run().\n") // debug
	http.ListenAndServe(
		// TODO: replace the "8080" with instance.ListenPort
		fmt.Sprintf(":%d", 8080),
		instance.Router,
	)
}

func StatusHandler(respW http.ResponseWriter, req *http.Request) {
	fmt.Printf("in status handler.\n") // debug

	respW.WriteHeader(http.StatusOK)
	io.WriteString(respW, "status OK\n")
}

func NewApiRouter() http.Handler {
	fmt.Printf("in NewApiRouter().")
	router := mux.NewRouter()

	router.HandleFunc("/api/configurations", ConfigurationsHandler)

	return router
}

func ConfigurationsHandler(respW http.ResponseWriter, req *http.Request) {
	fmt.Printf("in configurations handler.\n") // debug

	// respW.WriteHeader(http.StatusOK)
	// io.WriteString(respW, "api response here\n")

	respW.WriteHeader(http.StatusOK)                            // debug
	io.WriteString(respW, "configurations api response here\n") // debug
}

func NewConfigurationHandler(respW http.ResponseWriter, req *http.Request) {
	fmt.Printf("in new configuration handler.\n")

	respW.WriteHeader(http.StatusOK)                                // debug
	io.WriteString(respW, "new configuration api response here.\n") // debug
}

func GetPutConfigurationHandler(respW http.ResponseWriter, req *http.Request) {
	fmt.Printf("in GET/PUT configuration handler.\n")

	respW.WriteHeader(http.StatusOK)                               // debug
	io.WriteString(respW, "get put configuration handler here.\n") // debug
}
