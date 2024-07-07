package main

// Order is a struct that represents the order table in the database
type Order struct {
	ID    int64        `gorm:"primaryKey;autoIncrement"`
	Items []*OrderItem `gorm:"foreignKey:OrderID;references:ID"`
}

// OrderItem is a struct that represents the order_item table in the database
type OrderItem struct {
	OrderID  int64   `gorm:"primaryKey"`
	ItemID   int64   `gorm:"primaryKey"`
	Quantity uint    `gorm:"not null"`
	Price    float64 `gorm:"not null"`
}
