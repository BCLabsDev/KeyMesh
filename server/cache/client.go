package cache

import (
	"context"
	"fmt"
	"keymesh/utils/config"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/v9"
)

var rdb *redis.Client

func init() {
	addr := fmt.Sprintf("%s:%d", config.Cache_Host, config.Cache_Port)
	rdb = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: "",
		DB:       0,
	})
	if err := rdb.Ping(context.Background()).Err(); err != nil {
		if !fiber.IsChild() {
			log.Fatalf("Redis 连接失败: %v", err)
		}
	}
}
