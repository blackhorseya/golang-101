package main

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:changeme@tcp(127.0.0.1:3306)/golang_101?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("open db error: %v", err)
	}

	// ping the database to check if the connection is successful
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("get db error: %v", err)
	}
	defer sqlDB.Close()

	err = sqlDB.Ping()
	if err != nil {
		log.Fatalf("ping db error: %v", err)
	}
	log.Println("Database connection successful")

	err = db.AutoMigrate()
	if err != nil {
		log.Fatalf("auto migrate error: %v", err)
	}
	log.Println("Auto migration completed")
}
