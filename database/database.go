package database

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DatabaseInit() {

	var err error
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	const MYSQL = "root:@tcp(127.0.0.1:3306)/data_db?charset=utf8mb4&parseTime=True&loc=Local"

	dsn := MYSQL
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("cannot connect to database ")
	}

	fmt.Println("connected to database")

}
