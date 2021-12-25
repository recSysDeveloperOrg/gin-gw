package rpc

import (
	"context"
	"gin-gateway/idl/gen/rating"
)

func RateMovie(movieID, userID string, rate float64) (*rating.RateResp, error) {
	req := &rating.RateReq{
		MovieId: movieID,
		UserId:  userID,
		Rating:  rate,
	}

	return RatingClient.RateMovie(context.Background(), req)
}

func RateQueryRecords(userID string, page, offset int64) (*rating.QueryRateRecordsResp, error) {
	req := &rating.QueryRateRecordsReq{
		UserId: userID,
		Page:   page,
		Offset: offset,
	}

	return RatingClient.QueryRateRecords(context.Background(), req)
}

func RateQueryMovieRating(userID string, movieIds []string) (*rating.QueryMovieRatingResp, error) {
	req := &rating.QueryMovieRatingReq{
		UserId:      userID,
		MovieIdList: movieIds,
	}

	return RatingClient.BatchQueryMovieRating(context.Background(), req)
}
