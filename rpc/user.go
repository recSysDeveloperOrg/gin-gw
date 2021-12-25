package rpc

import (
	"context"
	"fmt"
	"gin-gateway/idl/gen/user"
)

const (
	TokenTypeAccess  = uint8(0)
	TokenTypeRefresh = uint8(1)
)

func QueryUser(tokenType uint8, tokenString string) (*user.QueryResp, error) {
	var req *user.QueryReq
	switch tokenType {
	case TokenTypeAccess:
		req = &user.QueryReq{
			AccessToken: tokenString,
		}
	case TokenTypeRefresh:
		req = &user.QueryReq{
			RefreshToken: tokenString,
		}
	default:
		return nil, fmt.Errorf("invalid token type:%d", tokenType)
	}

	return UserClient.Query(context.Background(), req)
}
