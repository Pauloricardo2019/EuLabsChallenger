package provider

import (
	"eulabs_challenger/internal/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type databaseProvider struct {
	config *model.Config
}

func NewDatabaseProvider(config *model.Config) *databaseProvider {
	return &databaseProvider{
		config: config,
	}
}

func (d *databaseProvider) Connect() (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(d.config.DBConfig.ConnString), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
