package cruds

import (
	"gorm.io/gorm"
	"practice/models"
	"practice/utils"
)

func CrudRegisterUser(user models.UserDTO, db *gorm.DB) {
	user.Password = utils.HashPassword(user.Password)
	db.Create(&user)
}
