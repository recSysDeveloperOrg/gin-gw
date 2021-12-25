package service

import (
	"gin-gateway/rpc"
	"github.com/gin-gonic/gin"
	"strconv"
)

func RecommendMovies(c *gin.Context) {
	page, _ := strconv.ParseInt(c.PostForm("page"), 10, 64)
	offset, _ := strconv.ParseInt(c.PostForm("offset"), 10, 64)
	resp, _ := rpc.MovieRecommend(c.GetString("userID"), page, offset)
	c.JSONP(200, resp)
}

func GetMovieDetail(c *gin.Context) {
	resp, _ := rpc.MovieGetDetail(c.GetString("userID"), c.PostForm("movieID"))
	c.JSONP(200, resp)
}

func SearchMovies(c *gin.Context) {
	page, _ := strconv.ParseInt(c.PostForm("page"), 10, 64)
	offset, _ := strconv.ParseInt(c.PostForm("offset"), 10, 64)
	resp, _ := rpc.MovieSearch(c.PostForm("keyword"), page, offset)
	c.JSONP(200, resp)
}

func RecommendFeedback(c *gin.Context) {
	resp, _ := rpc.MovieRecommendFeedback(c.PostForm("ft"), c.PostForm("sourceID"), c.GetString("userID"))
	c.JSONP(200, resp)
}
