package events

type OrderCreatedEvent struct {
	ID        int     `json:"id"`
	ProductID int     `json:"product_id"`
	Amount    float64 `json:"amount"`
	Qty       int     `json:"qty"`
}
