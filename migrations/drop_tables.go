package migrations

import (
	"gorm.io/gorm"
	"practice/models"
)

func DropTables(db *gorm.DB) {
	db.Migrator().DropTable(&models.UserDTO{}, &models.UserWeb{})
}
