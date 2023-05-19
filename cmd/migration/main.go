package main

import (
	"eulabs_challenger/cmd/migration/migrations"
	"eulabs_challenger/internal/config"
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func main() {
	cfg := config.NewConfig()

	db, err := gorm.Open(mysql.Open(cfg.GetConfig().DBConfig.ConnString), &gorm.Config{})
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
