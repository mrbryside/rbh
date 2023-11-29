//go:build unit

package service_test

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/mrbryside/rbh/domain/interview/appointment/domain/appointment"
	"github.com/mrbryside/rbh/domain/interview/appointment/domain/comment"
	"github.com/mrbryside/rbh/domain/interview/appointment/domain/creator"
	"github.com/mrbryside/rbh/domain/interview/appointment/domain/history"
	"github.com/mrbryside/rbh/domain/interview/appointment/pkg/generated/appointmentmock"
	"github.com/mrbryside/rbh/domain/interview/appointment/pkg/generated/commentmock"
	"github.com/mrbryside/rbh/domain/interview/appointment/pkg/generated/creatormock"
	"github.com/mrbryside/rbh/domain/interview/appointment/pkg/generated/historymock"
	"github.com/mrbryside/rbh/domain/interview/appointment/service"
	"github.com/mrbryside/rbh/domain/interview/appointment/types/mystatus"
	"github.com/stretchr/testify/assert"
)

func TestCreateAppointmentSuccess(t *testing.T) {
	// Arrange
	ctrl := gomock.NewController(t)
	mock := appointmentmock.NewMockRepository(ctrl)
	hMock := historymock.NewMockHistoryServicer(ctrl)
	cMock := creatormock.NewMockCreatorServicer(ctrl)
	csMock := commentmock.NewMockCommentServicer(ctrl)

	dto := service.CreateAppointmentDto{
		Name:        "test name",
		Description: "test description",
		CreatorId:   1,
	}
	agg := appointment.New(dto.Name, dto.Description)
	mock.EXPECT().Create(gomock.Any()).Return(agg.SetCreatorId(dto.CreatorId), nil)
	hMock.EXPECT().Create(gomock.Any()).Return(history.Aggregate{}, nil)
	cMock.EXPECT().GetById(gomock.Any()).Return(creator.Aggregate{}, nil)

	// Act
	underTest := service.NewAppointmentService(mock, hMock, cMock, csMock)
	result, err := underTest.Create(dto)

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, result.Appointment().Name, dto.Name)
	assert.Equal(t, result.Appointment().Description, dto.Description)

}

func TestCreateAppointmentSuccessButHistorySaveError(t *testing.T) {
	// Arrange
	ctrl := gomock.NewController(t)
	mock := appointmentmock.NewMockRepository(ctrl)
	hMock := historymock.NewMockHistoryServicer(ctrl)
	cMock := creatormock.NewMockCreatorServicer(ctrl)
	csMock := commentmock.NewMockCommentServicer(ctrl)

	dto := service.CreateAppointmentDto{
		Name:        "test name",
		Description: "test description",
		CreatorId:   1,
	}
	agg := appointment.New(dto.Name, dto.Description)
	mock.EXPECT().Create(gomock.Any()).Return(agg.SetCreatorId(dto.CreatorId), nil)
	hMock.EXPECT().Create(gomock.Any()).Return(history.Aggregate{}, errors.New("error saving history"))
	cMock.EXPECT().GetById(gomock.Any()).Return(creator.Aggregate{}, nil)

	// Act
	underTest := service.NewAppointmentService(mock, hMock, cMock, csMock)
	result, err := underTest.Create(dto)

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, result.Appointment().Name, dto.Name)
	assert.Equal(t, result.Appointment().Description, dto.Description)
}

func TestCreateAppointmentError(t *testing.T) {
	// Arrange
	ctrl := gomock.NewController(t)
	mock := appointmentmock.NewMockRepository(ctrl)
	hMock := historymock.NewMockHistoryServicer(ctrl)
	cMock := creatormock.NewMockCreatorServicer(ctrl)
	csMock := commentmock.NewMockCommentServicer(ctrl)

	mock.EXPECT().Create(gomock.Any()).Return(appointment.Aggregate{}, errors.New("error creating appointment"))

	// Act
	underTest := service.NewAppointmentService(mock, hMock, cMock, csMock)
	dto := service.CreateAppointmentDto{
		Name:        "test name",
		Description: "test description",
		CreatorId:   1,
	}
	_, err := underTest.Create(dto)

	// Assert
	assert.Error(t, err)
}

func TestGetAllAppointmentSuccess(t *testing.T) {
	// Arrange
	ctrl := gomock.NewController(t)
	mock := appointmentmock.NewMockRepository(ctrl)
	hMock := historymock.NewMockHistoryServicer(ctrl)
	cMock := creatormock.NewMockCreatorServicer(ctrl)
	csMock := commentmock.NewMockCommentServicer(ctrl)

	agg := appointment.New("test name", "test description")
	aggs := appointment.Aggregates{}
	aggs.Aggregates = append(aggs.Aggregates, agg)
	creatorAgg, err := creator.New(1, "test name", "test email")
	if err != nil {
		t.Error(err.Error())
	}
	mock.EXPECT().GetAll(gomock.Any(), gomock.Any()).Return(aggs, nil)
	cMock.EXPECT().GetById(gomock.Any()).Return(creatorAgg, nil)

	// Act
	underTest := service.NewAppointmentService(mock, hMock, cMock, csMock)
	dto := service.GetAllAppointmentDto{
		Page:     1,
		PageSize: 1,
	}
	result, err := underTest.GetAll(dto)

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, agg.Appointment().Id, result.Aggregates[0].Appointment().Id)
	assert.Equal(t, creatorAgg.Creator().Id, result.Aggregates[0].Creator().Id)
	assert.Equal(t, creatorAgg.Creator().Name, result.Aggregates[0].Creator().Name)
	assert.Equal(t, creatorAgg.Creator().Email, result.Aggregates[0].Creator().Email)
}

func TestGetAllAppointmentError(t *testing.T) {
	// Arrange
	ctrl := gomock.NewController(t)
	mock := appointmentmock.NewMockRepository(ctrl)
	hMock := historymock.NewMockHistoryServicer(ctrl)
	cMock := creatormock.NewMockCreatorServicer(ctrl)
	csMock := commentmock.NewMockCommentServicer(ctrl)

	agg := appointment.New("test name", "test description")
	aggs := appointment.Aggregates{}
	aggs.Aggregates = append(aggs.Aggregates, agg)
	mock.EXPECT().GetAll(gomock.Any(), gomock.Any()).Return(aggs, nil)

	cMock.EXPECT().GetById(gomock.Any()).Return(
		creator.Aggregate{}, errors.New("error getting creator by id"),
	)
	// Act
	underTest := service.NewAppointmentService(mock, hMock, cMock, csMock)
	dto := service.GetAllAppointmentDto{
		Page:     1,
		PageSize: 1,
	}
	_, err := underTest.GetAll(dto)

	// Assert
	assert.Error(t, err)
}

func TestGetAllAppointmentGetCreatorByIdError(t *testing.T) {
	// Arrange
	ctrl := gomock.NewController(t)
	mock := appointmentmock.NewMockRepository(ctrl)
	hMock := historymock.NewMockHistoryServicer(ctrl)
	cMock := creatormock.NewMockCreatorServicer(ctrl)
	csMock := commentmock.NewMockCommentServicer(ctrl)

	mock.EXPECT().GetAll(
		gomock.Any(),
		gomock.Any(),
	).Return(
		appointment.Aggregates{},
		errors.New("error getting all appointments"),
	)

	// Act
	underTest := service.NewAppointmentService(mock, hMock, cMock, csMock)
	dto := service.GetAllAppointmentDto{
		Page:     1,
		PageSize: 1,
	}
	_, err := underTest.GetAll(dto)

	// Assert
	assert.Error(t, err)
}

func TestGetAppointmentByIdSuccess(t *testing.T) {
	// Arrange
	ctrl := gomock.NewController(t)
	mock := appointmentmock.NewMockRepository(ctrl)
	hMock := historymock.NewMockHistoryServicer(ctrl)
	cMock := creatormock.NewMockCreatorServicer(ctrl)
	csMock := commentmock.NewMockCommentServicer(ctrl)

	agg := appointment.New("test name", "test description")
	creatorAgg, err := creator.New(1, "test name", "test email")
	if err != nil {
		t.Error(err.Error())
	}
	commentOne := comment.New("test comment one")
	commentTwo := comment.New("test comment two")
	comments := []comment.Aggregate{commentOne, commentTwo}

	historyOne, err := history.New(1, "test name", "test description", mystatus.Done)
	if err != nil {
		t.Error(err.Error())
	}
	historyTwo, err := history.New(2, "test name", "test description", mystatus.Done)
	if err != nil {
		t.Error(err.Error())
	}
	histories := []history.Aggregate{historyOne, historyTwo}

	mock.EXPECT().GetById(gomock.Any()).Return(agg, nil)
	cMock.EXPECT().GetById(gomock.Any()).Return(creatorAgg, nil)
	cMock.EXPECT().GetById(gomock.Any()).Return(creatorAgg, nil)
	cMock.EXPECT().GetById(gomock.Any()).Return(creatorAgg, nil)
	csMock.EXPECT().GetAllByAppointmentId(gomock.Any()).Return(comments, nil)
	hMock.EXPECT().GetAllByAppointmentId(gomock.Any()).Return(histories, nil)

	// Act
	underTest := service.NewAppointmentService(mock, hMock, cMock, csMock)
	result, err := underTest.GetById(1)

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, agg.Appointment().Id, result.Appointment().Id)
	assert.Equal(t, creatorAgg.Creator().Id, result.Creator().Id)
	assert.Equal(t, creatorAgg.Creator().Name, result.Creator().Name)
	assert.Equal(t, creatorAgg.Creator().Email, result.Creator().Email)
	assert.Equal(t, 2, len(result.Comments()))
	assert.Equal(t, 2, len(result.Histories()))
}

func TestGetAppointmentByIdFailedApplyHistories(t *testing.T) {
	// Arrange
	ctrl := gomock.NewController(t)
	mock := appointmentmock.NewMockRepository(ctrl)
	hMock := historymock.NewMockHistoryServicer(ctrl)
	cMock := creatormock.NewMockCreatorServicer(ctrl)
	csMock := commentmock.NewMockCommentServicer(ctrl)

	agg := appointment.New("test name", "test description")
	creatorAgg, err := creator.New(1, "test name", "test email")
	if err != nil {
		t.Error(err.Error())
	}
	comments := []comment.Aggregate{}
	histories := []history.Aggregate{}

	mock.EXPECT().GetById(gomock.Any()).Return(agg, nil)
	cMock.EXPECT().GetById(gomock.Any()).Return(creatorAgg, nil)
	csMock.EXPECT().GetAllByAppointmentId(gomock.Any()).Return(comments, nil)
	hMock.EXPECT().GetAllByAppointmentId(gomock.Any()).Return(histories, errors.New("error getting histories by appointment id"))

	// Act
	underTest := service.NewAppointmentService(mock, hMock, cMock, csMock)
	_, err = underTest.GetById(1)

	// Assert
	assert.Error(t, err)
}

func TestGetAppointmentByIdFailedApplyCreatorForComment(t *testing.T) {
	// Arrange
	ctrl := gomock.NewController(t)
	mock := appointmentmock.NewMockRepository(ctrl)
	hMock := historymock.NewMockHistoryServicer(ctrl)
	cMock := creatormock.NewMockCreatorServicer(ctrl)
	csMock := commentmock.NewMockCommentServicer(ctrl)

	agg := appointment.New("test name", "test description")
	creatorAgg, err := creator.New(1, "test name", "test email")
	if err != nil {
		t.Error(err.Error())
	}
	commentOne := comment.New("test comment one")
	commentTwo := comment.New("test comment two")
	comments := []comment.Aggregate{commentOne, commentTwo}
	mock.EXPECT().GetById(gomock.Any()).Return(agg, nil)
	cMock.EXPECT().GetById(gomock.Any()).Return(creatorAgg, nil)
	csMock.EXPECT().GetAllByAppointmentId(gomock.Any()).Return(comments, nil)
	cMock.EXPECT().GetById(gomock.Any()).Return(creatorAgg, errors.New("error getting creator by id"))

	// Act
	underTest := service.NewAppointmentService(mock, hMock, cMock, csMock)
	_, err = underTest.GetById(1)

	// Assert
	assert.Error(t, err)
}

func TestGetAppointmentByIdErrorNotFoundAppointment(t *testing.T) {
	// Arrange
	ctrl := gomock.NewController(t)
	mock := appointmentmock.NewMockRepository(ctrl)
	hMock := historymock.NewMockHistoryServicer(ctrl)
	cMock := creatormock.NewMockCreatorServicer(ctrl)
	csMock := commentmock.NewMockCommentServicer(ctrl)

	mock.EXPECT().GetById(gomock.Any()).Return(
		appointment.Aggregate{},
		errors.New("error getting appointment by id"),
	)

	// Act
	underTest := service.NewAppointmentService(mock, hMock, cMock, csMock)
	_, err := underTest.GetById(1)

	// Assert
	assert.Error(t, err)
}

func TestGetAppointmentByIdErrorApplyCreator(t *testing.T) {
	// Arrange
	ctrl := gomock.NewController(t)
	mock := appointmentmock.NewMockRepository(ctrl)
	hMock := historymock.NewMockHistoryServicer(ctrl)
	cMock := creatormock.NewMockCreatorServicer(ctrl)
	csMock := commentmock.NewMockCommentServicer(ctrl)

	mock.EXPECT().GetById(gomock.Any()).Return(
		appointment.Aggregate{},
		nil,
	)
	cMock.EXPECT().GetById(gomock.Any()).Return(
		creator.Aggregate{},
		errors.New("error getting creator by id"),
	)

	// Act
	underTest := service.NewAppointmentService(mock, hMock, cMock, csMock)
	_, err := underTest.GetById(1)

	// Assert
	assert.Error(t, err)
}

func TestGetAppointmentByIdErrorAddComments(t *testing.T) {
	// Arrange
	ctrl := gomock.NewController(t)
	mock := appointmentmock.NewMockRepository(ctrl)
	hMock := historymock.NewMockHistoryServicer(ctrl)
	cMock := creatormock.NewMockCreatorServicer(ctrl)
	csMock := commentmock.NewMockCommentServicer(ctrl)

	mock.EXPECT().GetById(gomock.Any()).Return(
		appointment.Aggregate{},
		nil,
	)
	cMock.EXPECT().GetById(gomock.Any()).Return(creator.Aggregate{}, nil)
	csMock.EXPECT().GetAllByAppointmentId(gomock.Any()).Return(
		[]comment.Aggregate{},
		errors.New("error getting comments by appointment id"),
	)

	// Act
	underTest := service.NewAppointmentService(mock, hMock, cMock, csMock)
	_, err := underTest.GetById(1)

	// Assert
	assert.Error(t, err)
}

func TestUpdateAppointmentByIdSuccess(t *testing.T) {
	// Arrange
	ctrl := gomock.NewController(t)
	mock := appointmentmock.NewMockRepository(ctrl)
	hMock := historymock.NewMockHistoryServicer(ctrl)
	cMock := creatormock.NewMockCreatorServicer(ctrl)
	csMock := commentmock.NewMockCommentServicer(ctrl)

	dto := service.UpdateAppointmentDto{
		Id:          1,
		Name:        "new name",
		Description: "new description",
		Enabled:     false,
		Status:      mystatus.Done,
	}
	agg := appointment.New(dto.Name, dto.Description)
	aggWithStatus, err := agg.SetStatus(dto.Status)
	if err != nil {
		t.Error(err.Error())
	}
	aggWithEnabled := aggWithStatus.SetEnabled(dto.Enabled)
	mock.EXPECT().UpdateById(gomock.Any()).Return(aggWithEnabled, nil)
	hMock.EXPECT().Create(gomock.Any()).Return(history.Aggregate{}, nil)
	cMock.EXPECT().GetById(gomock.Any()).Return(creator.Aggregate{}, nil)

	// Act
	underTest := service.NewAppointmentService(mock, hMock, cMock, csMock)
	result, err := underTest.UpdateById(dto)

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, result.Appointment().Name, dto.Name)
	assert.Equal(t, result.Appointment().Description, dto.Description)
	assert.Equal(t, result.Appointment().Enabled, dto.Enabled)
	assert.Equal(t, result.Appointment().Status.Value, dto.Status)
}

func TestUpdateAppointmentByIdApplyCreatorError(t *testing.T) {
	// Arrange
	ctrl := gomock.NewController(t)
	mock := appointmentmock.NewMockRepository(ctrl)
	hMock := historymock.NewMockHistoryServicer(ctrl)
	cMock := creatormock.NewMockCreatorServicer(ctrl)
	csMock := commentmock.NewMockCommentServicer(ctrl)

	dto := service.UpdateAppointmentDto{
		Id:          1,
		Name:        "new name",
		Description: "new description",
		Enabled:     false,
		Status:      mystatus.Done,
	}
	agg := appointment.New(dto.Name, dto.Description)
	aggWithStatus, err := agg.SetStatus(dto.Status)
	if err != nil {
		t.Error(err.Error())
	}
	aggWithEnabled := aggWithStatus.SetEnabled(dto.Enabled)
	mock.EXPECT().UpdateById(gomock.Any()).Return(aggWithEnabled, nil)
	hMock.EXPECT().Create(gomock.Any()).Return(history.Aggregate{}, nil)
	cMock.EXPECT().GetById(gomock.Any()).Return(creator.Aggregate{}, errors.New("error getting creator by id"))

	// Act
	underTest := service.NewAppointmentService(mock, hMock, cMock, csMock)
	_, err = underTest.UpdateById(dto)

	// Assert
	assert.Error(t, err)
}

func TestUpdateAppointmentByIdError(t *testing.T) {
	// Arrange
	ctrl := gomock.NewController(t)
	mock := appointmentmock.NewMockRepository(ctrl)
	hMock := historymock.NewMockHistoryServicer(ctrl)
	cMock := creatormock.NewMockCreatorServicer(ctrl)
	csMock := commentmock.NewMockCommentServicer(ctrl)

	dto := service.UpdateAppointmentDto{
		Id:          1,
		Name:        "new name",
		Description: "new description",
		Status:      mystatus.Done,
	}
	mock.EXPECT().UpdateById(gomock.Any()).Return(
		appointment.Aggregate{},
		errors.New("error updating appointment"),
	)

	// Act
	underTest := service.NewAppointmentService(mock, hMock, cMock, csMock)
	_, err := underTest.UpdateById(dto)

	// Assert
	assert.Error(t, err)
}

func TestUpdateAppointmentByIdAggregateError(t *testing.T) {
	// Arrange
	ctrl := gomock.NewController(t)
	mock := appointmentmock.NewMockRepository(ctrl)
	hMock := historymock.NewMockHistoryServicer(ctrl)
	cMock := creatormock.NewMockCreatorServicer(ctrl)
	csMock := commentmock.NewMockCommentServicer(ctrl)

	dto := service.UpdateAppointmentDto{
		Id:          1,
		Name:        "new name",
		Description: "new description",
		Status:      "invalid status",
	}

	// Act
	underTest := service.NewAppointmentService(mock, hMock, cMock, csMock)
	_, err := underTest.UpdateById(dto)

	// Assert
	assert.Error(t, err)
}
