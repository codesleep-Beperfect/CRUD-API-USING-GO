package routes

import (
	"github.com/codesleep-Beperfect/Bookstore/controller"
	"github.com/gorilla/mux"
)
var UserDetail = func(router *mux.Router){
	router.HandleFunc("/user", controller.CreateUser).Methods("POST")
	router.HandleFunc("/user", controller.GetUser).Methods("GET")
	router.HandleFunc("/user/{id}", controller.GetUserById).Methods("GET")
	router.HandleFunc("/user/{id}", controller.UpdateUser).Methods("PUT")
	router.HandleFunc("/user/{id}", controller.DeleteUser).Methods("DELETE")
	// router.HandleFunc("/user/{id}", controller.deleteUser).Methods("DELETE")
}