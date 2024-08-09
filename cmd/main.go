package main

import (
	"log"
	"net/http"
	"github.com/codesleep-Beperfect/Bookstore/routes"
	"github.com/gorilla/mux"
)

func main(){
	r := mux.NewRouter()
	routes.UserDetail(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:9010", r))
}