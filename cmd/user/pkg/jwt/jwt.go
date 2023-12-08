/*
* @Author: zgy
* @Date:   2023/7/26 11:07
 */
package jwt

import (
	"github.com/dgrijalva/jwt-go"
	"tiktok/cmd/user/config"
	"tiktok/pkg/errno"
	"time"
)

type CustomClaims struct {
	ID       int
	NickName string

	jwt.StandardClaims
}

//func JWTAuth() gin.HandlerFunc {
//	return func(c *gin.Context) {
//		//1、获取token
//		token := c.Request.Header.Get("x-token")
//		color.Yellow(token)
//		if token == "" {
//
//			response.Err(c, http.StatusUnauthorized, 401, "请登录", "")
//			//先终止中间件调用
//			c.Abort()
//			return
//		}
//		//创建一个新的验证key
//		j := NewJWT()
//		//解析token
//		claims, err := j.ParseToken(token)
//		if err != nil {
//			if err == TokenExpired {
//				//token过期
//				response.Err(c, http.StatusUnauthorized, 401, "授权已过期", "")
//				c.Abort()
//				return
//			}
//			//其他错误
//			response.Err(c, http.StatusUnauthorized, 401, "未登陆", "")
//			c.Abort()
//			return
//		}
//		//打印上下文
//		fmt.Println(c)
//		// gin的上下文记录claims和userId的值
//		c.Set("claims", claims)
//		c.Set("userId", claims.ID)
//		c.Next()
//	}
//}

type JWT struct {
	SigningKey []byte
}

// 错误类型变量
var (
	TokenExpired     = errno.NewErrNo(errno.TokenErrCode, "Token is expired")
	TokenNotValidYet = errno.NewErrNo(errno.TokenErrCode, "Token not active yet")
	TokenMalformed   = errno.NewErrNo(errno.TokenErrCode, "That's not even a token")
	TokenInvalid     = errno.NewErrNo(errno.TokenErrCode, "Couldn't handle this token:")
)

func NewJWT() *JWT {
	return &JWT{
		[]byte(config.Settings.JWTKey.SigningKey),
	}
}

// 创建一个token，针对j的内部方法
func (j *JWT) CreateToken(claims CustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.SigningKey)
}

// 创建key的解析方法
func (j *JWT) ParseToken(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})
	if err != nil {
		//校验错误类型。断言如果错误是token非法
		if ve, ok := err.(*jwt.ValidationError); ok {
			//根据校验错误类型判断,类似与字符串匹配
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				// Token is expired
				return nil, TokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYet
			} else {
				return nil, TokenInvalid
			}
		}
	}
	if token != nil {
		//token合法且是对应的声明
		if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
			return claims, nil
		}
		return nil, TokenInvalid
	} else {
		return nil, TokenInvalid
	}

}

// 针对对应ID和名称生成的token
func CreateToken(Id int, NickName string) (string, error) {
	//生成token信息
	j := NewJWT()
	//可以配合自定义封装想要的信息，在完成token验证后还可以通过上下文进行信息保存和验证其他权限等
	claims := CustomClaims{
		//自定义
		ID:       Id,
		NickName: NickName,

		//官方
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix(),
			// TODO 设置token过期时间
			ExpiresAt: time.Now().Unix() + 24*30, //token -->1天过期
			Issuer:    "test",
		},
	}
	//生成token
	token, err := j.CreateToken(claims)
	if err != nil {

		return "", errno.ConvertErr(err).WithMessage("create token error")
	}
	return token, nil

}
