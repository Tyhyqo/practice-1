package migrations

import (
	"gorm.io/gorm"
	"practice/models"
)

func OnAutoMigrate(db *gorm.DB) {
	db.AutoMigrate(&models.UserDTO{}, models.OrderDTO{})
}
