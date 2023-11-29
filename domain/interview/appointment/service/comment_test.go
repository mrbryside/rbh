//go:build unit

package service_test

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/mrbryside/rbh/domain/interview/appointment/domain/comment"
	"github.com/mrbryside/rbh/domain/interview/appointment/domain/creator"
	"github.com/mrbryside/rbh/domain/interview/appointment/pkg/generated/commentmock"
	"github.com/mrbryside/rbh/domain/interview/appointment/pkg/generated/creatormock"
	"github.com/mrbryside/rbh/domain/interview/appointment/service"
	"github.com/stretchr/testify/assert"
)

func TestCreateCommentSuccess(t *testing.T) {
	// Arrange
	ctrl := gomock.NewController(t)
	mock := commentmock.NewMockRepository(ctrl)
	cMock := creatormock.NewMockCreatorServicer(ctrl)
	agg := comment.New("Hello")
	mock.EXPECT().Create(gomock.Any()).Return(agg, nil)
	cMock.EXPECT().GetById(gomock.Any()).Return(creator.Aggregate{}, nil)

	// Act
	underTest := service.NewCommentService(mock, cMock)
	result, err := underTest.Create(service.CreateCommentDto{
		Message:       "Hello",
		AppointmentId: 1,
		CreatorId:     1,
	})

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, agg, result)
}

func TestCreateCommentError(t *testing.T) {
	// Arrange
	ctrl := gomock.NewController(t)
	mock := commentmock.NewMockRepository(ctrl)
	cMock := creatormock.NewMockCreatorServicer(ctrl)
	agg := comment.New("Hello")
	mock.EXPECT().Create(gomock.Any()).Return(agg, errors.New("create error"))

	// Act
	underTest := service.NewCommentService(mock, cMock)
	_, err := underTest.Create(service.CreateCommentDto{
		Message:       "Hello",
		AppointmentId: 1,
		CreatorId:     1,
	})

	// Assert
	assert.Error(t, err)
}

func TestCreateCommentGetCreatorError(t *testing.T) {
	// Arrange
	ctrl := gomock.NewController(t)
	mock := commentmock.NewMockRepository(ctrl)
	cMock := creatormock.NewMockCreatorServicer(ctrl)
	agg := comment.New("Hello")
	mock.EXPECT().Create(gomock.Any()).Return(agg, nil)
	cMock.EXPECT().GetById(gomock.Any()).Return(creator.Aggregate{}, errors.New("get creator error"))

	// Act
	underTest := service.NewCommentService(mock, cMock)
	_, err := underTest.Create(service.CreateCommentDto{
		Message:       "Hello",
		AppointmentId: 1,
		CreatorId:     1,
	})

	// Assert
	assert.Error(t, err)
}

func TestDeleteCommentSuccess(t *testing.T) {
	// Arrange
	ctrl := gomock.NewController(t)
	mock := commentmock.NewMockRepository(ctrl)
	cMock := creatormock.NewMockCreatorServicer(ctrl)
	creatorAgg, err := creator.New(1, "Hello", "mail@mail.com")
	if err != nil {
		t.Error(err.Error())
	}
	commentAgg := comment.New("Hello").
		SetCreator(creatorAgg.Creator())

	mock.EXPECT().GetById(gomock.Any()).Return(commentAgg, nil)
	mock.EXPECT().DeleteById(gomock.Any()).Return(nil)

	// Act
	underTest := service.NewCommentService(mock, cMock)
	err = underTest.DeleteById(1, 1)

	// Assert
	assert.NoError(t, err)
}
func TestDeleteCommentErrorCreatorNotMatch(t *testing.T) {
	// Arrange
	ctrl := gomock.NewController(t)
	mock := commentmock.NewMockRepository(ctrl)
	cMock := creatormock.NewMockCreatorServicer(ctrl)
	creatorAgg, err := creator.New(1, "Hello", "mail@mail.com")
	if err != nil {
		t.Error(err.Error())
	}
	commentAgg := comment.New("Hello").
		SetCreator(creatorAgg.Creator())

	mock.EXPECT().GetById(gomock.Any()).Return(commentAgg, nil)

	// Act
	underTest := service.NewCommentService(mock, cMock)
	err = underTest.DeleteById(1, 2)

	// Assert
	assert.Error(t, err)
}

func TestDeleteCommentErrorGetCreatorError(t *testing.T) {
	// Arrange
	ctrl := gomock.NewController(t)
	mock := commentmock.NewMockRepository(ctrl)
	cMock := creatormock.NewMockCreatorServicer(ctrl)
	creatorAgg, err := creator.New(1, "Hello", "mail@mail.com")
	if err != nil {
		t.Error(err.Error())
	}
	commentAgg := comment.New("Hello").
		SetCreator(creatorAgg.Creator())

	mock.EXPECT().GetById(gomock.Any()).Return(commentAgg, errors.New("get creator error"))

	// Act
	underTest := service.NewCommentService(mock, cMock)
	err = underTest.DeleteById(1, 1)

	// Assert
	assert.Error(t, err)
}

func TestUpdateCommentSuccess(t *testing.T) {
	// Arrange
	ctrl := gomock.NewController(t)
	mock := commentmock.NewMockRepository(ctrl)
	cMock := creatormock.NewMockCreatorServicer(ctrl)
	creatorAgg, err := creator.New(1, "Hello", "mail@mail.com")
	if err != nil {
		t.Error(err.Error())
	}
	commentAgg := comment.New("edited hello").
		SetCreator(creatorAgg.Creator())

	mock.EXPECT().GetById(gomock.Any()).Return(commentAgg, nil)
	mock.EXPECT().UpdateById(gomock.Any()).Return(commentAgg, nil)
	cMock.EXPECT().GetById(gomock.Any()).Return(creatorAgg, nil)

	// Act
	underTest := service.NewCommentService(mock, cMock)
	result, err := underTest.UpdateById(service.UpdateCommentDto{
		Id:        1,
		CreatorId: 1,
		Message:   "Hello",
	})

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, commentAgg, result)
}

func TestUpdateCommentCreatorError(t *testing.T) {
	// Arrange
	ctrl := gomock.NewController(t)
	mock := commentmock.NewMockRepository(ctrl)
	cMock := creatormock.NewMockCreatorServicer(ctrl)
	creatorAgg, err := creator.New(1, "Hello", "mail@mail.com")
	if err != nil {
		t.Error(err.Error())
	}
	commentAgg := comment.New("edited hello").
		SetCreator(creatorAgg.Creator())

	mock.EXPECT().GetById(gomock.Any()).Return(commentAgg, nil)
	mock.EXPECT().UpdateById(gomock.Any()).Return(commentAgg, errors.New("update error"))

	// Act
	underTest := service.NewCommentService(mock, cMock)
	_, err = underTest.UpdateById(service.UpdateCommentDto{
		Id:        1,
		CreatorId: 1,
		Message:   "edited hello",
	})

	// Assert
	assert.Error(t, err)
}
func TestUpdateCommentCreatorNotMatchError(t *testing.T) {
	// Arrange
	ctrl := gomock.NewController(t)
	mock := commentmock.NewMockRepository(ctrl)
	cMock := creatormock.NewMockCreatorServicer(ctrl)
	creatorAgg, err := creator.New(1, "Hello", "mail@mail.com")
	if err != nil {
		t.Error(err.Error())
	}
	commentAgg := comment.New("edited hello").
		SetCreator(creatorAgg.Creator())

	mock.EXPECT().GetById(gomock.Any()).Return(commentAgg, nil)

	// Act
	underTest := service.NewCommentService(mock, cMock)
	_, err = underTest.UpdateById(service.UpdateCommentDto{
		Id:        1,
		CreatorId: 9,
		Message:   "Hello",
	})

	// Assert
	assert.Error(t, err)
}

func TestGetAllByAppointmentIdSuccess(t *testing.T) {
	// Arrange
	ctrl := gomock.NewController(t)
	mock := commentmock.NewMockRepository(ctrl)
	cMock := creatormock.NewMockCreatorServicer(ctrl)
	var aggs []comment.Aggregate
	agg := comment.New("Hello")
	agg.SetCommentId(1)
	aggs = append(aggs, agg)
	mock.EXPECT().GetAllByAppointmentId(gomock.Any()).Return(aggs, nil)
	cMock.EXPECT().GetById(gomock.Any()).Return(creator.Aggregate{}, nil)

	// Act
	underTest := service.NewCommentService(mock, cMock)
	result, err := underTest.GetAllByAppointmentId(1)

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, aggs, result)
}

func TestGetAllByAppointmentIdError(t *testing.T) {
	// Arrange
	ctrl := gomock.NewController(t)
	mock := commentmock.NewMockRepository(ctrl)
	cMock := creatormock.NewMockCreatorServicer(ctrl)
	var aggs []comment.Aggregate
	agg := comment.New("Hello")
	agg.SetCommentId(1)
	aggs = append(aggs, agg)
	mock.EXPECT().GetAllByAppointmentId(gomock.Any()).Return(aggs, errors.New("get all error"))

	// Act
	underTest := service.NewCommentService(mock, cMock)
	_, err := underTest.GetAllByAppointmentId(1)

	// Assert
	assert.Error(t, err)
}

func TestGetAllByAppointmentIdCreatorError(t *testing.T) {
	// Arrange
	ctrl := gomock.NewController(t)
	mock := commentmock.NewMockRepository(ctrl)
	cMock := creatormock.NewMockCreatorServicer(ctrl)
	var aggs []comment.Aggregate
	agg := comment.New("Hello")
	agg.SetCommentId(1)
	aggs = append(aggs, agg)
	mock.EXPECT().GetAllByAppointmentId(gomock.Any()).Return(aggs, nil)
	cMock.EXPECT().GetById(gomock.Any()).Return(creator.Aggregate{}, errors.New("get creator error"))

	// Act
	underTest := service.NewCommentService(mock, cMock)
	_, err := underTest.GetAllByAppointmentId(1)

	// Assert
	assert.Error(t, err)
}
