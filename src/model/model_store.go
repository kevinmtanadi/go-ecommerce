package model

import "time"

type Store struct {
	ID          uint
	StoreName   string
	OwnerID     uint
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Status      int       `gorm:"default:10"`
	Products    []Product `gorm:"foreignKey:StoreID"`
}

func (Store) TableName() string {
	return "store"
}
