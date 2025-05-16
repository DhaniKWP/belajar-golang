package main

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	ID uint `gorm:"primaryKey"`
	Name string
	Email string
}

func main() {
	dsn := "host=localhost user=postgres password=admin dbname=testdb port=5432 sslmode=disable"
	db, err :=gorm.Open(postgres.Open(dsn),&gorm.Config{})

	if err != nil {
		panic("Failed to connect to database!")
	}


	// Auto migrate schema
	err = db.AutoMigrate(&User{})
	if err != nil {
	panic("Failed to migrate database!")
	}

	// insert sample database
	db.Create(&User{Name:"Dhani", Email:"Dhanikusuma256@gmail.com"})

	// Read first user 
	var user User 
	db.First(&user) 
	fmt.Println("User:", user) 
} 

