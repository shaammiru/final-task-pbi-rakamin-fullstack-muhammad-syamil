package models

import (
	"time"
)

type User struct {
	ID        uint      `gorm:"column:id;primary_key;autoIncrement"`
	Username  string    `gorm:"column:username;not null"`
	Email     string    `gorm:"column:email;unique;not null"`
	Password  string    `gorm:"column:password;not null"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
	// Relations
	Photos []Photo `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
}

type Photo struct {
	ID        uint      `gorm:"column:id;primary_key;autoIncrement"`
	Title     string    `gorm:"column:title"`
	Caption   string    `gorm:"column:caption"`
	PhotoURL  string    `gorm:"column:photo_url"`
	UserID    uint      `gorm:"column:user_id"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
}
