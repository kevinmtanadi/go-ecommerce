package model

import (
	"fmt"
	"log"

	"gorm.io/gorm"
)

var tables []Table

type Table interface {
	TableName() string
}

func GenerateTableList() {
	tables = append(tables,
		&User{},
		&Store{},
		&Product{},
		&ProductAdditionalData{},
		&ProductRating{},
		&CartContent{},
		&MasterProvince{},
		&MasterCity{},
		&MasterKelurahan{},
		&Order{},
		&Shipment{},
		&Transaction{},
		&TransactionDetail{},
		&UserAddress{},
		&Voucher{},
		&Wallet{},
		&Wishlist{},
	)
}

func MigrateTable(db *gorm.DB) {
	GenerateTableList()

	for _, table := range tables {
		err := db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(table)

		created := db.Migrator().HasTable(table.TableName())
		if created && err == nil {
			log.Println(fmt.Sprintf("Table %s created", table.TableName()))
		} else {
			log.Println(fmt.Sprintf("Error creating table %s: %s", table.TableName(), err.Error()))
		}
	}
}
