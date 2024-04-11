package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDb() {
	// Carregar variáveis de ambiente do arquivo .env
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	dsn := "host=" + os.Getenv("aws-0-us-east-1.pooler.supabase.com") + " user=" + os.Getenv("postgres.fqbbbpnevfsjzudrujnq") + " password=" + os.Getenv("M9a&iSRKp2?sa5!") + " dbname=" + os.Getenv("postgres") + " port=" + os.Getenv("5432") + " sslmode=" + os.Getenv("disable")

	// Conectar ao banco de dados
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the database")
	}
}
