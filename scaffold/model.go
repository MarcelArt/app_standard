package scaffold

import (
	"fmt"
	"os"
	"strings"
)

const modelTemplate = `
package models

import "gorm.io/gorm"

const ${modelCamel}TableName = "${modelPlural}"

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
	return ${modelCamel}TableName
}

`

func ScaffoldModel(modelName string, modelCamel string, modelSnake string) {
	modelPlural := PluralizeWord(modelSnake)
	newModel := strings.ReplaceAll(modelTemplate, "${modelCamel}", modelCamel)
	newModel = strings.ReplaceAll(newModel, "${modelPlural}", modelPlural)
	newModel = strings.ReplaceAll(newModel, "${modelName}", modelName)

	filename := fmt.Sprintf("models/%s.model.go", modelSnake)
	if err := os.WriteFile(filename, []byte(newModel), 0644); err != nil {
		panic("Failed writing model file")
	}
}
