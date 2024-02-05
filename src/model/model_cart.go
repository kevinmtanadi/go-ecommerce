package model

type CartContent struct {
	ID        uint
	UserID    uint
	ProductID uint
	Quantity  int
}

func (CartContent) TableName() string {
	return "cart_content"
}
