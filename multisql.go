package multisql

import (
	"context"
	"database/sql"
	"log"
	"time"

	//used for testing in postgres_test.go
	_ "github.com/lib/pq"
	"github.com/phazon85/multisql/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

//NewSQLDBObject takes in file name and driver name and returns a SQL DB object
func NewSQLDBObject(file string, driverName string) *sql.DB {
	cfg, err := config.NewConfig(driverName, file)
	if err != nil {
		log.Printf("Error creating connection string for SQL: %s", err)
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

//NewNoSQLDBObject takes in file name and driver name and returns a NoSQL DB object
func NewNoSQLDBObject(file string, driverName string) *mongo.Client {
	cfg, err := config.NewConfig(driverName, file)
	if err != nil {
		log.Printf("Error creating connection string for NoSQL: %s", err)
	}
	client, err := mongo.NewClient(options.Client().ApplyURI(cfg))
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}

	return client
}
