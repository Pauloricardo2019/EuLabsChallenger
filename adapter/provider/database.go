package provider

import (
	"eulabs_challenger/internal/model"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

type databaseProvider struct {
	config *model.Config
	logger *zap.Logger
}

func NewDatabaseProvider(config *model.Config, logger *zap.Logger) *databaseProvider {
	return &databaseProvider{
		config: config,
		logger: logger,
	}
}

func (d *databaseProvider) Connect() (*gorm.DB, error) {
	d.logger.Info("Init database",
		zap.Time("StartedAt", time.Now()),
	)

	db, err := gorm.Open(mysql.Open(d.config.DBConfig.ConnString), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	d.logger.Info("Database connected",
		zap.Time("StartedAt", time.Now()),
	)
	return db, nil
}
