package order

type CreateOrderRequest struct {
	UserID uint64 `json:"user_id" binding:"required"`
	Price  uint64 `json:"price" binding:"required,gt=0"`
}
