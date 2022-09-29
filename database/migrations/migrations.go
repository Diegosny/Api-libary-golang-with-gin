package migrations

import (
	"api/model"
	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) {
	err := db.AutoMigrate(model.Book{})
	if err != nil {
		return
	}
}
