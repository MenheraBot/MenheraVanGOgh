package main

import (
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

type Http struct {
	Uptime int64 `json:"uptime"`
}

type Ws struct {
	Id     uint8 `json:"id"`
	Ping   uint16 `json:"ping"`
	Uptime int64 `json:"uptime"`
}

type PingStruct struct {
	Http Http `json:"http"`
	Ws   []Ws `json:"ws"`
}

func main() {
	app := fiber.New(fiber.Config{
		DisableStartupMessage: true,
		ReduceMemoryUsage: true,
		StreamRequestBody: true,
		Prefork: true,
	})

	app.Use(cors.New())

	app.Use(compress.New(compress.Config{
		Level: compress.LevelBestSpeed,
	}))

	startTime := time.Now()

	app.Get("/ping", func(c *fiber.Ctx) error {
		return returnPing(c, startTime)
	})

	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404) // => 404 "Not Found"
	})

	if fiber.IsChild() {
		log.Printf("[%d] Child Running\n", os.Getpid())
	} else {
		log.Printf("[%d] Master Running Port 2080\n", os.Getpid())
	}

	log.Fatal(app.Listen(":2080"))
}

func returnPing(c *fiber.Ctx, startTime time.Time) error {
	now := time.Now()

	ws := []Ws{}
	http := Http{
		Uptime: now.Sub(startTime).Milliseconds(),
	}

	returnData := PingStruct{
		Http: http,
		Ws:   ws,
	}

	return c.JSON(returnData)
}
