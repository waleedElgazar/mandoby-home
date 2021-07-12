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
	fmt.Println("11")
	db_name := os.Getenv("DB_NAME")
	in := "INSERT INTO " + db_name + ".post set phone =?, name =?, productiontype =?, productionName=? ,amount =?,unit=? ,imageUrl=?, government =?, usertype =?, area =?, date =?"
	insert, err := db.Prepare(in)
	insert.Exec(post.Phone, post.Name, post.ProductType, post.ProductName, post.Amount, post.Unit, post.ImageUrl, post.Government, post.UserType, post.Area, post.Date)
	if err != nil {
		fmt.Println("12")
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
	query := "SELECT id, phone, name, productiontype, productionName , amount,unit ,imageUrl, government, usertype , area, date FROM " + db_name + ".post WHERE productiontype = ?"
	result, err := dbb.Query(query, typeProduction)
	if err != nil {
		fmt.Println("error ", err.Error())

	}
	var id, amount int
	var userPhone, name, productiontype, productionName, unit, url, government, usertype, area, date string

	for result.Next() {
		err = result.Scan(&id, &userPhone, &name, &productiontype, &productionName, &amount, &unit, &url, &government, &usertype, &area, &date)
		if err != nil {
			fmt.Println(err.Error())
		}
		post := db.Post{
			PostID:      id,
			Phone:       userPhone,
			Name:        name,
			ProductType: productiontype,
			Amount:      amount,
			Unit:        unit,
			ImageUrl:    url,
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
	query := "SELECT id ,phone, name, productiontype,productionName , amount,unit, imageUrl,government, usertype , area, date FROM " + db_name + ".post WHERE phone = ?"
	result, err := dbb.Query(query, phone)
	if err != nil {
		fmt.Println("error ", err.Error())

	}
	var id, amount int
	var userPhone, name, productiontype, productionName, unit, url, government, usertype, area, date string

	for result.Next() {
		err = result.Scan(&id, &userPhone, &name, &productiontype, &productionName, &amount, &unit, &url, &government, &usertype, &area, &date)
		if err != nil {
			fmt.Println(err.Error())
		}
		post := db.Post{
			PostID:      id,
			Phone:       userPhone,
			Name:        name,
			ProductType: productiontype,
			ProductName: productionName,
			Amount:      amount,
			Unit:        unit,
			ImageUrl:    url,
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
	query := "SELECT id ,phone, name, productiontype,productionName , amount, unit ,imageUrl, government, usertype , area, date FROM " + db_name + ".post WHERE usertype = ?"
	result, err := dbb.Query(query, userType)
	if err != nil {
		fmt.Println("error ", err.Error())

	}
	var id, amount int
	var userPhone, name, productiontype, productionName, unit, url, government, usertype, area, date string

	for result.Next() {
		err = result.Scan(&id, &userPhone, &name, &productiontype, &productionName, &amount, &unit, &url, &government, &usertype, &area, &date)
		if err != nil {
			fmt.Println(err.Error())
		}
		post := db.Post{
			PostID:      id,
			Phone:       userPhone,
			Name:        name,
			ProductType: productiontype,
			ProductName: productionName,
			Amount:      amount,
			Unit:        unit,
			ImageUrl:    url,
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
	query := "SELECT id,phone, name, productiontype ,productionName, amount,unit , imageUrl , government, usertype , area, date FROM " + db_name + ".post "
	result, err := dbb.Query(query)
	if err != nil {
		fmt.Println("error ", err.Error())

	}
	var id, amount int
	var userPhone, name, productiontype, productionName, unit, url, government, usertype, area, date string

	for result.Next() {
		err = result.Scan(&id, &userPhone, &name, &productiontype, &productionName, &amount, &unit, &url, &government, &usertype, &area, &date)
		if err != nil {
			fmt.Println(err.Error())
		}
		post := db.Post{
			PostID:      id,
			Phone:       userPhone,
			Name:        name,
			ProductType: productiontype,
			ProductName: productionName,
			Amount:      amount,
			Unit:        unit,
			ImageUrl:    url,
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
func DeletePostDb(id string) bool {
	dbb := db.DBConn()
	defer dbb.Close()
	db_name := os.Getenv("DB_NAME")
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
	query := "UPDATE " + db_name + ".post set phone =?, name =?,productiontype =?,productionName=? ,amount =?,unit=? ,imageUrl=? , government =?, usertype =?, area =?, date =? WHERE id =?"
	update, err := db.Prepare(query)
	if err != nil {
		panic(err.Error())
	}
	_, err = update.Exec(post.Phone, post.Name, post.ProductType, post.ProductName, post.Amount, post.Unit, post.ImageUrl, post.Government, post.UserType, post.Area, post.Date, id)
	if err != nil {
		panic(err.Error())
	}
	return true
}

//done
func GetPostsWithSearch(text string) []db.Post {

	posts := []db.Post{}
	dbb := db.DBConn()
	defer dbb.Close()
	//query := "SELECT * FROM post WHERE productionName LIKE /%?/% OR productiontype LIKE %?% OR name LIKE %?% OR phone LIKE %?% OR government LIKE %?% OR area LIKE %?%"
	result, err := dbb.Query("SELECT * FROM post WHERE productionName LIKE concat('%', ?, '%') OR productiontype LIKE concat('%', ?, '%')  OR phone LIKE concat('%', ?, '%') OR name LIKE concat('%', ?, '%') OR amount LIKE concat('%', ?, '%') OR government LIKE concat('%', ?, '%') OR area LIKE concat('%', ?, '%') OR usertype LIKE concat('%', ?, '%')", text, text, text, text, text, text, text, text)
	//result, err := dbb.Query(query, text, text, text, text, text, text)
	if err != nil {
		fmt.Println("error ", err.Error())

	}
	var id, amount int
	var userPhone, name, productiontype, productionName, unit, url, government, usertype, area, date string

	for result.Next() {
		err = result.Scan(&id, &userPhone, &name, &productiontype, &productionName, &amount, &unit, &url, &government, &usertype, &area, &date)
		if err != nil {
			fmt.Println(err.Error())
		}
		post := db.Post{
			PostID:      id,
			Phone:       userPhone,
			Name:        name,
			ProductType: productiontype,
			ProductName: productionName,
			Amount:      amount,
			Unit:        unit,
			ImageUrl:    url,
			Government:  government,
			UserType:    usertype,
			Area:        area,
			Date:        date,
		}
		posts = append(posts, post)
	}
	return posts
}

func CreateReport() string {
	var report string
	report ="you sell "
	dbb := db.DBConn()
	defer dbb.Close()
	arr := [7]string{"Bag", "Box", "carton", "packet", "bottle", "can", "bar"}
	for i := 0; i < 7; i++ {
		rows, err := dbb.Query("SELECT sum(amount) FROM post where unit = ?", arr[i])
		if err != nil {
			fmt.Println(err.Error())
			return "0"
		} else {
			var sum_quantities int64
			for rows.Next() {
				rows.Scan(&sum_quantities)
			}
			x := fmt.Sprintf(" %s %d ", arr[i], sum_quantities)
			report+=x

		}
	}

	return report

}
