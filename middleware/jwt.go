package middleware

import (
	"gin-gateway/rpc"
	"github.com/gin-gonic/gin"
)

func JwtAuth(c *gin.Context) {
	accessToken, ok := c.GetPostForm("accessToken")
	if ok {
		resp, err := rpc.UserQuery(rpc.TokenTypeAccess, accessToken)
		if err != nil {
			c.JSONP(500, err)
			return
		}
		if resp.User != nil {
			c.Set("userID", resp.User.Id)
		}
	}

	c.Next()
}
