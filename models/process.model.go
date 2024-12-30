package models

import (
	"fmt"
	"strconv"

	"gorm.io/gorm"
)

const processTableName = "processes"

type Process struct {
	gorm.Model
	Name       string `json:"name"`
	TemplateID uint   `json:"templateId"`

	Template *Template `json:"template"`
}

type ProcessDTO struct {
	DTO
	Name       string `json:"name" field:"display=name,key=name"`
	TemplateID uint   `json:"templateId" field:"display=name,key=name"`
}

type ProcessPage struct {
	ID           uint   `gorm:"primarykey"`
	Name         string `json:"name"`
	TemplateID   uint   `json:"templateId"`
	TemplateName string `json:"templateName"`
}

func (ProcessDTO) TableName() string {
	return processTableName
}

func (m ProcessPage) ToView() View {
	return View{
		"ID":         fmt.Sprintf("%d", m.ID),
		"Name":       fmt.Sprintf("%s", m.Name),
		"TemplateID": fmt.Sprintf("%d", m.TemplateID),
	}
}

func (m ProcessDTO) ToView() View {
	return View{
		"Name":       fmt.Sprintf("%s", m.Name),
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
