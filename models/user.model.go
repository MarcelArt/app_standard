package models

import "gorm.io/gorm"

const userTableName = "users"

type User struct {
	gorm.Model
	// Insert your fields here
	Username string `json:"username" gorm:"unique;not null"`
	Email    string `json:"email" gorm:"unique;not null"`
	Password string `json:"-" gorm:"not null"`
	Salt     string `json:"-" gorm:"not null"`
}

type UserDTO struct {
	DTO
	// Insert your fields here
	Username string `json:"username" gorm:"unique;not null"`
	Email    string `json:"email" gorm:"unique;not null"`
	Password string `json:"password" gorm:"not null"`
	Salt     string `json:"-" gorm:"not null"`
}

type UserPage struct {
	// Insert your fields here
	Username string `json:"username" gorm:"unique;not null"`
	Email    string `json:"email" gorm:"unique;not null"`
}

type LoginInput struct {
	Username  string `json:"username" gorm:"unique;not null"`
	Password  string `json:"password" gorm:"not null"`
	IsRemeber bool   `json:"isRemember"`
}

type LoginResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type RefreshInput struct {
	RefreshToken string `json:"refreshToken"`
}

func (UserDTO) TableName() string {
	return userTableName
}
