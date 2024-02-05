package model

type Wishlist struct {
	ID        uint
	UserID    uint
	ProductID uint
}

func (w *Wishlist) TableName() string {
	return "wishlists"
}
