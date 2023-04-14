package configs

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// postgres database konektor
func GormPostgresConn() *gorm.DB {
	configDB := map[string]string{
		"DB_Username": Config("DB_USERNAME"),
		"DB_Password": Config("DB_PASSWORD"),
		"DB_Host":     Config("DB_ADDRESS"),
		"DB_Name":     Config("DB_NAME"),
		"DB_PORT":     Config("DB_PORT"),
	}

	openConnection := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
		configDB["DB_Host"],
		configDB["DB_Username"],
		configDB["DB_Password"],
		configDB["DB_Name"],
		configDB["DB_PORT"],
	)

	// Open the connection
	db, err := gorm.Open(postgres.Open(openConnection), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	log.Println("DATABASE Successfully connected!")
	// return the connection
	return db
}
