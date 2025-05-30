package domain

type Payment struct {
	ID        int     `gorm:"primaryKey" json:"id"`
	ProductID int     `gorm:"not null" json:"product_id"`
	Amount    float64 `gorm:"type:numeric(10,2);not null" json:"price"`
	Status    string  `gorm:"size:50;not null" json:"status"`
}