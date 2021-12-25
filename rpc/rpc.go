package rpc

import (
	"flag"
	"fmt"
	"gin-gateway/idl/gen/movie"
	"gin-gateway/idl/gen/rating"
	"gin-gateway/idl/gen/tag"
	"gin-gateway/idl/gen/user"
	"google.golang.org/grpc"
)

var (
	MovieClient  movie.MovieServiceClient
	UserClient   user.UserServiceClient
	RatingClient rating.RatingServiceClient
	TagClient    tag.TagServiceClient
)

func Init() {
	var ip string
	var pMovie int
	var pUser int
	var pRating int
	var pTag int
	flag.StringVar(&ip, "ip", "127.0.0.1", "service ip")
	flag.IntVar(&pMovie, "pMovie", 50001, "movie service port")
	flag.IntVar(&pUser, "pUser", 50002, "user service port")
	flag.IntVar(&pRating, "pRating", 50003, "rating service port")
	flag.IntVar(&pTag, "pTag", 50004, "tag service port")
	flag.Parse()

	MovieClient = movie.NewMovieServiceClient(mustDial(ip, pMovie))
	UserClient = user.NewUserServiceClient(mustDial(ip, pUser))
	RatingClient = rating.NewRatingServiceClient(mustDial(ip, pRating))
	TagClient = tag.NewTagServiceClient(mustDial(ip, pTag))
}

func mustDial(ip string, port int) *grpc.ClientConn {
	conn, err := dial(ip, port)
	if err != nil {
		panic(err)
	}

	return conn
}

func dial(ip string, port int) (*grpc.ClientConn, error) {
	return grpc.Dial(fmt.Sprintf("%s:%d", ip, port), grpc.WithInsecure())
}
