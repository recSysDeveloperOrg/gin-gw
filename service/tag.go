package service

import (
	"gin-gateway/rpc"
	"github.com/gin-gonic/gin"
	"strconv"
)

func CreateTag(c *gin.Context) {
	resp, _ := rpc.TagCreate(c.GetString("userID"), c.PostForm("movieID"), c.PostForm("tag"))
	c.JSONP(200, resp)
}

func QueryMovieTag(c *gin.Context) {
	resp, _ := rpc.TagQueryMovie(c.GetString("userID"), c.PostForm("movieID"))
	c.JSONP(200, resp)
}

func QueryUserTagCloud(c *gin.Context) {
	kMax, _ := strconv.ParseInt(c.PostForm("n"), 10, 64)
	resp, _ := rpc.TagQueryUserCloud(c.GetString("userID"), kMax)
	c.JSONP(200, resp)
}

func QueryTagRecords(c *gin.Context) {
	page, _ := strconv.ParseInt(c.PostForm("page"), 10, 64)
	offset, _ := strconv.ParseInt(c.PostForm("offset"), 10, 64)
	resp, _ := rpc.TagQueryRecords(c.GetString("userID"), page, offset)
	c.JSONP(200, resp)
}

func QueryMovieTopNTags(c *gin.Context) {
	kMax, _ := strconv.ParseInt(c.PostForm("n"), 10, 64)
	resp, _ := rpc.TagQueryMovieTopNTags(c.PostForm("movieID"), kMax)
	c.JSONP(200, resp)
}
