package multisql

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	postgresFile = "tests/postgres.yaml"
	mysqlFile    = "tests/mysql.yaml"
)

func TestNewSQL(t *testing.T) {
	t.Run("Postgres DB test", func(t *testing.T) {
		ps := NewSQL(postgresFile)
		assert.NotNil(t, ps)
	})

	t.Run("MySQL DB test", func(t *testing.T) {
		mysql := NewSQL(mysqlFile)
		assert.NotNil(t, mysql)
	})
}
