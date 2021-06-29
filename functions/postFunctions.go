package functions

import (
	"demo/db"
	"fmt"
	"os"
)

//done

func InsertPostDB(post db.Post) bool {
	db := db.DBConn()
	defer db.Close()

	db_name := os.Getenv("DB_NAME")
	in := "INSERT INTO " + db_name + ".post set phone =?, name =?,productiontype =? ,amount =?, government =?, usertype =?, area =?, date =?"
	insert, err := db.Prepare(in)
	insert.Exec(post.Phone, post.Name, post.ProductType, post.Amount, post.Government, post.UserType, post.Area, post.Date)
	if err != nil {
		panic(err.Error())
	}
	return true
}

//done
func GetPostDB(typeProduction string) ([]db.Post, bool) {
	posts := []db.Post{}
	dbb := db.DBConn()
	defer dbb.Close()
	db_name := os.Getenv("DB_NAME")
	query := "SELECT id, phone, name, productiontype , amount, government, usertype , area, date FROM " + db_name + ".post WHERE productiontype = ?"
	result, err := dbb.Query(query, typeProduction)
	if err != nil {
		fmt.Println("error ", err.Error())

	}
	var id int
	var userPhone, name, productiontype, amount, government, usertype, area, date string

	for result.Next() {
		err = result.Scan(&id, &userPhone, &name, &productiontype, &amount, &government, &usertype, &area, &date)
		if err != nil {
			fmt.Println(err.Error())
		}
		post := db.Post{
			PostID:      id,
			Phone:       userPhone,
			Name:        name,
			ProductType: productiontype,
			Amount:      amount,
			Government:  government,
			UserType:    usertype,
			Area:        area,
			Date:        date,
		}
		posts = append(posts, post)
	}
	return posts, true
}

//done
func GetPostWithPhoneDB(phone string) ([]db.Post, bool) {
	posts := []db.Post{}
	dbb := db.DBConn()
	defer dbb.Close()
	db_name := os.Getenv("DB_NAME")
	fmt.Println(phone, "|db")
	query := "SELECT id ,phone, name, productiontype , amount, government, usertype , area, date FROM " + db_name + ".post WHERE productiontype = ?"
	result, err := dbb.Query(query, phone)
	if err != nil {
		fmt.Println("error ", err.Error())

	}
	var id int
	var userPhone, name, productiontype, amount, government, usertype, area, date string

	for result.Next() {
		err = result.Scan(&id, &userPhone, &name, &productiontype, &amount, &government, &usertype, &area, &date)
		if err != nil {
			fmt.Println(err.Error())
		}
		post := db.Post{
			PostID:      id,
			Phone:       userPhone,
			Name:        name,
			ProductType: productiontype,
			Amount:      amount,
			Government:  government,
			UserType:    usertype,
			Area:        area,
			Date:        date,
		}
		posts = append(posts, post)
	}
	return posts, true
}

//done
func GetPostForUSerDB(userType string) ([]db.Post, bool) {
	posts := []db.Post{}
	dbb := db.DBConn()
	defer dbb.Close()
	db_name := os.Getenv("DB_NAME")
	fmt.Println(userType, "|db")
	query := "SELECT id ,phone, name, productiontype , amount, government, usertype , area, date FROM " + db_name + ".post WHERE usertype = ?"
	result, err := dbb.Query(query, userType)
	if err != nil {
		fmt.Println("error ", err.Error())

	}
	var id int
	var userPhone, name, productiontype, amount, government, usertype, area, date string

	for result.Next() {
		err = result.Scan(&id, &userPhone, &name, &productiontype, &amount, &government, &usertype, &area, &date)
		if err != nil {
			fmt.Println(err.Error())
		}
		post := db.Post{
			PostID:      id,
			Phone:       userPhone,
			Name:        name,
			ProductType: productiontype,
			Amount:      amount,
			Government:  government,
			UserType:    usertype,
			Area:        area,
			Date:        date,
		}
		posts = append(posts, post)
	}
	return posts, true
}

//done
func GetPostsDB() ([]db.Post, bool) {
	posts := []db.Post{}
	dbb := db.DBConn()
	defer dbb.Close()
	db_name := os.Getenv("DB_NAME")
	query := "SELECT id,phone, name, productiontype , amount, government, usertype , area, date FROM " + db_name + ".post "
	result, err := dbb.Query(query)
	if err != nil {
		fmt.Println("error ", err.Error())

	}
	var id int
	var userPhone, name, productiontype, amount, government, usertype, area, date string

	for result.Next() {
		err = result.Scan(&id, &userPhone, &name, &productiontype, &amount, &government, &usertype, &area, &date)
		if err != nil {
			fmt.Println(err.Error())
		}
		post := db.Post{
			PostID:      id,
			Phone:       userPhone,
			Name:        name,
			ProductType: productiontype,
			Amount:      amount,
			Government:  government,
			UserType:    usertype,
			Area:        area,
			Date:        date,
		}
		posts = append(posts, post)
	}
	return posts, true
}

//done
func DeletePostDb(id int) bool {
	dbb := db.DBConn()
	defer dbb.Close()
	db_name := os.Getenv("DB_NAME")
	fmt.Println("from db", id)
	delete, err := dbb.Prepare("DELETE FROM " + db_name + ".post WHERE id =?")
	if err != nil {
		fmt.Println("err", err.Error())
		return false
	}
	_, err = delete.Exec(id)
	if err != nil {
		panic(err.Error())
		//return false
	}

	return true

}

//done
func UpdatePostDB(id int, post db.Post) bool {
	db := db.DBConn()
	defer db.Close()
	db_name := os.Getenv("DB_NAME")
	query := "UPDATE " + db_name + ".post set phone =?, name =?,productiontype =? ,amount =?, government =?, usertype =?, area =?, date =? WHERE id =?"
	update, err := db.Prepare(query)
	if err != nil {
		panic(err.Error())
	}
	_, err = update.Exec(post.Phone, post.Name, post.ProductType, post.Amount, post.Government, post.UserType, post.Area, post.Date, id)
	if err != nil {
		panic(err.Error())
	}
	return true
}
