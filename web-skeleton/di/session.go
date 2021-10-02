package di

import (
	"time"

	"github.com/Summer2021-Project/di"
	"github.com/Summer2021-Project/dotenv"
	"github.com/go-session/redis"
	"github.com/go-session/session"
)

func init() {
	obj := di.Object{
		Name: "session",
		New: func() (i interface{}, e error) {
			opts := redis.Options{
				Addr:        dotenv.Getenv("REDIS_ADDR").String(),
				Password:    dotenv.Getenv("REDIS_PASSWORD").String(),
				DB:          int(dotenv.Getenv("REDIS_DATABASE").Int64()),
				DialTimeout: time.Duration(dotenv.Getenv("REDIS_DIAL_TIMEOUT").Int64(10)) * time.Second,
			}
			opt := redis.NewRedisStore(&opts)
			return session.NewManager(session.SetStore(opt)), nil
		},
	}
	if err := di.Provide(&obj); err != nil {
		panic(err)
	}
}

func Session() (manager *session.Manager) {
	if err := di.Populate("session", &manager); err != nil {
		panic(err)
	}
	return
}
