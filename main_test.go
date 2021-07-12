package main

import (
	"bytes"
	"demo/functions"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

//done
func TestGetAllPosts(t *testing.T){
	request, _ := http.NewRequest("GET", "/getPosts", nil)
	response := httptest.NewRecorder()
	Router().ServeHTTP(response, request)
	assert.Equal(t, 200, response.Code, "OK response is expected")
}

func TestGetPostWithKind(t *testing.T){
	request, _ := http.NewRequest("GET", "/getPost/{productType}", nil)
	response := httptest.NewRecorder()
	Router().ServeHTTP(response, request)
	assert.Equal(t, 200, response.Code, "OK response is expected")
}

func TestGetPostPhone(t *testing.T){
	request, _ := http.NewRequest("GET", "/getPostsWithPhone/{phone}", nil)
	response := httptest.NewRecorder()
	Router().ServeHTTP(response, request)
	assert.Equal(t, 200, response.Code, "OK response is expected")
}

func TestGetReport(t *testing.T){
	request, _ := http.NewRequest("GET", "/report", nil)
	response := httptest.NewRecorder()
	Router().ServeHTTP(response, request)
	assert.Equal(t, 200, response.Code, "OK response is expected")
}

func TestGetPostForUser(t *testing.T){
	request, _ := http.NewRequest("GET", "/getPostsForUser/{userType}", nil)
	response := httptest.NewRecorder()
	Router().ServeHTTP(response, request)
	assert.Equal(t, 200, response.Code, "OK response is expected")
}

func TestInsertPost(t *testing.T){

	var jsonStr = []byte(`{ "name":"waleed reda",
    "phone":"01018148645",
    "productiontype":"food",
    "productionName":"chebsi",
    "amount":12,
    "unit":"carton",
    "imageUrl":"until",
    "government":"menofia",
    "usertype":"mandop",
    "area":"birket elsab"}`)
	request, _ := http.NewRequest("POST", "/add", bytes.NewBuffer(jsonStr))
	response := httptest.NewRecorder()
	Router().ServeHTTP(response, request)
	assert.Equal(t, http.StatusAccepted, response.Code, "OK response is expected")
}

func TestContactUSER(t *testing.T){

	request, _ := http.NewRequest("GET", "/contactUser/{phone}/{toCall}", nil)
	response := httptest.NewRecorder()
	Router().ServeHTTP(response, request)
	assert.Equal(t, 200, response.Code, "OK response is expected")
}

func TestSearch(t *testing.T){
	request, _ := http.NewRequest("GET", "/search/{text}", nil)
	response := httptest.NewRecorder()
	Router().ServeHTTP(response, request)
	assert.Equal(t, 200, response.Code, "OK response is expected")
}

func TestDeletePost(t *testing.T){
	var jsonStr = []byte(`{ "name":"waleed reda",
    "phone":"01018148645",
    "productiontype":"food",
    "productionName":"chebsi",
    "amount":12,
    "unit":"carton",
    "imageUrl":"until",
    "government":"menofia",
    "usertype":"mandop",
    "area":"birket elsab"}`)
	request, _ := http.NewRequest("DELETE", "/deletePost/{id}", bytes.NewBuffer(jsonStr))
	response := httptest.NewRecorder()
	Router().ServeHTTP(response, request)
	assert.Equal(t, http.StatusAccepted, response.Code, "response is expected")
}

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/getPosts", functions.GetAllPosts).Methods("GET")
	router.HandleFunc("/getPost/{productType}", functions.GetPostWithKind).Methods("GET")
	router.HandleFunc("/getPostsWithPhone/{phone}", functions.GetPostPhone).Methods("GET")
	router.HandleFunc("/report", functions.GetReport).Methods("GET")
	router.HandleFunc("/search/{text}", functions.Search).Methods("GET")
	router.HandleFunc("/contactUser/{phone}/{toCall}", functions.ContactUSER).Methods("GET")
	router.HandleFunc("/getPostsForUser/{userType}", functions.GetPostForUser).Methods("GET")
	router.HandleFunc("/add", functions.InsertPost).Methods("POST")
	router.HandleFunc("/deletePost/{id}", functions.InsertPost).Methods("DELETE")

	return router
}
