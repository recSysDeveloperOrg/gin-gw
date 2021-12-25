package rpc

import (
	"context"
	"gin-gateway/idl/gen/tag"
)

func TagCreate(userID, movieID, tagContent string) (*tag.CreateTagResp, error) {
	req := &tag.CreateTagReq{
		UserId:     userID,
		MovieId:    movieID,
		TagContent: tagContent,
	}

	return TagClient.CreateTag(context.Background(), req)
}

func TagQueryMovie(userID, movieID string) (*tag.QueryMovieTagResp, error) {
	req := &tag.QueryMovieTagReq{
		UserId:  userID,
		MovieId: movieID,
	}

	return TagClient.QueryMovieTag(context.Background(), req)
}

func TagQueryUserCloud(userID string, kMax int64) (*tag.QueryUserTagCloudResp, error) {
	req := &tag.QueryUserTagCloudReq{
		UserId: userID,
		NTags:  kMax,
	}

	return TagClient.QueryUserTagCloud(context.Background(), req)
}

func TagQueryRecords(userID string, page, offset int64) (*tag.QueryTagRecordsResp, error) {
	req := &tag.QueryTagRecordsReq{
		UserId: userID,
		Page:   page,
		Offset: offset,
	}

	return TagClient.QueryTagRecords(context.Background(), req)
}

func TagQueryMovieTopNTags(movieID string, topN int64) (*tag.QueryMovieTopNTagsResp, error) {
	req := &tag.QueryMovieTopNTagsReq{
		MovieId: movieID,
		NTags:   topN,
	}

	return TagClient.QueryMovieTopNTags(context.Background(), req)
}
