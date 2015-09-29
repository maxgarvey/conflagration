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
	fmt.Printf("in server New().\n") // debug

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

	if req.Method != "GET" {
		respW.WriteHeader(http.StatusNotImplemented)
		io.WriteString(respW, "requests to this endpoint must use GET Method.\n")
		return
	}

	respW.WriteHeader(http.StatusOK)                            // debug
	io.WriteString(respW, "configurations api response here\n") // debug
}

func NewConfigurationHandler(respW http.ResponseWriter, req *http.Request) {
	fmt.Printf("in new configuration handler.\n") // debug

	if req.Method != "POST" {
		respW.WriteHeader(http.StatusNotImplemented)
		io.WriteString(respW, "requests to this endpoint must use POST Method.\n")
		return
	}

	respW.WriteHeader(http.StatusOK)                                // debug
	io.WriteString(respW, "new configuration api response here.\n") // debug
}

func GetPutConfigurationHandler(respW http.ResponseWriter, req *http.Request) {
	fmt.Printf("in GET/PUT configuration handler.\n") // debug

	if req.Method != "GET" || req.Method != "PUT" {
		respW.WriteHeader(http.StatusNotImplemented)
		io.WriteString(
			respW, "requests to this endpoint must use GET or PUT Methods.\n")
		return
	}

	if req.Method == "PUT" {
		respW.WriteHeader(http.StatusAccepted)                     // debug
		io.WriteString(respW, "put configuration handler here.\n") // debug
		return
	}

	respW.WriteHeader(http.StatusOK)                           // debug
	io.WriteString(respW, "put configuration handler here.\n") // debug
}
