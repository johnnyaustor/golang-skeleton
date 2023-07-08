package configuration

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type configuration struct {
	Database database
}

type database struct {
	Host     string
	Port     string
	Username string
	Password string
	Name     string
	SSLMode  string
	TimeZone string
}

var (
	Conf = &configuration{}
)

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func ReadConfiguration() {
	log.Println("Read Configuration...")
	if err := godotenv.Load(); err != nil {
		log.Panic("Cannot read .env file")
	}

	db := &database{}
	{
		db.Host = getEnv("DB_HOST", "localhost")
		db.Port = getEnv("DB_PORT", "5432")
		db.Username = getEnv("DB_USERNAME", "root")
		db.Password = getEnv("DB_PASSWORD", "")
		db.Name = getEnv("DB_NAME", "")
		db.SSLMode = getEnv("DB_SSLMODE", "disable")
		db.TimeZone = getEnv("DB_TIMEZONE", "UTC")
	}

	Conf.Database = *db
}
