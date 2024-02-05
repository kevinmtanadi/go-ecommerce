package model

import (
	"gorm.io/gorm"
)

var tables []interface{}

func GenerateTableList() {
	tables = append(tables, &User{})
}

func MigrateTable(db *gorm.DB) {
	GenerateTableList()
	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(tables...)
}
