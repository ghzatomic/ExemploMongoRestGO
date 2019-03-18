package main

import (
	. "ExemploMongoRest1/config"
	. "ExemploMongoRest1/repository"
	. "ExemploMongoRest1/rest/test"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

var dao = ProcessamentoDAO{}
var config = Config{}

func init() {
	config.Read()

	dao.Server = config.Server
	dao.Database = config.Database
	dao.Connect()
}

// função principal
func main() {
	router := mux.NewRouter()

	router.HandleFunc("/processamento", GetAll).Methods("GET")

	log.Fatal(http.ListenAndServe(":8000", router))
}

