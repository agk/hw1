package utils

import (
	"hw1/handlers"
	"github.com/gorilla/mux"
)

func BuildItemResource(router *mux.Router, prefix string) {
	router.HandleFunc(prefix+"/{id}", handlers.GetItemById).Methods("GET")
	router.HandleFunc(prefix, handlers.CreateItem).Methods("POST")
	router.HandleFunc(prefix+"/{id}", handlers.UpdateItemById).Methods("PUT")
	router.HandleFunc(prefix+"/{id}", handlers.DeleteItemById).Methods("DELETE")
}

func BuildManyItemsResourcePrefix(router *mux.Router, prefix string) {
	router.HandleFunc(prefix, handlers.GetAllItems).Methods("GET")
}
