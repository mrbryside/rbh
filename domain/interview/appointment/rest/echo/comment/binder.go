package comment

import "github.com/mrbryside/rbh/domain/interview/appointment/domain/comment"

func toCommentResponse(agg comment.Aggregate) Response {
	return Response{
		Id:      agg.Comment().Id,
		Message: agg.Comment().Message,
		Creator: Creator{
			Id:    agg.Comment().Creator.Id,
			Name:  agg.Comment().Creator.Name,
			Email: agg.Comment().Creator.Email,
		},
	}
}

func toCommentGetAllResponse(aggs []comment.Aggregate) []Response {
	res := make([]Response, 0)
	for _, agg := range aggs {
		res = append(res, toCommentResponse(agg))
	}
	return res
}
