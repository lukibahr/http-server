// package app implements the main app logic

package main

import (
	"io/ioutil"
	"os"

	log "github.com/sirupsen/logrus"

	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type App struct {
	Router *mux.Router
}

func (a *App) Initialize() {
	a.Router = mux.NewRouter()
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.InfoLevel)

	a.initializeRoutes()
}

func (a *App) Run(addr string) {
	log.Fatal(http.ListenAndServe(addr, a.Router))
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func (a *App) getHandler(w http.ResponseWriter, r *http.Request) {
	log.WithFields(log.Fields{
		"remotehost": r.Host,
		"addr":       r.RemoteAddr,
		"body":       r.Body,
	}).Info("get request received")
	respondWithJSON(w, http.StatusOK, "-XGET")
}

func (a *App) getHeaderHandler(w http.ResponseWriter, r *http.Request) {
	log.WithFields(log.Fields{
		"remotehost": r.Host,
		"addr":       r.RemoteAddr,
		"body":       r.Body,
		"headers":    r.Header,
	}).Info("get request received")
	respondWithJSON(w, http.StatusOK, "-XHEADERS")
}

func (a *App) postHandler(w http.ResponseWriter, r *http.Request) {
	// Read body
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	log.WithFields(log.Fields{
		"remotehost": r.Host,
		"addr":       r.RemoteAddr,
		"body":       string(b),
		"headers":    r.Header,
	}).Info("post request received")
	respondWithJSON(w, http.StatusOK, "-XPOST")
}

func (a *App) initializeRoutes() {
	a.Router.HandleFunc("/get", a.getHandler).Methods("GET")
	a.Router.HandleFunc("/headers", a.getHeaderHandler).Methods("GET")
	a.Router.HandleFunc("/post", a.postHandler).Methods("POST")
}
