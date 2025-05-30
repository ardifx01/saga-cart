package events

type StockReservedEvent struct {
	ID           int     `json:"id"`
	ProductID    int     `json:"product_id"`
	Qty          int     `json:"qty"`
	JumlahBeli   int     `json:"jumlah_beli"`
	Amount       float64 `json:"amount"`
	HargaProduct float64 `json:"harga_product"`
}
