package models

import "gorm.io/gorm"

type TrainingSplit struct {
	gorm.Model
	ID      string `gorm:"primaryKey"`
	Name    string `gorm:"column:name"`
	Enabled string `gorm:"column:enabled"`
}

func (TrainingSplit) TableName() string {
	return "TrainingSplit"
}

func (TrainingSplit) SchemaName() string {
	return "public"
}
