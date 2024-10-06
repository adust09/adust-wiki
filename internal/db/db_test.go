package db_test

import (
	"errors"
	"testing"

	"imagera/internal/mocks"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestConnect_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDB := mocks.NewMockDatabase(ctrl)

	mockDB.EXPECT().Connect().Return(&gorm.DB{}, nil)

	db, err := mockDB.Connect()

	assert.NoError(t, err)
	assert.NotNil(t, db)
}

func TestConnect_Failure(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDB := mocks.NewMockDatabase(ctrl)

	mockDB.EXPECT().Connect().Return(nil, errors.New("failed to connect to database"))

	db, err := mockDB.Connect()

	assert.Error(t, err)
	assert.Nil(t, db)
	assert.Equal(t, "failed to connect to database", err.Error())
}
