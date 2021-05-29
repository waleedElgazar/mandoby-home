package functions

import (
	"demo/db"
	"encoding/json"
	"net/http"
	"time"
)

func InsertPost(w http.ResponseWriter, r *http.Request) {
	var creds db.Post
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		w.Header().Set("Content-Type", "text/plain")
		w.Write([]byte("failed"))
		w.Write([]byte(err.Error()))
		w.WriteHeader(http.StatusBadRequest)
		return
	} else {
		post := db.Post{
			Name:       creds.Name,
			Phone:      creds.Phone,
			Type:       creds.Type,
			Amount:     creds.Amount,
			Date:       time.Now().UTC(),
			Government: creds.Government,
		}
		InsertPostDB(post)
		w.Header().Set("Content-Type", "text/plain")
		w.Write([]byte("added succ"))
		w.WriteHeader(http.StatusAccepted)
		return
	}
}

func GetPost(w http.ResponseWriter, r *http.Request) {
	var post db.Post
	err := json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		w.Header().Set("Content-Type", "text/plain")
		w.Write([]byte("there is error happened\n try again"))
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	kind := post.Type
	posts, founded := GetPostDB(kind)
	if founded {
		w.Header().Set("Content-Type", "text/plain")
		w.Write([]byte("founded posts like that "))
		json.NewEncoder(w).Encode(&posts)
		w.WriteHeader(http.StatusAccepted)
		return
	} else {
		json.NewEncoder(w).Encode(&posts)
		w.Header().Set("Content-Type", "text/plain")
		w.Write([]byte("there is no posts with that type"))
		w.WriteHeader(http.StatusAccepted)
		return
	}
}

func GetPosts(w http.ResponseWriter, r *http.Request) {
	var post db.Post
	err := json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		w.Header().Set("Content-Type", "text/plain")
		w.Write([]byte("there is error happened\n try again"))
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	posts, _ := GetPostsDB()
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("founded posts like that "))
	json.NewEncoder(w).Encode(&posts)

}
