package main

import (
	"eulabs_challenger/cmd/migration/migrations"
	"eulabs_challenger/internal/config"
	"github.com/go-gormigrate/gormigrate/v2"
	_ "github.com/go-gormigrate/gormigrate/v2"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func main() {
	logger := zap.NewExample()
	defer logger.Sync()

	cfg := config.NewConfig(logger).GetConfig()

	db, err := gorm.Open(mysql.Open(cfg.DBConfig.ConnString), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	migrationsToExec := migrations.GetMigrationsToExec()
	m := gormigrate.New(db, gormigrate.DefaultOptions, migrationsToExec)
	if err := m.Migrate(); err != nil {
		log.Fatalln("Could not migrate: ", err)
	}

	log.Println("Migration did run successfully")

}
