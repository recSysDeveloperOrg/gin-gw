package rpc

import (
	"context"
	"gin-gateway/idl/gen/movie"
)

var feedBackTypeStr2FeedbackType = map[string]movie.FeedbackType{
	FeedbackTypeMovie:    movie.FeedbackType_FEEDBACK_TYPE_MOVIE,
	FeedbackTypeCategory: movie.FeedbackType_FEEDBACK_TYPE_CATEGORY,
}

const (
	FeedbackTypeMovie    = "movie"
	FeedbackTypeCategory = "tag"
)

func MovieRecommend(userID string, page, offset int64) (*movie.RecommendResp, error) {
	req := &movie.RecommendReq{
		UserId: userID,
		Page:   page,
		Offset: offset,
	}

	return MovieClient.RecommendMovies(context.Background(), req)
}

func MovieGetDetail(userID, movieID string) (*movie.MovieDetailResp, error) {
	req := &movie.MovieDetailReq{
		UserId: userID,
		Id:     movieID,
	}

	return MovieClient.GetMovieDetail(context.Background(), req)
}

func MovieSearch(keyword string, page, offset int64) (*movie.SearchResp, error) {
	req := &movie.SearchReq{
		Keyword: keyword,
		Page:    page,
		Offset:  offset,
	}

	return MovieClient.SearchMovies(context.Background(), req)
}

func MovieRecommendFeedback(feedbackType, sourceID, userID string) (*movie.FeedbackResp, error) {
	req := &movie.FeedbackReq{
		UserId:   userID,
		SourceId: sourceID,
		FbType:   feedBackTypeStr2FeedbackType[feedbackType],
	}

	return MovieClient.RecommendFeedback(context.Background(), req)
}
