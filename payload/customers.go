package payload

type CustomersPayload struct {
	Email      string `json:"email" binding:"required"`
	Name       string `json:"name" binding:"required"`
	Address    string `json:"address" binding:"required"`
	City       string `json:"city" binding:"required"`
	State      string `json:"state" binding:"required"`
	PostalCode string `json:"postal_code" binding:"required"`
	Country    string `json:"country" binding:"required"`
	Phone      string `json:"phone" binding:"required"`
}
