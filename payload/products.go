package payload

type ProductsPayload struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
	Stock       int    `json:"stock" binding:"required"`
	Price       int    `json:"price" binding:"required"`
	CategoryId  int    `json:"category_id" binding:"required"`
}
