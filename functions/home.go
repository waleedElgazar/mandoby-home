package functions

import (
	"demo/db"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
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
			Unit: 		 creds.Unit,
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
func Search(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	text := param["text"]
	json.NewEncoder(w).Encode(GetPostsWithSearch(text))
}

func GetReport(w http.ResponseWriter, r *http.Request ){
	json.NewEncoder(w).Encode(CreateReport())
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

func ContactUSER(w http.ResponseWriter, r *http.Request) {
	parm := mux.Vars(r)
	phone := parm["phone"]
	tocall := parm["toCall"]
	validToken, err := GenerateJWT(phone)
	if err != nil {
		fmt.Println("failed to generate token ", err.Error())
	}
	client := &http.Client{}

	url := "https://gp-mandoob-users.herokuapp.com/isAuthorized/" + tocall

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("error ", err.Error())
	}
	req.Header.Set("Token", validToken)

	res, err := client.Do(req)
	if err != nil {
		fmt.Println("error ", err.Error())
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("error ", err)
	}
	var user db.User
	err = json.Unmarshal(body, &user)
	if err != nil {
		fmt.Println("error", err.Error())
	}
	json.NewEncoder(w).Encode(user)

}

var mySigningKey = []byte("mandopy-project")

func GenerateJWT(phone string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["authorized"] = true
	claims["phone"] = phone
	claims["exp"] = time.Now().Add(time.Minute * 5).Unix()
	tokenString, err := token.SignedString(mySigningKey)

	if err != nil {
		fmt.Println("Something Went Wrong: ", err.Error())
		return "", err
	}

	return tokenString, nil
}

func Welcome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "welcome")
}
