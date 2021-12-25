package service

import (
	"gin-gateway/rpc"
	"github.com/gin-gonic/gin"
	"strconv"
)

func RateMovie(c *gin.Context) {
	rate, _ := strconv.ParseFloat(c.PostForm("rating"), 64)
	resp, _ := rpc.RateMovie(c.PostForm("movieID"), c.GetString("userID"), rate)
	c.JSONP(200, resp)
}

func QueryRateRecords(c *gin.Context) {
	page, _ := strconv.ParseInt(c.PostForm("page"), 10, 64)
	offset, _ := strconv.ParseInt(c.PostForm("offset"), 10, 64)
	resp, _ := rpc.RateQueryRecords(c.GetString("userID"), page, offset)
	c.JSONP(200, resp)
}

func QueryMovieRating(c *gin.Context) {
	resp, _ := rpc.RateQueryMovieRating(c.GetString("userID"), c.PostFormArray("movieIDs"))
	c.JSONP(200, resp)
}
