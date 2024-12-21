package models

import "gorm.io/gorm"

const processTableName = "processes"

type Process struct {
	gorm.Model
	TemplateID uint `json:"templateId"`

	Template *Template `json:"template"`
}

type ProcessDTO struct {
	DTO
	TemplateID uint `json:"templateId"`
}

type ProcessPage struct {
	ID           uint   `gorm:"primarykey"`
	TemplateID   uint   `json:"templateId"`
	TemplateName string `json:"templateName"`
}

func (ProcessDTO) TableName() string {
	return processTableName
}
