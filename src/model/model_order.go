package model

type Order struct {
}

func (Order) TableName() string {
	return "order"
}
