package database

import (
	"errors"
	"log"
	"time"
)

func (db *Database) GetCachedProfileImage(userId, profileString string) (string, error) {
	ctx, finishCtx := RedisContext()

	res, err := db.Client.Get(ctx, "profile:"+userId+":string").Result()
	finishCtx()

	if err != nil {
		return "", err
	}

	if res == profileString {
		ctx, finishCtx = RedisContext()

		res, err = db.Client.Get(ctx, "profile:"+userId+":image").Result()
		finishCtx()

		if err != nil {
			return "", err
		}

		return res, nil
	}

	return "", errors.New("no image found")
}

func (db *Database) SetCachedProfileImage(userId, profileString, image string) error {
	ctx := BackgroundContext()

	err := db.Client.Set(ctx, "profile:"+userId+":image", image, time.Minute*30).Err()

	if err != nil {
		log.Print(err)
		return err
	}

	err = db.Client.Set(ctx, "profile:"+userId+":string", profileString, time.Minute*30).Err()

	return err
}
