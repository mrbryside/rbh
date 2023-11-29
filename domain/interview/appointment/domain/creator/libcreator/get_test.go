//go:build unit

package libcreator_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/mrbryside/rbh/domain/interview/appointment/domain/creator/libcreator"
	"github.com/mrbryside/rbh/domain/user/domain/user"
	"github.com/mrbryside/rbh/domain/user/pkg/generated/mockgen"
	"github.com/stretchr/testify/assert"
)

func TestGetCreator(t *testing.T) {
	// Arrange
	ctrl := gomock.NewController(t)
	mock := mockgen.NewMockAuthServicer(ctrl)
	userAgg, err := user.New("test", "test@mail.com", "123456")
	if err != nil {
		t.Error(err)
	}
	userAggWithId := userAgg.SetId(1)
	mock.EXPECT().GetById(gomock.Any()).Return(userAggWithId, nil)

	// Act
	underTest := libcreator.NewRepository(mock)
	creatorAgg, err := underTest.GetById(1)

	// Assert
	assert.Nil(t, err)
	assert.Equal(t, userAggWithId.Person().Email.Value, creatorAgg.Creator().Email)
	assert.Equal(t, userAggWithId.Person().Name, creatorAgg.Creator().Name)
	assert.Equal(t, userAggWithId.Person().Id, creatorAgg.Creator().Id)
}
