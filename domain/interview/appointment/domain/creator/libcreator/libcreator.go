package libcreator

import (
	"github.com/mrbryside/rbh/domain/interview/appointment/domain/creator"
	"github.com/mrbryside/rbh/domain/user/service"
)

type repository struct {
	authService service.AuthServicer
}

func NewRepository(as service.AuthServicer) creator.Repository {
	return &repository{authService: as}
}
