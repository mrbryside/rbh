//go:build unit

package service_test

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/mrbryside/rbh/domain/interview/appointment/domain/creator"
	"github.com/mrbryside/rbh/domain/interview/appointment/pkg/generated/creatormock"
	"github.com/mrbryside/rbh/domain/interview/appointment/service"
	"github.com/stretchr/testify/assert"
)

func TestGetCreatorByIdSuccess(t *testing.T) {
	// Arrange
	ctrl := gomock.NewController(t)

	cAgg, err := creator.New(1, "name", "email@a.com")
	if err != nil {
		t.Error(err.Error())
	}

	mockRepo := creatormock.NewMockRepository(ctrl)
	mockRepo.EXPECT().GetById(gomock.Any()).Return(cAgg, nil)

	// Act
	underTest := service.NewCreatorService(mockRepo)
	result, err := underTest.GetById(1)

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, cAgg, result)
}

func TestGetCreatorByIdError(t *testing.T) {
	// Arrange
	ctrl := gomock.NewController(t)

	mockRepo := creatormock.NewMockRepository(ctrl)
	mockRepo.EXPECT().GetById(gomock.Any()).Return(creator.Aggregate{}, errors.New("get by id error"))

	// Act
	underTest := service.NewCreatorService(mockRepo)
	_, err := underTest.GetById(1)

	// Assert
	assert.Error(t, err)
}
