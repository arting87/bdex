package db

import (
	"time"

	"github.com/garyburd/redigo/redis"
	"github.com/spf13/viper"
)

var Redis *redis.Pool

func setupRedis() {
	cfg := viper.Sub("redis")

	Redis = &redis.Pool{
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", cfg.GetString("host"))
			if err != nil {
				return nil, err
			}
			if cfg.GetString("password") != "" {
				if _, err := c.Do("AUTH", cfg.GetString("password")); err != nil {
					c.Close()
					return nil, err
				}
			}
			if _, err := c.Do("SELECT", cfg.GetInt("index")); err != nil {
				c.Close()
				return nil, err
			}
			return c, nil
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			if time.Since(t) < time.Minute {
				return nil
			}
			_, err := c.Do("PING")
			return err
		},
		MaxIdle:     cfg.GetInt("maxIdle"),
		IdleTimeout: 240 * time.Second,
	}
}
