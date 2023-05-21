package config

import (
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

var (
	DBConn *gorm.DB
)

func InitDB(host string, port string, username string, password string, dbname string) error {
	if password == "" {
		password = "@" + password
	}

	// Keyword : Connection String, String Concatenation
	dsn := "sqlserver://DESKTOP-EN0V4OG:1433?database=wbb_dashboard&trusted+connection=yes"
	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	DBConn = db

	return nil

	// Migrate the schema
	// db.AutoMigrate(&models.User{})

	// Create
	// db.Create(&models.User{Name: "Luffy", Age: 7})

	// Read
	// var user models.User
	// db.First(&user, 1) // find user with id 1
	// db.First(&user, "name = ?", "Luffy") // find user with name luffy

	// Update - update user's age to 20
	// db.Model(&user).Update("Age", 20)

	// Delete - delete user
	// db.Delete(&user, 1)

	// db.Close()
}
