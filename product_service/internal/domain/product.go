package domain

import "time"

type Product struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Name      string    `gorm:"size:255;not null" json:"name"`
	Price     float64   `gorm:"type:numeric(10,2);not null" json:"price"`
	Quantity  int       `gorm:"not null" json:"quantity"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
