package multisql

import (
	"database/sql"
	"log"

	//used for testing in postgres_test.go
	_ "github.com/lib/pq"
	"github.com/phazon85/multisql/config"
)

//NewDBObject takes in file name and driver name and return an appropriate DB connection object
// depending on which driver name is specified - Currently supports Postgresql only
func NewDBObject(file string, driverName string) *sql.DB {
	cfg, err := config.NewConfig(driverName, file)
	if err != nil {
		log.Printf("Error crearing connection string, %s", err)
	}
	db, err := sql.Open(driverName, cfg)
	if err != nil {
		log.Printf("Error opening SQL db: %s", err.Error())
	}
	err = db.Ping()
	if err != nil {
		log.Printf("Error pinging SQL DB: %s", err.Error())
	}
	log.Println("Succesfully connected.")
	return db
}
