package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func DBConn() (db *sql.DB) {

	db_driver := os.Getenv("DB_DRIVER")
	db_User := os.Getenv("DB_ROOT")
	db_Password := os.Getenv("DB_PASSWORD")
	db_Port := os.Getenv("DB_PORT")
	db_Name := os.Getenv("DB_NAME")

	//path:="bb8d3f8c173589:b9a00cff@(us-cdbr-east-04.cleardb.com)/heroku_09b7638a416ab31"
	//path:=fmt.Sprintf(db_User,db_Password,"@tcp",db_Port,"/",db_Name)
	path := db_User + ":" + db_Password + "@tcp" + db_Port + "/" + db_Name
	fmt.Println(path)
	db, err := sql.Open(db_driver, path)
	if err != nil {
		panic(err.Error())
	}
	return db
}
