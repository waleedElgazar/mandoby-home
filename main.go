package main

import (
	"demo/functions"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	port := functions.GetPort()
	router.HandleFunc("/", functions.Welcome)
	router.HandleFunc("/add", functions.InsertPost).Methods("POST")
	router.HandleFunc("/getPost", functions.GetPostWithKind).Methods("GET")
	router.HandleFunc("/getPosts", functions.GetAllPosts).Methods("GET")
	router.HandleFunc("/getPostsForUser", functions.GetPostForUser).Methods("GET")
	router.HandleFunc("/getPostsWithPhone", functions.GetPostPhone).Methods("GET")
	router.HandleFunc("/updatePost", functions.UpdatePost).Methods("PUT")
	router.HandleFunc("/deletePost", functions.DeletePost).Methods("DELETE")
	http.ListenAndServe(":"+port, router)
}
