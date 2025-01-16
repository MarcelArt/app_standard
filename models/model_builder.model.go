package models

type ModelBuilder struct {
	Column   string `json:"column"`
	JSON     string `json:"json"`
	Gorm     string `json:"gorm"`
	Validate string `json:"validate"`
}

type ModelBuilderRequest struct {
	Model ModelBuilder `json:"model_builder"`
}
