package storage

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Config struct {
	DBName   string
	Password string
}

func NewConnection(config *Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("root:%s@tcp(localhost:3306)/%s", config.Password, config.DBName)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		return db, err
	}

	return db, err
}