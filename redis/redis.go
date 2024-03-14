package redis

import (
	"encoding/json"
	"time"

	"github.com/anujsharma13/model"
	"github.com/go-redis/redis"
)

var client = redis.Nil

func Set(key string, value model.Workers) {
	if client == redis.Nil {
		client := redis.NewClient(&redis.Options{
			Addr:     "localhost:6379",
			Password: "",
			DB:       0,
		})
		client.Set(key, value, time.Hour)
	}

}

func Get(key string) (model.Workers, error) {
	var work model.Workers
	if client == redis.Nil {
		client := redis.NewClient(&redis.Options{
			Addr:     "localhost:6379",
			Password: "",
			DB:       0,
		})
		ans := client.Get(key)
		if ans.Err() != nil {
			return work, ans.Err()
		}
		value, err := ans.Result()
		if err != nil {
			return work, err
		}
		if err := json.Unmarshal([]byte(value), &work); err != nil {
			return work, err
		}
	}
	return work, nil
}

func Delete(key string) {
	if client == redis.Nil {
		client := redis.NewClient(&redis.Options{
			Addr:     "localhost:6379",
			Password: "",
			DB:       0,
		})
		client.Del(key)
	}
}
