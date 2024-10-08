package models

import (
	"time"
	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	UserName    string    `gorm:"type:text;not null" json:"username"`
	Amount  float32   `gorm:"type:decimal(10,2);not null" json:"amount"`
	Status  string    `gorm:"type:varchar(50);not null" json:"status"`
	Date    time.Time `gorm:"type:timestamp;not null" json:"date"`
}


