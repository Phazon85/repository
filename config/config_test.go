package config

import "testing"

const (
	postgresDriverName = "postgres"
	postgresFile       = "../tests/postgres.yaml"
	mysqlFile          = "../tests/mysql.yaml"
)

func TestNewConfig(t *testing.T) {
	t.Run("Postgres config", func(t *testing.T) {
		driver, cfg := NewConfig(postgresFile)
		driverwant := "postgres"
		cfgwant := "host=localhost port=5432 user=postgres password=changeme dbname=test sslmode=disable"
		if cfg != cfgwant || driver != driverwant {
			t.Errorf("Error loading config, got: %s %s, want %s %s", cfg, driver, cfgwant, driverwant)
		}
	})

	t.Run("MySQL config", func(t *testing.T) {
		driver, cfg := NewConfig(mysqlFile)
		driverwant := "mysql"
		cfgwant := "mysql:changeme@localhost/test"
		if cfg != cfgwant || driver != driverwant {
			t.Errorf("Error loading config, got: %s %s, want %s %s", cfg, driver, cfgwant, driverwant)
		}
	})

	t.Run("Nil file name", func(t *testing.T) {
		_, cfg := NewConfig("")
		if cfg != "" {
			t.Errorf("Wanted to get an empty string, but didn't")
		}
	})
}
