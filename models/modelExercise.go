package models

import (
	"database/sql/driver"
	"errors"
	"strings"

	"gorm.io/gorm"
)

// Tratando tipo array
// MultiString é um tipo de dados personalizado que representa uma lista de strings.
type MultiString []string

// Scan converte um valor do banco de dados em um tipo MultiString.
func (s *MultiString) Scan(src interface{}) error {
	// Verifica se o valor de origem é uma string
	str, ok := src.(string)
	if !ok {
		return errors.New("failed to scan multistring field - source is not a string")
	}

	// Remove os caracteres "{}" do início e do fim da string
	str = strings.TrimPrefix(str, "{")
	str = strings.TrimSuffix(str, "}")

	// Separa a string pela "," e armazena os resultados em MultiString
	*s = strings.Split(str, ",")
	return nil
}

// Junta os elementos de MultiString em uma única string, separados por ",".
func (s MultiString) Value() (driver.Value, error) {
	// Verifica se MultiString está vazio
	if s == nil || len(s) == 0 {
		return nil, nil
	}
	// Retorna a fatia de strings sem adicionar chaves
	return s, nil
}

type Exercise struct {
	gorm.Model
	ID           string      `gorm:"primaryKey"`
	Enabled      bool        `gorm:"column:enabled"`
	Difficulty   string      `gorm:"column:difficulty"`
	Instructions string      `gorm:"column:instructions"`
	Image        MultiString `gorm:"type:text"`
	Slug         string      `gorm:"column:slug"`
	MuscleID     string      `gorm:"column:muscleId"`
	Name         string      `gorm:"column:name"`
}

func (Exercise) TableName() string {
	return "Exercise"
}

func (Exercise) SchemaName() string {
	return "public"
}
