package payload

type OrderDetailPayload struct {
	ProductId int `json:"product_id" binding:"required"`
	Qty       int `json:"qty" binding:"required"`
}
