package main

import(
	"log"
	"github.com/irvandandung/gokriyatest/config"
	"github.com/gorilla/mux"
	"net/http"
)

func RunApp(){
	log.Println("init configuration...")
	pg := config.DbConn()
	r := mux.NewRouter()
	config.Routes(pg, r)
	port:= ":1234"
	log.Println("Starting server at "+port)
	log.Fatal(http.ListenAndServe(port, r))
}