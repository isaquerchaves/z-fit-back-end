package config

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDb() {
	dsn := "host=aws-0-us-east-1.pooler.supabase.com user=postgres.fqbbbpnevfsjzudrujnq password=M9a&iSRKp2?sa5! dbname=postgres port=5432 sslmode=disable"
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the database")
	}

	// Sobrescreve o comportamento padrão do GORM para desativar soft delete globalmente
	DB = DB.Session(&gorm.Session{SkipDefaultTransaction: true}).Unscoped()
}
