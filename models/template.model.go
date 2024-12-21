package models

import (
	"gorm.io/gorm"
)

const templateTableName = "templates"

type Template struct {
	gorm.Model
	Name string `gorm:"not null" json:"name"`
}

type TemplateDTO struct {
	DTO
	Name string `gorm:"not null" json:"name"`
}

type TemplatePage struct {
	ID   uint   `gorm:"primarykey"`
	Name string `gorm:"not null" json:"name"`
}

func (TemplateDTO) TableName() string {
	return templateTableName
}

// func (TemplatePage) TableName() string {
// 	return templateTableName
// }
