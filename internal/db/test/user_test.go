package db_test

import (
	"imagera/internal/db"
	"imagera/internal/db/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {
	err := db.Connect()
	if err != nil {
		t.Fatalf("Failed to connect to database: %v", err)
	}

	user := models.User{
		Username:     "testuser",
		Email:        "testuser@example.com",
		PasswordHash: "hashedpassword",
	}

	tx := db.DB.Begin()

	err = tx.Create(&user).Error
	if err != nil {
		t.Errorf("Failed to create user: %v", err)
	}

	tx.Rollback()

	assert.Nil(t, err)
	assert.NotNil(t, user.ID)
}
