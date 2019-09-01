package config

import (
	"errors"
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

const (
	driverNotSupported   = "Drivername not supported or nil"
	psqlConnectionString = "host=%s port=%d user=%s password=%s dbname=%s sslmode=disable"
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

// NewConfig loads the info from the config file
func NewConfig(driverName string, file string) (string, error) {
	switch driverName {
	case "postgres":
		cfg := &SQLConnectionInfo{}
		if err := load(cfg, file); err != nil {
			return "", err
		}
		psql := fmt.Sprintf(psqlConnectionString, cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Name)
		return psql, nil
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
