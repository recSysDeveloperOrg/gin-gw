package rpc

import (
	"context"
	"fmt"
	"gin-gateway/idl/gen/user"
)

var genderStr2Gender = map[string]user.Gender{
	GenderTypeMale:    user.Gender_GENDER_MALE,
	GenderTypeFemale:  user.Gender_GENDER_FEMALE,
	GenderTypeUnknown: user.Gender_GENDER_UNDEFINED,
}

const (
	GenderTypeMale    = "MALE"
	GenderTypeFemale  = "FEMALE"
	GenderTypeUnknown = "UNKNOWN"
)

const (
	TokenTypeAccess  = uint8(0)
	TokenTypeRefresh = uint8(1)
)

func UserQuery(tokenType uint8, tokenString string) (*user.QueryResp, error) {
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

func UserLogin(username, password string) (*user.LoginResp, error) {
	req := &user.LoginReq{
		Username: username,
		Password: password,
	}

	return UserClient.Login(context.Background(), req)
}

func UserRegister(username, password string, strGender string) (*user.RegisterResp, error) {
	req := &user.RegisterReq{
		User: &user.User{
			Name:     username,
			Password: password,
			Gender:   genderStr2Gender[strGender],
		},
	}

	return UserClient.Register(context.Background(), req)
}
