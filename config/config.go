package config

import (
	"errors"
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

const (
	driverNotSupported    = "Drivername not supported or nil"
	psqlConnectionString  = "host=%s port=%d user=%s password=%s dbname=%s sslmode=disable"
	nosqlConnectionString = "mongodb://%s:%d"
)

//SQLConnectionInfo holds connection info for SQL implementations
type SQLConnectionInfo struct {
	Host     string `yaml:"host"`
	Port     int64  `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Name     string `yaml:"name"`
	DBName   string `yaml:"dbname"`
}

//NoSQLConnectionInfo holds connection info for a NoSQL Implementation
type NoSQLConnectionInfo struct {
	URI        string `yaml:"uri"`
	Port       int    `yaml:"port"`
	DBName     string `yaml:"dbname"`
	Collection string `yaml:"collection"`
}

// NewConfig loads the info from the config file
func NewConfig(driverName string, file string) (string, error) {
	switch driverName {
	case "postgres":
		cfg := &SQLConnectionInfo{}
		if err := load(cfg, file); err != nil {
			return "", err
		}
		psql := fmt.Sprintf(psqlConnectionString, cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DBName)
		return psql, nil
	case "mongodb":
		cfg := &NoSQLConnectionInfo{}
		if err := load(cfg, file); err != nil {
			return "", err
		}
		nosql := fmt.Sprintf(nosqlConnectionString, cfg.URI, cfg.Port)
		return nosql, nil
	default:
		err := errors.New(driverNotSupported)
		return "", err

	}
}

func load(config interface{}, fname string) error {
	data, err := ioutil.ReadFile(fname)
	if err != nil {
		return err
	}

	if err := yaml.Unmarshal(data, config); err != nil {
		return err
	}
	return nil
}
