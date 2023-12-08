package utils

import "golang.org/x/crypto/bcrypt"

// 功能函数，密码加密存储
func EncodePassword(password string) (string, error) {
	//加密方式：hash(password+randomnum)
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost) //加密处理
	if err != nil {
		return password, err
	}
	encodePWD := string(hash)
	return encodePWD, nil
}
