package model

type Wallet struct {
	ID     uint
	UserID uint
}

func (Wallet) TableName() string {
	return "wallet"
}
