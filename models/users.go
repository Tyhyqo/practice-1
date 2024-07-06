package models

import "github.com/golang-jwt/jwt"

type UserDTO struct {
	Login     string `gorm:"type:varchar(20);primaryKey;unique"`
	Password  string `gorm:"type:varchar(100);not null"`
	IsCourier bool   `gorm:"type:boolean;not null"`
}

type UserWeb struct {
	Login     string `json:"login"`
	Password  string `json:"password"`
	IsCourier bool   `json:"isCourier"`
}

type UserClaims struct {
	Login string `json:"login"`
	jwt.StandardClaims
}
