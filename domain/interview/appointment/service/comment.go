package service

import (
	"errors"

	"github.com/mrbryside/rbh/domain/interview/appointment/domain/comment"
)

//go:generate mockgen -source=comment.go -destination=../pkg/generated/commentmock/service.go -package=commentmock
type CommentServicer interface {
	Create(CreateCommentDto) (comment.Aggregate, error)
	UpdateById(UpdateCommentDto) (comment.Aggregate, error)
	GetAllByAppointmentId(uint) ([]comment.Aggregate, error)
	DeleteById(commentId uint, creatorId uint) error
}

type commentService struct {
	commentDomain  comment.Repository
	creatorService CreatorServicer
}

func NewCommentService(cr comment.Repository, cs CreatorServicer) CommentServicer {
	return commentService{
		commentDomain:  cr,
		creatorService: cs,
	}
}

// CreateCommentDto is a data transfer object for creating a comment
type CreateCommentDto struct {
	Message       string
	AppointmentId uint
	CreatorId     uint
}

func (cs commentService) Create(dto CreateCommentDto) (comment.Aggregate, error) {
	agg, err := cs.commentDomain.Create(
		comment.New(dto.Message).
			SetAppointmentId(dto.AppointmentId).
			SetCreatorId(dto.CreatorId),
	)
	if err != nil {
		return agg, err
	}
	return cs.applyCreator(agg)
}

// UpdateCommentDto is a data transfer object for updating a comment
type UpdateCommentDto struct {
	Id        uint
	CreatorId uint
	Message   string
}

func (cs commentService) UpdateById(dto UpdateCommentDto) (comment.Aggregate, error) {
	mismatch, err := cs.isCreatorIdNotMatch(dto.Id, dto.CreatorId)
	if mismatch {
		return comment.Aggregate{}, err
	}
	agg, err := cs.commentDomain.UpdateById(
		comment.New(dto.Message).
			SetCommentId(dto.Id),
	)
	if err != nil {
		return agg, err
	}
	return cs.applyCreator(agg)
}

func (cs commentService) GetAllByAppointmentId(appointmentId uint) ([]comment.Aggregate, error) {
	aggs, err := cs.commentDomain.GetAllByAppointmentId(appointmentId)
	if err != nil {
		return aggs, err
	}
	return cs.applyCreators(aggs)
}

func (cs commentService) DeleteById(id uint, creatorId uint) error {
	mismatch, err := cs.isCreatorIdNotMatch(id, creatorId)
	if mismatch {
		return err
	}
	return cs.commentDomain.DeleteById(id)
}

func (cs commentService) applyCreator(agg comment.Aggregate) (comment.Aggregate, error) {
	creatorAgg, err := cs.creatorService.GetById(agg.Comment().Creator.Id)
	if err != nil {
		return agg, err
	}
	return agg.SetCreator(creatorAgg.Creator()), nil
}

func (cs commentService) applyCreators(aggs []comment.Aggregate) ([]comment.Aggregate, error) {
	results := make([]comment.Aggregate, 0)
	for _, agg := range aggs {
		aggWithCreator, err := cs.applyCreator(agg)
		if err != nil {
			return aggs, err
		}
		results = append(results, aggWithCreator)
	}
	return results, nil
}

func (cs commentService) isCreatorIdNotMatch(commentId uint, creatorId uint) (bool, error) {
	getAgg, err := cs.commentDomain.GetById(commentId)
	if err != nil {
		return true, errors.New("not found comment")
	}
	if getAgg.Comment().Creator.Id != creatorId {
		return true, errors.New("creator mismatch")
	}
	return false, nil
}
