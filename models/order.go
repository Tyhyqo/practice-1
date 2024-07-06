package models

import (
	"time"
)

type OrderDTO struct {
	OrderId   int       `gorm:"type:int;not null;auto_increment;primaryKey"`
	CreatedAt time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP"`
	DeliverAt time.Time `gorm:"type:timestamp"`
	Address   string    `gorm:"type:varchar(200);not null"`
	Client    UserDTO   `gorm:"type:varchar(200);not null;ForeignKey:Login"`
	Executor  UserDTO   `gorm:"type:varchar(200);not null;ForeignKey:Login"`
	Status    string    `gorm:"type:varchar(20);not null"`
}

type OrderWeb struct {
	DeliverAt time.Time `json:"deliver_at"`
	Address   string    `json:"address"`
	Client    UserDTO   `json:"client"`
	Executor  UserDTO   `json:"executor"`
	Status    string    `json:"status"`
}
