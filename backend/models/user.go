package models

import (
    "gorm.io/gorm"
    "time"
)

type User struct {
    gorm.Model
    Email     string    `gorm:"unique;not null" json:"email"`
    Password  string    `gorm:"not null" json:"-"`
    Role      string    `gorm:"type:varchar(20);not null;default:'user'" json:"role"`
    FirstName string    `gorm:"type:varchar(100)" json:"first_name"`
    LastName  string    `gorm:"type:varchar(100)" json:"last_name"`
    IsActive  bool      `gorm:"default:true" json:"is_active"`
    LastLogin time.Time `json:"last_login"`
}

