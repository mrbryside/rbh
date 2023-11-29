package appointment

import (
	"github.com/mrbryside/rbh/domain/interview/appointment/domain/appointment"
)

func toCommentResp(agg appointment.Aggregate) []Comment {
	comments := make([]Comment, 0)
	for _, comment := range agg.Comments() {
		c := Comment{
			Id:        comment.Id,
			Message:   comment.Message,
			CreatedAt: comment.CreatedAt,
			Creator: Creator{
				Id:    comment.Creator.Id,
				Name:  comment.Creator.Name,
				Email: comment.Creator.Email,
			},
		}
		comments = append(comments, c)
	}
	return comments
}

func toHistoryResp(agg appointment.Aggregate) []History {
	histories := make([]History, 0)
	for _, history := range agg.Histories() {
		h := History{
			Id:          history.Id,
			Name:        history.Name,
			Description: history.Description,
			Status:      history.Status.Value,
		}
		histories = append(histories, h)
	}
	return histories
}

func toAppointmentResp(agg appointment.Aggregate) Response {
	return Response{
		Id:          agg.Appointment().Id,
		Name:        agg.Appointment().Name,
		Description: agg.Appointment().Description,
		Status:      agg.Appointment().Status.Value,
		Enabled:     agg.Appointment().Enabled,
		Creator: Creator{
			Id:    agg.Creator().Id,
			Name:  agg.Creator().Name,
			Email: agg.Creator().Email,
		},
		Comments:  toCommentResp(agg),
		Histories: toHistoryResp(agg),
		CreatedAt: agg.Appointment().CreatedAt,
	}
}

func toAppointmentUpdateResp(agg appointment.Aggregate) UpdateResponse {
	return UpdateResponse{
		Id:          agg.Appointment().Id,
		Name:        agg.Appointment().Name,
		Description: agg.Appointment().Description,
		Status:      agg.Appointment().Status.Value,
		Enabled:     agg.Appointment().Enabled,
		Creator: Creator{
			Id:    agg.Creator().Id,
			Name:  agg.Creator().Name,
			Email: agg.Creator().Email,
		},
		CreatedAt: agg.Appointment().CreatedAt,
	}
}

func toAppointmentPaginateResp(aggs appointment.Aggregates) Paginate {
	results := make([]GetAllResp, 0)
	for _, agg := range aggs.Aggregates {
		results = append(results, GetAllResp{
			Id:          agg.Appointment().Id,
			Name:        agg.Appointment().Name,
			Description: agg.Appointment().Description,
			Status:      agg.Appointment().Status.Value,
			Creator: Creator{
				Id:    agg.Creator().Id,
				Name:  agg.Creator().Name,
				Email: agg.Creator().Email,
			},
			CreatedAt: agg.Appointment().CreatedAt,
		})
	}
	return Paginate{
		Results: results,
		Next:    aggs.Next,
	}
}
