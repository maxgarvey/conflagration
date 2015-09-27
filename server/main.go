package server

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

type Instance struct {
	ListenAddress string
	ListenPort    string

	Router *mux.Router
}

func New() Instance {
	fmt.Printf("in server New()\n")

	var instance Instance
	return instance
}

func (instance *Instance) Initialize() {
	router := mux.NewRouter()

	router.HandleFunc("/status/", StatusHandler)
	router.HandleFunc("/api/", ApiHandler)
	router.Handle("/static/{staticPath}",
		http.StripPrefix(
			"/static/",
			http.FileServer(
				http.Dir(
					"static"))))

	instance.Router = router
}

func (instance *Instance) Run() {
	fmt.Printf("in server.Run().\n")
	http.ListenAndServe(":8080", instance.Router)
}

func StatusHandler(respW http.ResponseWriter, req *http.Request) {
	fmt.Printf("in status handler.\n")
}

func ApiHandler(respW http.ResponseWriter, req *http.Request) {
	fmt.Printf("in api handler.\n")
}

// func StaticHandler() http.HandlerFunc {
// 	fmt.Printf("in static handler.\n")
// 	return http.FileServer(http.Dir("/Users/maxgarvey/go/src/github.com/maxgarvey/conflagration/static")).ServeHTTP
// }
