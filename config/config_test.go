package config

import (
	"testing"
)

const (
	postgresFile = "../tests/postgres.yaml"
)

func TestNewConfig(t *testing.T) {
	t.Run("Postgres config", func(t *testing.T) {
		cfg := NewConfig("postgres", postgresFile)
		want := "host=localhost port=5432 user=postgres password=changeme dbname=test sslmode=disable"
		if cfg != want {
			t.Errorf("Error loading config, got: %s, want %s", cfg, want)
		}
	})
}

// func TestLoad(t *testing.T) {
// 	defer func() {
// 		if r := recover(); r == nil {
// 			t.Errorf("The code did not panic")
// 		}
// 	}()
// 	_ = NewConfig("", "")
// }
