package errno

import (
	"errors"
	"fmt"
)

const (
	SuccessCode                = 0
	ServiceErrCode             = 10001
	ParamErrCode               = 10002
	UserAlreadyExistErrCode    = 10003
	AuthorizationFailedErrCode = 10004
	KeyMissErrCode             = 10005
	RedisErrCode               = 10006
	TokenErrCode               = 10007
	UserNotExistCode           = 10008
)

// 通过空变量进行接口实现校验
var _ error = &ErrNo{}

// 自定义错误返回类型
type ErrNo struct {
	ErrCode int64
	ErrMsg  string
}

// 继承error接口，实现自定义Error方法
func (e ErrNo) Error() string {
	return fmt.Sprintf("err_code=%d, err_msg=%s", e.ErrCode, e.ErrMsg)
}

func NewErrNo(code int64, msg string) ErrNo {
	return ErrNo{code, msg}
}

func (e ErrNo) WithMessage(msg string) ErrNo {
	e.ErrMsg = msg
	return e
}

// 常用封装错误类型
var (
	Success                = NewErrNo(SuccessCode, "Success")
	ServiceErr             = NewErrNo(ServiceErrCode, "Service is unable to start successfully")
	ParamErr               = NewErrNo(ParamErrCode, "Wrong Parameter has been given")
	UserAlreadyExistErr    = NewErrNo(UserAlreadyExistErrCode, "User already exists")
	AuthorizationFailedErr = NewErrNo(AuthorizationFailedErrCode, "Authorization failed")
	KeyMissErr             = NewErrNo(KeyMissErrCode, "key Miss")
	RedisErr               = NewErrNo(RedisErrCode, "redis error")
	UserNotExist           = NewErrNo(UserNotExistCode, "user Not exsit")
)

// ConvertErr convert error to Errno
func ConvertErr(err error) ErrNo {
	Err := ErrNo{}
	if errors.As(err, &Err) {
		return Err
	}

	s := ServiceErr
	s.ErrMsg = err.Error()
	return s
}
