package db_test

import (
	"imagera/internal/db"
	"os"
	"testing"

	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

func setupPostgresEnv() {
	os.Setenv("DB_HOST", "localhost")
	os.Setenv("DB_USER", "testuser")
	os.Setenv("DB_PASSWORD", "testpassword")
	os.Setenv("DB_NAME", "testdb")
	os.Setenv("DB_PORT", "5432")
}

func TestConnect_Success(t *testing.T) {
	setupPostgresEnv()

	err := db.Connect()

	assert.NoError(t, err)
	assert.NotNil(t, db.DB) // db.DBがnilでないことを確認
}

func TestMigrate_Success(t *testing.T) {
	setupPostgresEnv()

	err := db.Connect()
	assert.NoError(t, err)

	db.Migrate()

	assert.NotNil(t, db.DB) // db.DBがnilでないことを確認
}
