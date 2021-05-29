package main

import (
	"demo/functions"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/add", functions.InsertPost).Methods("POST")
	router.HandleFunc("/posts/{kind}", functions.GetPost).Methods("GET")
	router.HandleFunc("/posts", functions.GetPosts).Methods("GET")
	http.ListenAndServe(":8088", router)
}
