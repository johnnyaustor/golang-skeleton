package database

import (
	"fmt"
	"log"

	"github.com/johnnyaustor/golang-skeleton/server/configuration"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DB struct {
	Connection *gorm.DB
}

func ConnectDatabase() *DB {
	log.Println("Connect to database " + infoDB())

	dsn := fmt.Sprintf("host=%v port=%v user=%v password=%v dbname=%v sslmode=%v TimeZone=%v",
		configuration.Conf.Database.Host, configuration.Conf.Database.Port, configuration.Conf.Database.Username, configuration.Conf.Database.Password,
		configuration.Conf.Database.Name, configuration.Conf.Database.SSLMode, configuration.Conf.Database.TimeZone)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Panic(fmt.Sprintf("Cannot Connect Database %v", err))
	}
	dbs := DB{
		Connection: db,
	}
	return &dbs
}

func infoDB() string {
	return fmt.Sprintf("%v@%v:%v/%v?sslMode=%v&timeZone=%v",
		configuration.Conf.Database.Username, configuration.Conf.Database.Host, configuration.Conf.Database.Port, configuration.Conf.Database.Name, configuration.Conf.Database.SSLMode, configuration.Conf.Database.TimeZone)
}
