package db_test

import (
	"imagera/internal/db"
	"imagera/internal/db/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testUser models.User
var testImage models.Image

func setup(t *testing.T) {
	err := db.Connect()
	if err != nil {
		t.Fatalf("Failed to connect to database: %v", err)
	}

	testUser = models.User{
		Username:     "imageuser",
		Email:        "imageuser@example.com",
		PasswordHash: "hashedpassword",
	}

	tx := db.DB.Begin()
	err = tx.Create(&testUser).Error
	if err != nil {
		t.Fatalf("Failed to create user: %v", err)
	}

	testImage = models.Image{
		UserID:      testUser.ID,
		Filename:    "test_image.png",
		Size:        1024,
		UploadURL:   "http://example.com/test_image.png",
		Description: "Test image description",
		Tags:        "test, image",
	}

	tx.Commit()
}

func TestCreateImage(t *testing.T) {
	setup(t)

	tx := db.DB.Begin()

	err := tx.Create(&testImage).Error
	assert.Nil(t, err)
	assert.NotNil(t, testImage.ID)

	tx.Rollback()
}

func TestGetImage(t *testing.T) {
	setup(t)

	tx := db.DB.Begin()

	var fetchedImage models.Image
	err := tx.First(&fetchedImage, "id = ?", testImage.ID).Error
	assert.Nil(t, err)
	assert.Equal(t, testImage.Filename, fetchedImage.Filename)

	tx.Rollback()
}

func TestUpdateImage(t *testing.T) {
	setup(t)

	tx := db.DB.Begin()

	newDescription := "Updated image description"
	err := tx.Model(&testImage).Update("description", newDescription).Error
	assert.Nil(t, err)

	var updatedImage models.Image
	err = tx.First(&updatedImage, "id = ?", testImage.ID).Error
	assert.Nil(t, err)
	assert.Equal(t, newDescription, updatedImage.Description)

	tx.Rollback()
}

func TestDeleteImage(t *testing.T) {
	setup(t)

	tx := db.DB.Begin()

	err := tx.Delete(&testImage).Error
	assert.Nil(t, err)

	var deletedImage models.Image
	err = tx.First(&deletedImage, "id = ?", testImage.ID).Error
	assert.NotNil(t, err)

	tx.Rollback()
}
