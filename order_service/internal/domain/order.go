package domain

import "time"

type Order struct {
	ID           int       `gorm:"primaryKey" json:"id"`
	CustomerName string    `gorm:"size:255;not null" json:"customer_name"`
	ProductID    int       `gorm:"not null" json:"product_id"`
	Qty          int       `gorm:"not null" json:"qty"`
	OrderDate    time.Time `gorm:"not null" json:"order_date"`
	Status       string    `gorm:"size:50;not null" json:"status"`
}
