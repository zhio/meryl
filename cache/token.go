package cache

import "time"

func SaveUserToken(token, userID string, exp time.Duration) error {
	err := RedisClient.Set(UserTokenKey(token), userID, exp).Err()
	return err
}

func GetUserByToken(token string) (string, error) {
	return RedisClient.Get(UserTokenKey(token)).Result()
}

func DelUserToken(token string) error {
	return RedisClient.Del(UserTokenKey(token)).Err()
}
