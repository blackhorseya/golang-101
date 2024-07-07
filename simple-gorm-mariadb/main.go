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

	// Get the underlying sql.DB object to close the connection later
	sqlDB, err := db.DB()
	if err != nil {
		log.Printf("get db error: %v", err)
		return
	}
	defer sqlDB.Close()

	// Ping the database to check if the connection is successful
	err = sqlDB.Ping()
	if err != nil {
		log.Printf("ping db error: %v", err)
		return
	}
	log.Println("Database connection successful")

	// Perform auto migration
	err = db.AutoMigrate(&Order{}, &OrderItem{})
	if err != nil {
		log.Printf("auto migrate error: %v", err)
		return
	}
	log.Println("Auto migration completed")

	// Create a new order
	order := &Order{
		Items: []*OrderItem{
			{ItemID: 1, Quantity: 2, Price: 10.0},
			{ItemID: 2, Quantity: 1, Price: 20.0},
		},
	}

	// Save the order using transaction
	tx := db.Begin()
	if tx.Error != nil {
		log.Printf("begin tx error: %v", tx.Error)
		return
	}

	err = tx.Create(order).Error
	if err != nil {
		tx.Rollback()
		log.Printf("create order error: %v", err)
		return
	}

	err = tx.Commit().Error
	if err != nil {
		log.Printf("commit tx error: %v", err)
		return
	}
	log.Printf("order created: %v", order.ID)

	// todo: 2024/7/7|sean|list all orders

	// todo: 2024/7/7|sean|get order by id
}
