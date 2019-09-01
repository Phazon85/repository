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
		ps := NewDBObject(postgresFile, "postgres")
		assert.NotNil(t, ps)
	})
}

// func TestInitDB(t *testing.T) {
// 	t.Run("Postgres DB Init", func(t *testing.T) {
// 		ps := NewDBObject(postgresFile, "postgres")
// 		db := ps.InitDB()
// 		assert.NotNil(t, db)
// 	})
// }
