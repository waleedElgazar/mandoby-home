package functions

import (
	"demo/db"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

//done
func InsertPost(w http.ResponseWriter, r *http.Request) {
	var creds db.Post
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	} else {
		currentTime := time.Now()
		post := db.Post{
			Phone:       creds.Phone,
			Name:        creds.Name,
			ProductType: creds.ProductType,
			ProductName: creds.ProductName,
			Amount:      creds.Amount,
			ImageUrl:    creds.ImageUrl,
			Government:  creds.Government,
			UserType:    creds.UserType,
			Area:        creds.Area,
			Date:        currentTime.Format("2006.01.02 15:04:05"),
		}
		InsertPostDB(post)
		w.WriteHeader(http.StatusAccepted)
		return
	}
}

//done
func GetPostWithKind(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	posts, founded := GetPostDB(params["productType"])
	if founded {
		json.NewEncoder(w).Encode(&posts)
		w.WriteHeader(http.StatusAccepted)
		return
	} else {
		json.NewEncoder(w).Encode(&posts)
		w.WriteHeader(http.StatusAccepted)
		return
	}
}

//done
func GetPostPhone(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	phone := params["phone"]
	posts, founded := GetPostWithPhoneDB(phone)
	if founded {
		json.NewEncoder(w).Encode(&posts)
		w.WriteHeader(http.StatusAccepted)
		return
	} else {
		json.NewEncoder(w).Encode(&posts)
		w.WriteHeader(http.StatusAccepted)
		return
	}
}

//done
func GetPostForUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	userType := params["userType"]
	posts, founded := GetPostForUSerDB(userType)
	if founded {
		json.NewEncoder(w).Encode(&posts)
		w.WriteHeader(http.StatusAccepted)
		return
	} else {
		json.NewEncoder(w).Encode(&posts)
		w.WriteHeader(http.StatusAccepted)
		return
	}
}

//done
func GetAllPosts(w http.ResponseWriter, r *http.Request) {
	posts, founded := GetPostsDB()
	if founded {
		json.NewEncoder(w).Encode(&posts)
		w.WriteHeader(http.StatusAccepted)
		return
	} else {
		json.NewEncoder(w).Encode(&posts)
		w.WriteHeader(http.StatusAccepted)
		return
	}
}

//done
func UpdatePost(w http.ResponseWriter, r *http.Request) {
	var creds db.Post
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	} else {
		currentTime := time.Now()
		post := db.Post{
			Phone:       creds.Phone,
			Name:        creds.Name,
			ProductType: creds.ProductType,
			ProductName: creds.ProductName,
			Amount:      creds.Amount,
			ImageUrl:    creds.ImageUrl,
			Government:  creds.Government,
			UserType:    creds.UserType,
			Area:        creds.Area,
			Date:        currentTime.Format("2020.01.02 15:04:05"),
		}
		UpdatePostDB(creds.PostID, post)
		w.WriteHeader(http.StatusAccepted)
		return
	}
}

//done
func DeletePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id := params["id"]
	deleted := DeletePostDb(id)
	if deleted {
		w.WriteHeader(http.StatusAccepted)
		return
	}
}

func Welcome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "welcome")
}
