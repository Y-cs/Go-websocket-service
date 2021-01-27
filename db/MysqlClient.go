package db

import (
	"GoTest/constant"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func init() {
	fmt.Println("connection to mysql")
	clientUrl := fmt.Sprintf("%s:%s@tcp(%s)/%s", constant.DbUser, constant.DbPassword, constant.DbUrl, constant.DbSchemas)
	fmt.Println("connection mysql url -> " + clientUrl)
	db, _ = sql.Open(constant.DbDrive, clientUrl)
	//验证连接
	if err := db.Ping(); err != nil {
		fmt.Println("mysql connection error :", err)
		return
	}
	fmt.Println("mysql connection is ready")
}
