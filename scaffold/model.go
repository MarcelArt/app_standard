package scaffold

import (
	"fmt"
	"os"
	"strings"
)

const modelTemplate = `
package models

import "gorm.io/gorm"

const ${modelLower}TableName = "${modelPlural}"

type ${modelName} struct {
	gorm.Model
	// Insert your fields here
}

type ${modelName}DTO struct {
	DTO
	// Insert your fields here
}

type ${modelName}Page struct {
	// Insert your fields here
}

func (${modelName}DTO) TableName() string {
	return ${modelLower}TableName
}

`

func ScaffoldModel(modelName string) {
	modelLower := strings.ToLower(modelName)
	modelPlural := modelName + "s" // need to handle other english plural grammar

	newModel := strings.ReplaceAll(modelTemplate, "${modelLower}", modelLower)
	newModel = strings.ReplaceAll(newModel, "${modelPlural}", modelPlural)
	newModel = strings.ReplaceAll(newModel, "${modelName}", modelName)

	filename := fmt.Sprintf("models/%s.model.go", modelLower)
	if err := os.WriteFile(filename, []byte(newModel), 0644); err != nil {
		panic("Failed writing model file")
	}
}
