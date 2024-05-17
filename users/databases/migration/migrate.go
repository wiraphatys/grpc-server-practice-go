package main

import (
	"fmt"

	"user.services/config"
	"user.services/databases"
	"user.services/entities"
)

func main() {
	cfg := config.GetConfig()
	db := databases.NewPostgresDatabase(cfg)

	// migrate schema
	if err := db.GetDb().AutoMigrate(&entities.User{}); err != nil {
		fmt.Printf("Error while migrate user table: %v", err.Error())
	}

	fmt.Println("migrate successful")
}
