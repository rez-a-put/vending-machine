package main

import (
	"log"
	"net/http"
	"vending-machine/handler"
	"vending-machine/utils"

	"github.com/gorilla/mux"
)

var r *mux.Router

func init() {
	r = mux.NewRouter()
}

func main() {
	r.HandleFunc("/vending_machine", handler.GetItemList).Methods("GET")
	r.HandleFunc("/vending_machine/buy", handler.BuyItem).Methods("POST")
	r.HandleFunc("/vending_machine/items", handler.AddItem).Methods("POST")
	r.HandleFunc("/vending_machine/items/{id}", handler.ModifyItem).Methods("PUT")
	r.HandleFunc("/vending_machine/items/{id}", handler.RemoveItem).Methods("DELETE")

	log.Fatal(http.ListenAndServe(utils.GetEnvByKey("SERVER_HOST"), r))
}
