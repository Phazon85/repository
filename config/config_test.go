package config

import (
	"testing"
)

const (
	postgresFile = "../tests/postgres.yaml"
)

func TestNewConfig(t *testing.T) {
	t.Run("Postgres config", func(t *testing.T) {
		cfg, _ := NewConfig("postgres", postgresFile)
		want := "host=localhost port=5432 user=postgres password=changeme dbname=test sslmode=disable"
		if cfg != want {
			t.Errorf("Error loading config, got: %s, want %s", cfg, want)
		}
	})

	t.Run("Nil driveName", func(t *testing.T) {
		_, err := NewConfig("", "")
		if err == nil {
			t.Errorf("Error handling nil driverName, wanted error, didn't get one")
		}
	})
}
