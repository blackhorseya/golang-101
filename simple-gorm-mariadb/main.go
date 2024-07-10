package main

import (
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	dsn          = "root:changeme@tcp(127.0.0.1:3306)/golang_101?charset=utf8mb4&parseTime=True&loc=Local"
	defaultLimit = 100
)

func setupDatabase(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	sqlDB.SetMaxOpenConns(500)
	sqlDB.SetMaxIdleConns(100)
	sqlDB.SetConnMaxLifetime(5 * time.Minute)

	if err = sqlDB.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

func main() {
	db, err := setupDatabase(dsn)
	if err != nil {
		log.Printf("setup database error: %v", err)
		return
	}
	log.Println("Database connected")

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
	err = createOrder(db, order)
	if err != nil {
		log.Printf("create order error: %v", err)
		return
	}
	log.Println("Order created successfully with ID:", order.ID)

	// List all orders
	orders, total, err := listOrders(db, listCondition{
		Limit:  10,
		Offset: 0,
	})
	if err != nil {
		log.Printf("list orders error: %v", err)
		return
	}
	for _, o := range orders {
		log.Printf("order: %+v\n", o)
	}
	log.Printf("total orders: %d\n", total)

	// Get order by ID
	orderID := order.ID
	order, err = getOrderByID(db, orderID)
	if err != nil {
		log.Printf("get order by ID error: %v", err)
		return
	}
	log.Printf("order with ID %d: %+v\n", orderID, order)
}

func createOrder(db *gorm.DB, order *Order) (err error) {
	tx := db.Begin()
	if tx.Error != nil {
		return tx.Error
	}
	defer func() {
		if r := recover(); r != nil || err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	return tx.Create(order).Error
}

type listCondition struct {
	Limit  int
	Offset int
}

func listOrders(db *gorm.DB, cond listCondition) (items []*Order, total int, err error) {
	if cond.Limit <= 0 {
		cond.Limit = defaultLimit
	}
	if cond.Offset < 0 {
		cond.Offset = 0
	}

	var count int64
	var orders []*Order
	err = db.Model(&Order{}).Limit(cond.Limit).Offset(cond.Offset).Order("id DESC").Count(&count).Find(&orders).Error
	if err != nil {
		return nil, 0, err
	}

	return orders, int(count), nil
}
func getOrderByID(db *gorm.DB, orderID int64) (*Order, error) {
	order := new(Order)
	err := db.Preload("Items").First(order, orderID).Error
	if err != nil {
		return nil, err
	}
	return order, nil
}
