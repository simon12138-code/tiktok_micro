package request

import (
	"tiktok-api/pkg/myerr"
	"tiktok-common/errcode"
)

type UserCreateRequest struct {
	UserName     string `thrift:"user_name,1" frugal:"1,default,string" json:"user_name"`
	UserPassword string `thrift:"user_password,2" frugal:"2,default,string" json:"user_password"`
}

func (r UserCreateRequest) Verify() errcode.Err {
	if r.UserName == "" || len(r.UserName) > 32 || r.UserPassword == "" || len(r.UserPassword) > 32 {
		return myerr.ErrUserNameORPassWord
	}
	return nil
}

type UserLoginRequest struct {
	UserName     string `json:"user_name" from:"user_name" query:"user_name"`
	UserPassword string `json:"user_password" from:"user_password" query:"user_password"`
}

func (r UserLoginRequest) Verify() errcode.Err {
	if r.UserName == "" || len(r.UserName) > 32 || r.UserPassword == "" || len(r.UserPassword) > 32 {
		return myerr.ErrUserNameORPassWord
	}
	return nil
}
