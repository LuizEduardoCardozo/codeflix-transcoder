package database_test

import (
	"encoder/framework/database"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIfNewDBTestReturnsNotNil(t *testing.T) {
	db := database.NewTestDB()
	require.NotNil(t, db)
}

func TestIfMigrationWorks(t *testing.T) {
	conn := database.NewTestDB()

	var tableName []string

	conn.Raw("SELECT name FROM sqlite_master WHERE type='table';").Scan(&tableName)

	require.NotEmpty(t, tableName)
	require.Equal(t, 2, len(tableName))
}
