package model

import "time"

type Product struct {
	ID              uint
	StoreID         uint
	ProductName     string
	Price           float64
	DiscountedPrice float64
	Description     string
	Stock           int
	AdditionalDatas []ProductAdditionalData `gorm:"foreignKey:ProductID"`
	Rating          []ProductRating         `gorm:"foreignKey:ProductID"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

func (Product) TableName() string {
	return "product"
}

type ProductAdditionalDataType string

var (
	ProductImages ProductAdditionalDataType = "product_images"
	Discussion    ProductAdditionalDataType = "discussion"
)

type ProductAdditionalData struct {
	ID        uint
	ProductID uint
	Key       string
	Value     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (ProductAdditionalData) TableName() string {
	return "product_additional_data"
}

type ProductRating struct {
	ID        uint
	ProductID uint
	Rating    float64
	Comment   string
	UserID    uint
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (ProductRating) TableName() string {
	return "product_rating"
}
