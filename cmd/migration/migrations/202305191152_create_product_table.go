package migrations

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

func init() {
	newMigration := &gormigrate.Migration{
		ID: "202305191152_create_product_table",
		Migrate: func(tx *gorm.DB) error {

			sql := `CREATE TABLE IF NOT EXISTS products (
					id INT AUTO_INCREMENT PRIMARY KEY,
					name VARCHAR(255) NOT NULL,
					description VARCHAR(255) NOT NULL,
					price FLOAT NOT NULL,
					created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
					updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
				)`
			if err := tx.Exec(sql).Error; err != nil {
				return err
			}

			return nil

		},
		Rollback: func(tx *gorm.DB) error {
			return tx.Migrator().DropTable("products")
		},
	}

	MigrationsToExec = append(MigrationsToExec, newMigration)
}
