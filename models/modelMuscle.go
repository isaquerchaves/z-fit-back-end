package models

import "gorm.io/gorm"

type Muscle struct {
	gorm.Model
	ID        string `gorm:"primaryKey"`
	Name      string `gorm:"column:name"`
	Slug      string `gorm:"column:slug"`
	ImageURL  string `gorm:"column:imageUrl"`
	DeletedAt string `gorm:"column:deleted_at"`
}

func (Muscle) TableName() string {
	return "Muscle"
}

func (Muscle) SchemaName() string {
	return "public"
}
