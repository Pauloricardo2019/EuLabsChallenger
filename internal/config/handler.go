package config

import (
	"eulabs_challenger/internal/model"
	"fmt"
	"go.uber.org/zap"
	"os"
	"strconv"
	"time"
)

type config struct {
	cfg    model.Config
	logger *zap.Logger
}

func NewConfig(logger *zap.Logger) *config {
	return &config{
		logger: logger,
	}
}

func (c *config) GetConfig() *model.Config {
	c.logger.Info("Init config",
		zap.Time("StartedAt", time.Now()),
	)
	port := os.Getenv("REST_PORT")

	if port != "" {
		portInt, err := strconv.Atoi(port)
		if err != nil {
			c.cfg.RestPort = 9090
		}
		c.cfg.RestPort = portInt

	} else {
		c.cfg.RestPort = 9090
	}

	c.logger.Debug("REST_PORT", zap.Int("REST_PORT", c.cfg.RestPort), zap.Time("StartedAt", time.Now()))

	c.cfg.DBConfig = model.DBConfig{
		Username: os.Getenv("DB_USERNAME"),
		Password: os.Getenv("DB_PASSWORD"),
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Database: os.Getenv("DB_NAME"),
	}

	dbConfig := c.cfg.DBConfig

	c.cfg.DBConfig.ConnString = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", dbConfig.Username, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.Database)

	c.logger.Info("Config loaded",
		zap.Time("StartedAt", time.Now()))

	return &c.cfg
}
