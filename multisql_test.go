package multisql

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	postgresFile = "tests/postgres.yaml"
)

func TestNewDBObject(t *testing.T) {
	t.Run("Postgres DB test", func(t *testing.T) {
		db := NewDBObject(postgresFile, "postgres")
		assert.NotNil(t, db)
	})
}
