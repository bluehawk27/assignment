package httpapi

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"

	log "github.com/Sirupsen/logrus"
	"github.com/bluehawk27/assignment/service"
	"github.com/gorilla/mux"
)

var serv = service.NewService()

// Ping : Ping http Handler
func Ping(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	res, err := serv.Ping(ctx)
	if err != nil {
		log.Error("Error Redis not running: ", err)
	}
	io.WriteString(w, res)
}

// Add : Add http Handler
func Add(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	vars := mux.Vars(req)
	key := vars["arg"]
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Error("error getting requestBody: ", err)
	}

	serr := serv.Add(ctx, key, body)
	if serr != nil {
		log.Error("Error getting from the Service:", err)
	}

	io.WriteString(w, "Successfully added")
}

// Get : Get http Handler
func Get(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	vars := mux.Vars(req)
	arg := vars["arg"]

	res, err := serv.Get(ctx, arg)
	if err != nil {
		log.Error("Error connecting with the Service: ", err)
	}

	respBytes, jerr := json.Marshal(res)
	if jerr != nil {
		log.Error("Error Marshaling Response", jerr)
	}

	jsonString := string(respBytes)

	io.WriteString(w, jsonString)
}
