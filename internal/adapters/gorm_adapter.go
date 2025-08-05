package adapters

import (
	"fmt"
	"go-clean-achitech/internal/configs"

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

type GormAdapter struct {
	*gorm.DB
}

var adapterInstant *GormAdapter

func NewGormAdapter() (*GormAdapter, error) {

	if adapterInstant != nil {
		return adapterInstant, nil
	}

	env := configs.GetEnv()
	dsn := fmt.Sprintf("sqlserver://%s:%s@%s:%s?database=%s", env.DBUserName, env.DBPassword, env.DBHost, env.DBPort, env.DBName)
	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	adapterInstant = &GormAdapter{
		DB: db,
	}

	return adapterInstant, nil
}
