package database

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DatabaseInit() {

	// inisiasi var error
	var err error

	// membuat var sql untuk koneksi dsn
	const MYSQL = "root:@tcp(127.0.0.1:3306)/fiber_db?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := MYSQL
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Tidak bisa terhubung ke DATABASE")
	}

	fmt.Println("Terhubung ke database")
}
