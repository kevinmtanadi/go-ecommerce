package pkg

import (
	"fmt"
	"go-backend-template/src/model"
	"log"
	"os"
	"strconv"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

func NewMysqlClient() (*gorm.DB, error) {
	var (
		conn *gorm.DB
		err  error
	)

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("MYSQL_ROOT_USER"),
		os.Getenv("MYSQL_ROOT_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("MYSQL_PORT"),
		os.Getenv("MYSQL_DATABASE"),
	)

	shouldLog := os.Getenv("LOG_DB")
	if shouldLog == "true" {
		conn, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Error),
			NamingStrategy: schema.NamingStrategy{
				SingularTable: true,
			},
		})
	} else {
		conn, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Silent),
			NamingStrategy: schema.NamingStrategy{
				SingularTable: true,
			},
		})
	}

	if err != nil {
		return conn, err
	}

	db, err := conn.DB()

	if _, found := os.LookupEnv("DB_MAX_OPEN_CONNECTION"); found {
		if maxOpenConnection, err := strconv.Atoi(os.Getenv("DB_MAX_OPEN_CONNECTION")); err == nil {
			db.SetMaxOpenConns(maxOpenConnection)
		}
	}

	if _, found := os.LookupEnv("DB_MAX_IDLE_CONNECTION"); found {
		if maxIdleConnection, err := strconv.Atoi(os.Getenv("DB_MAX_IDLE_CONNECTION")); err == nil {
			db.SetMaxIdleConns(maxIdleConnection)
		}
	}

	if _, found := os.LookupEnv("DB_MAX_LIFETIME"); found {
		if maxLifetime, err := strconv.Atoi(os.Getenv("DB_MAX_LIFETIME")); err == nil {
			db.SetConnMaxLifetime(time.Duration(maxLifetime) * time.Minute)
		}
	}

	if os.Getenv("MIGRATE_MODEL") == "true" {
		model.MigrateTable(conn)
	}

	log.Println(fmt.Sprintf("Connected to database: %s", dsn))
	return conn, err
}
