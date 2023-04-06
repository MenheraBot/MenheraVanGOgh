package database

import (
	"context"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
)

type Database struct {
	Client *redis.Client
}

func RedisContext() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), time.Second*2)
}

func BackgroundContext() context.Context {
	return context.Background()
}

func NewDatabase(address string, db int) (*Database, error) {
	client := redis.NewClient(&redis.Options{
		Addr:         address,
		Password:     "",
		DB:           db,
		ReadTimeout:  time.Second * 2,
		WriteTimeout: time.Second * 2,
	})

	ctx, finishCtx := RedisContext()

	if err := client.Ping(ctx).Err(); err != nil {
		finishCtx()
		return nil, err
	}

	log.Println("Connected to Redis")
	finishCtx()

	return &Database{
		Client: client,
	}, nil

}
