package config

import (
	"eulabs_challenger/internal/model"
	"fmt"
	"os"
)

type config struct {
	cfg *model.Config
}

func NewConfig() *config {
	return &config{}
}

func (c *config) GetConfig() *model.Config {

	c.cfg.RestPort = os.Getenv("REST_PORT")

	c.cfg.DBConfig = model.DBConfig{
		Username: os.Getenv("DB_USERNAME"),
		Password: os.Getenv("DB_PASSWORD"),
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Database: os.Getenv("DB_DATABASE"),
	}

	dbConfig := c.cfg.DBConfig

	c.cfg.DBConfig.ConnString = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", dbConfig.Username, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.Database)

	return c.cfg
}
