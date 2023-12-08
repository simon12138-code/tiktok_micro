package redis

func GetUserVideoInfo(userids []string) ([]string, error) {
	res := make([]string, len(userids))
	resInterface, err := Redis.MGet(userids...).Result()
	for i, e := range resInterface {
		res[i] = e.(string)
	}
	if err != nil {
		return res, err
	}
	return res, nil
}
