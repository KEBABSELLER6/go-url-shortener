package redis

import (
	"context"
	"errors"
	"os"

	"github.com/KEBABSELLER6/go-url-shortener/shortener"
	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

func SetIfNotExist(url string) (string, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_HOST"),
		Password: os.Getenv("REDIS_PW"),
		DB:       0,
	})

	storedShortId, getError := rdb.Get(ctx, url).Result()
	if getError == redis.Nil {
		newShortId, shortIdErr := shortener.GenerateShortId()

		if shortIdErr == nil {
			insertErr := rdb.Set(ctx, url, newShortId, 0).Err()

			if insertErr != nil {
				return "", errors.New("Error while insert")
			} else {
				return newShortId, nil
			}
		} else {
			return "", errors.New("Error while generating id")
		}

	} else if getError != nil {
		return "", errors.New("Error while query")
	} else {
		return storedShortId, nil
	}
}
