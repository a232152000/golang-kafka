package repository

import (
	"fmt"
	"golang-kafka/models"
	"golang-kafka/util/database"
	"golang-kafka/util/log"
)

//var db = database.GetDB()

func CreateUser() {
	db := database.GetDB()

	fmt.Println(db)
	if err := db.AutoMigrate(&models.User{}); err != nil {
		log.Errorf("table建立失敗: %v", err)
	}
}

func InsertUser() {
	db := database.GetDB()

	newUser := models.User{Name: "Alice", Email: "alice@example.com", Age: 25}
	result := db.Create(&newUser)
	if result.Error != nil {
		log.Errorf("insert失敗: %v", result.Error)
	}
}

func GetAllUser() []models.User {
	db := database.GetDB()

	var users []models.User
	db.Find(&users)

	return users
}
