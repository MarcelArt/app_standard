package models

import (
	"fmt"

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

func (m TemplatePage) ToView() View {
	return View{
		"ID":   fmt.Sprintf("%d", m.ID),
		"Name": fmt.Sprintf("%s", m.Name),
	}
}

func (m TemplateDTO) ToView() View {
	return View{
		"name": fmt.Sprintf("%s", m.Name),
	}
}

func (m TemplateDTO) FromView(view View) (IDTO, error) {
	m.Name = view["name"]
	return m, nil
}
