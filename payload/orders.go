package payload

type OrdersPayload struct {
	CustomerId int                  `json:"customer_id" binding:"required"`
	ShippedAt  string               `json:"shipped_at" binding:"required"`
	Order      []OrderDetailPayload `json:"orders" binding:"required"`
}
