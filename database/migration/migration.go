package migration

import (
	"fiber-go/database"
	"fiber-go/model/entity"
	"fmt"
	"log"
)

func RunMigration() {
	err := database.DB.AutoMigrate(&entity.User{})

	if err != nil {
		log.Println(err)
	}

	fmt.Println("database Migrated")
}
