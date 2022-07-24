package configuration

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB gorm.DB

func ConnectDatabase() (DB *gorm.DB) {
	log.Println("Connect to database " + infoDB())

	dsn := fmt.Sprintf("host=%v port=%v user=%v password=%v dbname=%v sslmode=%v TimeZone=%v",
		Conf.Database.Host, Conf.Database.Port, Conf.Database.Username, Conf.Database.Password,
		Conf.Database.Name, Conf.Database.SSLMode, Conf.Database.TimeZone)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Panic(fmt.Sprintf("Cannot Connect Database %v", err))
	}

	DB = db
	return
}

func infoDB() string {
	return fmt.Sprintf("%v@%v:%v/%v?sslMode=%v&timeZone=%v",
		Conf.Database.Username, Conf.Database.Host, Conf.Database.Port, Conf.Database.Name, Conf.Database.SSLMode, Conf.Database.TimeZone)
}
