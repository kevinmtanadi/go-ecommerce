package model

import "time"

type User struct {
	ID         uint
	Username   string
	Email      string
	Password   string
	Token      string
	Salt       string
	CreatedAt  time.Time
	UpdatedAt  time.Time
	IsVerified uint8           `gorm:"default:0"`
	Stores     []Store         `gorm:"foreignKey:OwnerID"`
	Ratings    []ProductRating `gorm:"foreignKey:UserID"`
}

func (User) TableName() string {
	return "users"
}

type UserAddress struct {
	UserID      uint
	ProvinceID  uint
	CityID      uint
	KelurahanID uint
	FullAddress string
	Note        string
}

func (UserAddress) TableName() string {
	return "user_address"
}
