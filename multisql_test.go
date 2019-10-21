package multisql

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	postgresFile = "tests/postgres.yaml"
	mongoFile    = "tests/mongo.yaml"
)

func TestNewSQLDBObject(t *testing.T) {
	t.Run("Postgres DB test", func(t *testing.T) {
		ps := NewSQLDBObject(postgresFile, "postgres")
		assert.NotNil(t, ps)
	})
}

func TestNewNoSQLDBObject(t *testing.T) {
	t.Run("Mongo DB test", func(t *testing.T) {
		mongo := NewNoSQLDBObject(mongoFile, "mongodb")
		assert.NotNil(t, mongo)
	})
}
