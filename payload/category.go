package payload

type CategoryPayload struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
}
