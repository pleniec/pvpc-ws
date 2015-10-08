package authentication

import (
	"gopkg.in/redis.v3"
	"log"
)

var (
	rc *redis.Client
)

func init() {
	rc = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	_, err := rc.Ping().Result()
	if err != nil {
		log.Fatal("can't connect to redis (", err, ")")
	}
}

func AuthenticateRequest(at string) (int64, error) {
	id, err := rc.Get("access_token:" + at).Int64()
	if err != nil {
		if err.Error() != "redis: nil" {
			return -1, err
		} else {
			return -1, nil
		}
	}

	return id, nil
}
