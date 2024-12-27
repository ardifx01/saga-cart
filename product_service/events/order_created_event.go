package events

type OrderCreatedEvent struct {
	ProductID int `json:"product_id"`
	Qty       int `json:"qty"`
}
