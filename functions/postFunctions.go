package functions

import (
	"demo/db"
	"fmt"
	"time"
)

func InsertPostDB(post db.Post) bool {
	db := db.DBConn()
	defer db.Close()
	insert, err := db.Prepare("INSERT INTO login.post VALUES(?,?,?,?,?,?)")
	insert.Exec(post.Phone, post.Name, post.Type, post.Amount, post.Government, post.Date)
	if err != nil {
		panic(err.Error())
	}
	return true
}

func GetPostDB(kindd string) ([]db.Post, bool) {
	posts := []db.Post{}
	dbb := db.DBConn()
	defer dbb.Close()
	get, err := dbb.Query("SELECT phone, name, type, amount, government, date FROM login.post WHERE type = ?", kindd)
	if err != nil {
		return nil, false
	}

	var phon, name, amount, government, kind string
	var date time.Time
	for get.Next() {
		err = get.Scan(&phon, &name, &kind, &amount, &government, &date)
		if err != nil {
			fmt.Println(err.Error())
		}
		post := db.Post{
			Phone:      phon,
			Name:       name,
			Type:       kind,
			Amount:     amount,
			Government: government,
			Date:       date,
		}
		//fmt.Println(phon, name, kind, amount, government, date)
		posts = append(posts, post)
	}
	return posts, true

}

func GetPostsDB() ([]db.Post, bool) {
	posts := []db.Post{}
	dbb := db.DBConn()
	defer dbb.Close()
	get, err := dbb.Query("SELECT phone, name, type, amount, government, date FROM login.post")
	if err != nil {
		return nil, false
	}

	var phon, name, amount, government, kind string
	var date time.Time
	for get.Next() {
		err = get.Scan(&phon, &name, &kind, &amount, &government, &date)
		if err != nil {
			fmt.Println(err.Error())
		}
		post := db.Post{
			Phone:      phon,
			Name:       name,
			Type:       kind,
			Amount:     amount,
			Government: government,
			Date:       date,
		}
		//fmt.Println(phon, name, kind, amount, government, date)
		posts = append(posts, post)
	}
	return posts, true

}
