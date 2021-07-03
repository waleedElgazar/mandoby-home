package main

import (
	"demo/functions"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {

	os.Setenv("DB_HOST", "us-cdbr-east-04.cleardb.com")
	os.Setenv("DB_PORT", "(us-cdbr-east-04.cleardb.com)")
	os.Setenv("DB_DRIVER", "mysql")
	os.Setenv("DB_ROOT", "bb8d3f8c173589")
	os.Setenv("DB_PASSWORD", "b9a00cff")
	os.Setenv("DB_NAME", "heroku_09b7638a416ab31")
	//	os.Setenv("PORT", "8081")
	router := mux.NewRouter()
	port := os.Getenv("PORT")
	router.HandleFunc("/", functions.Welcome)
	router.HandleFunc("/add", functions.InsertPost).Methods("POST")
	router.HandleFunc("/getPost/{productType}", functions.GetPostWithKind).Methods("GET")
	router.HandleFunc("/getPosts", functions.GetAllPosts).Methods("GET")
	router.HandleFunc("/getPostsForUser/{userType}", functions.GetPostForUser).Methods("GET")
	router.HandleFunc("/getPostsWithPhone/{phone}", functions.GetPostPhone).Methods("GET")
	router.HandleFunc("/updatePost", functions.UpdatePost).Methods("PUT")
	router.HandleFunc("/deletePost/{id}", functions.DeletePost).Methods("DELETE")
	http.ListenAndServe(":"+port, router)
}
