package service

import (
	"github.com/mrbryside/rbh/domain/interview/appointment/domain/creator"
)

//go:generate mockgen -source=creator.go -destination=../pkg/generated/creatormock/service.go -package=creatormock
type CreatorServicer interface {
	GetById(id uint) (creator.Aggregate, error)
}

type creatorService struct {
	CreatorDomain creator.Repository
}

func NewCreatorService(cr creator.Repository) CreatorServicer {
	return creatorService{cr}

}

func (s creatorService) GetById(id uint) (creator.Aggregate, error) {
	return s.CreatorDomain.GetById(id)
}
