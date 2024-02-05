package model

import "time"

type Voucher struct {
	ID             uint
	StoreID        uint
	VoucherCode    string
	Discount       float64
	MinTransaction float64
	MaxDiscount    float64
	StartDate      time.Time
	EndDate        time.Time
	CreatedAt      time.Time
}

func (Voucher) TableName() string {
	return "voucher"
}
