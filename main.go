package main

import (
	"fmt"
	"net/http"
	"sync"

	log "github.com/Sirupsen/logrus"
	"github.com/bluehawk27/assignment/config"
	"github.com/bluehawk27/assignment/httpapi"
	"github.com/gorilla/mux"
)

func main() {

	addr := config.GetProxyConnectionString()

	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		r := mux.NewRouter()
		log.Info("Running HTTP Server on: " + addr)
		r.HandleFunc("/ping", httpapi.Ping).Methods("GET")
		r.HandleFunc("/add/{arg}", httpapi.Add).Methods("POST")
		r.HandleFunc("/get/{arg}", httpapi.Get).Methods("GET")

		http.Handle("/", r)
		if err := http.ListenAndServe(addr, nil); err != nil {
			log.Fatal(fmt.Sprintf("Fatal error server.Serve: %s \n", err))
		}
	}()
	wg.Wait()
}
