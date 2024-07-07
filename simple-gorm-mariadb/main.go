package main

import (
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:changeme@tcp(127.0.0.1:3306)/golang_101?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("open db error: %v", err)
	}

	// Get the underlying sql.DB object to close the connection later
	sqlDB, err := db.DB()
	if err != nil {
		log.Printf("get db error: %v", err)
		return
	}
	defer sqlDB.Close()

	sqlDB.SetMaxOpenConns(500)
	sqlDB.SetMaxIdleConns(100)
	sqlDB.SetConnMaxLifetime(5 * time.Minute)

	// Ping the database to check if the connection is successful
	err = sqlDB.Ping()
	if err != nil {
		log.Printf("ping db error: %v", err)
		return
	}
	log.Println("Database connection successful")

	// Perform auto migration
	err = db.AutoMigrate()
	if err != nil {
		log.Printf("auto migrate error: %v", err)
		return
	}
	log.Println("Auto migration completed")
}
