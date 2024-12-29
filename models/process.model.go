package models

import (
	"fmt"
	"strconv"

	"gorm.io/gorm"
)

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

func (m ProcessPage) ToView() View {
	return View{
		"ID":         fmt.Sprintf("%d", m.ID),
		"TemplateID": fmt.Sprintf("%d", m.TemplateID),
	}
}

func (m ProcessDTO) ToView() View {
	return View{
		"TemplateID": fmt.Sprintf("%d", m.TemplateID),
	}
}

func (m ProcessDTO) FromView(view View) (IDTO, error) {
	templateID, err := strconv.Atoi(view["TemplateID"])
	if err != nil {
		return nil, err
	}
	m.TemplateID = uint(templateID)
	return m, nil
}
