package main

import (
	"log"
	"time"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/MenheraBot/MenheraVanGOgh/src/controllers"
)

type Http struct {
	Uptime int64 `json:"uptime"`
}

type Ws struct {
	Id     uint8  `json:"id"`
	Ping   uint16 `json:"ping"`
	Uptime int64  `json:"uptime"`
}

type PingStruct struct {
	Http Http `json:"http"`
	Ws   []Ws `json:"ws"`
}

func main() {
	app := fiber.New(fiber.Config{
		DisableStartupMessage: true,
		ReduceMemoryUsage:     true,
		StreamRequestBody:     true,
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
		token := c.Get(fiber.HeaderAuthorization)

		if token != os.Getenv("TOKEN") {
			return c.SendStatus(401)
		}	

		c.Set("Content-Type", "image/png")

		return c.Next()
	})

	app.Post("/astolfo", controllers.Astolfo)
	app.Post("/philo", controllers.Philo)
	app.Post("/trisal", controllers.Trisal)
	app.Post("/ship", controllers.Ship)
	app.Post("/gado", controllers.Gado)
	app.Post("/macetava", controllers.Macetava)
	app.Post("/8ball", controllers.Eightball)

	app.Use(func(c *fiber.Ctx) error {
		c.Set("Content-Type", "text")
		return c.Status(404).SendString("Welp, there is nothing for you right here")
	})

	log.Println("Server Running Port 2080")

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
