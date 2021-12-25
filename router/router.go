package router

import (
	"gin-gateway/service"
	"github.com/gin-gonic/gin"
)

func Init() {
	r := gin.Default()
	v1 := r.Group("/v1")
	{
		movie := v1.Group("/movie")
		{
			movie.POST("/recommend", service.RecommendMovies)
			movie.GET("/:id", service.GetMovieDetail)
			movie.POST("/search", service.SearchMovies)
			movie.POST("/recommend-feedback", service.RecommendFeedback)
		}
		rate := v1.Group("/rate")
		{
			rate.POST("/", service.RateMovie)
			rate.POST("/history", service.QueryRateRecords)
			rate.POST("/movie", service.QueryMovieRating)
		}
		tag := v1.Group("/tag")
		{
			tag.POST("/", service.CreateTag)
			tag.POST("/user-movie-tags", service.QueryMovieTag)
			tag.POST("/user-tag-cloud", service.QueryUserTagCloud)
			tag.POST("/history", service.QueryTagRecords)
			tag.POST("/movie-top-k", service.QueryMovieTopNTags)
		}
		user := v1.Group("/user")
		{
			user.POST("/register", service.Register)
			user.POST("/login", service.Login)
			user.POST("/", service.Query)
		}
	}
	if err := r.Run(); err != nil {
		panic(err)
	}
}
