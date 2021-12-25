package service

import (
	"gin-gateway/constant"
	"gin-gateway/rpc"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	resp, _ := rpc.UserLogin(c.PostForm("username"), c.PostForm("password"))
	c.JSONP(200, constant.NewWrapSuccessResponse(resp))
}

func Register(c *gin.Context) {
	resp, _ := rpc.UserRegister(c.PostForm("username"), c.PostForm("password"), c.PostForm("gender"))
	c.JSONP(200, constant.NewWrapSuccessResponse(resp))
}

func Query(c *gin.Context) {
	accessToken, ok := c.GetPostForm("accessToken")
	if ok {
		resp, _ := rpc.UserQuery(rpc.TokenTypeAccess, accessToken)
		c.JSONP(0, constant.NewWrapSuccessResponse(resp.User))
		return
	}
	refreshToken, ok := c.GetPostForm("refreshToken")
	if ok {
		resp, _ := rpc.UserQuery(rpc.TokenTypeRefresh, refreshToken)
		c.JSONP(0, constant.NewWrapSuccessResponse(resp.User))
	}
}
