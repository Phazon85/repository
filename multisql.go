package multisql

import (
	"database/sql"
	"log"

	//used for testing in postgres_test.go
	_ "github.com/lib/pq"
	"github.com/phazon85/multisql/config"
)

//SQLService holds connection blob from NewDBObject
type SQLService struct {
	DB *sql.DB
}

//NewDBObject takes in file name and driver name and return an appropriate DB connection string
// depending on which driver name is specified - Currently supports Postgresql only
func NewDBObject(file string, driverName string) *SQLService {
	cfg := config.NewConfig(driverName, file)
	if cfg == "" {
		log.Printf("Could not properly format string: %s", err.Error())
	}
	db, err := sql.Open(driverName, cfg)
	if err != nil {
		log.Printf("Error opening SQL db: %s", err.Error())
	}
	err = db.Ping()
	if err != nil {
		log.Printf("Error pinging SQL DB: %s", err.Error())
	}
	return &SQLService{
		DB: db,
	}
}
