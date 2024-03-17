package models

import "time"

type User struct {
	ID        uint      `gorm:"column:id;primary_key;autoIncrement" json:"id"`
	Username  string    `gorm:"column:username;not null" json:"username"`
	Email     string    `gorm:"column:email;unique;not null" json:"email"`
	Password  string    `gorm:"column:password;not null" json:"password"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime" json:"createdAt"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoCreateTime;autoUpdateTime" json:"updatedAt"`
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

type UserRegister struct {
	Username string `json:"username" valid:"required~username is required"`
	Email    string `json:"email" valid:"email,required~email is required"`
	Password string `json:"password" valid:"required~password is required"`
}

type UserLogin struct {
	Email    string `json:"email" valid:"email,required~email is required"`
	Password string `json:"password" valid:"required~password is required"`
}
