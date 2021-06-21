package main

import (
	"demo/functions"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	port := os.Getenv("PORT")
	
	router := mux.NewRouter()
	router.HandleFunc("/add", functions.InsertPost).Methods("POST")
	router.HandleFunc("/posts/{kind}", functions.GetPost).Methods("GET")
	router.HandleFunc("/posts", functions.GetPosts).Methods("GET")
	http.ListenAndServe(":"+port, router)
}
