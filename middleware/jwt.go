package middleware

import (
	"gin-gateway/rpc"
	"github.com/gin-gonic/gin"
)

func JwtAuth(c *gin.Context) {
	accessToken, ok := c.GetPostForm("accessToken")
	if ok {
		resp, err := rpc.UserQuery(rpc.TokenTypeAccess, accessToken)
		if resp.BaseResp.ErrNo != 0 || err != nil {
			c.JSONP(500, err)
			return
		}
		c.Set("userID", resp.User.Id)
	}

	c.Next()
}
