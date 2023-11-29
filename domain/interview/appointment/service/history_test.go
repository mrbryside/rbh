//go:build unit

package service_test

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/mrbryside/rbh/domain/interview/appointment/domain/history"
	"github.com/mrbryside/rbh/domain/interview/appointment/pkg/generated/historymock"
	"github.com/mrbryside/rbh/domain/interview/appointment/service"
	"github.com/mrbryside/rbh/domain/interview/appointment/types/mystatus"
	"github.com/stretchr/testify/assert"
)

func TestCreateHistorySuccess(t *testing.T) {
	// Arrange
	ctrl := gomock.NewController(t)
	mock := historymock.NewMockRepository(ctrl)

	dto := service.CreateHistoryDto{
		Id:          1,
		Name:        "name",
		Description: "description",
		Status:      mystatus.Done,
	}
	agg, err := history.New(dto.Id, dto.Name, dto.Description, dto.Status)
	if err != nil {
		t.Error(err.Error())
	}
	mock.EXPECT().Create(gomock.Any()).Return(agg, nil)

	// Act
	underTest := service.NewHistoryService(mock)
	result, err := underTest.Create(dto)

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, agg, result)
}

func TestCreateHistoryError(t *testing.T) {
	// Arrange
	ctrl := gomock.NewController(t)
	mock := historymock.NewMockRepository(ctrl)

	dto := service.CreateHistoryDto{
		Id:          1,
		Name:        "name",
		Description: "description",
		Status:      mystatus.Done,
	}
	mock.EXPECT().Create(gomock.Any()).Return(
		history.Aggregate{},
		errors.New("error creating history"),
	)

	// Act
	underTest := service.NewHistoryService(mock)
	_, err := underTest.Create(dto)

	// Assert
	assert.Error(t, err)
}
func TestCreateHistoryAggregateError(t *testing.T) {
	// Arrange
	ctrl := gomock.NewController(t)
	mock := historymock.NewMockRepository(ctrl)

	dto := service.CreateHistoryDto{
		Id:          1,
		Name:        "name",
		Description: "description",
		Status:      "invalid status",
	}

	// Act
	underTest := service.NewHistoryService(mock)
	_, err := underTest.Create(dto)

	// Assert
	assert.Error(t, err)
}

func TestGetHistoryAllByAppointmentId(t *testing.T) {
	// Arrange
	ctrl := gomock.NewController(t)
	mock := historymock.NewMockRepository(ctrl)

	var aggs []history.Aggregate
	agg, err := history.New(1, "name", "description", mystatus.Done)
	if err != nil {
		t.Error(err.Error())
	}
	aggs = append(aggs, agg)

	mock.EXPECT().GetAllByAppointmentId(gomock.Any()).Return(aggs, nil)
	// Act
	underTest := service.NewHistoryService(mock)
	result, err := underTest.GetAllByAppointmentId(1)

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, aggs, result)
}
