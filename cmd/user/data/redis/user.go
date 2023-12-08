package redis

import (
	"time"
)

func CheckUser(key string) (int, error) {
	res, err := Redis.Exists(key).Result()
	if err != nil {
		return 0, err
	}
	return int(res), nil
}

func SetUser(userid string, userinfo string, random int) error {
	rand := time.Duration(random)
	duration := rand*time.Minute + 24*time.Hour
	_, err := Redis.Set(userid, userinfo, duration).Result()
	return err
}

// username -> ()
func SetUserIndex(username string, userid string, random int) error {
	rand := time.Duration(random)
	//随机过期时间防止雪崩
	duration := rand*time.Minute + 24*time.Hour
	_, err := Redis.Set(username, userid, duration).Result()
	return err
}
func GetUserIndex(username string) (string, error) {
	userId, err := Redis.Get(username).Result()
	if err != nil {
		return "", err
	}
	return userId, nil
}

// id -> protoUser
func GetUser(userid string) (string, error) {
	res, err := Redis.Get(userid).Result()
	if err != nil {
		return "", err
	}
	return res, nil
}
