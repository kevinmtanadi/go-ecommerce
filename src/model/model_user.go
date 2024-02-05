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
	IsVerified uint8 `gorm:"default:0"`
}

func (User) TableName() string {
	return "users"
}
