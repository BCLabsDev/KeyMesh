package main

import (
	_ "keymesh/cache"
	_ "keymesh/dao"

	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/logger"

	cfg "keymesh/utils/config"
)

func main() {
	app := fiber.New(fiber.Config{
		Prefork:      true,
		ServerHeader: "BCLABS/KeyMesh",
		JSONEncoder:  json.Marshal,
		JSONDecoder:  json.Unmarshal,
	})
	//  日志级别
	log.SetLevel(log.Level(cfg.Log_Level))

	// 日志中间件
	app.Use(logger.New())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, KeyMesh!")
	})

	app.Listen(":8080")
}
