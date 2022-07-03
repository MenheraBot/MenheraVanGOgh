package main

import (
	"log"
	"os"
	"time"

	"github.com/MenheraBot/MenheraVanGOgh/src/controllers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type HttpPing struct {
	Uptime int64 `json:"uptime"`
}

type PingStruct struct {
	Http HttpPing `json:"http"`
}

func main() {
	gin.SetMode(gin.ReleaseMode)

	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(cors.Default())

	httpStartTime := time.Now()

	router.GET("/ping", func(c *gin.Context) {
		returnPing(c, httpStartTime)
	})

	router.Use(func(c *gin.Context) {
		token := c.GetHeader("Authorization")

		if token != os.Getenv("TOKEN") {
			c.Status(401)
			c.Abort()
			return
		}

		c.Header("Content-Type", "text/plain")

		c.Next()
	})

	router.POST("/astolfo", controllers.Astolfo)
	router.POST("/philo", controllers.Philo)
	router.POST("/ship", controllers.Ship)
	router.POST("/trisal", controllers.Trisal)
	router.POST("/gado", controllers.Gado)
	router.POST("/macetava", controllers.Macetava)
	router.POST("/blackjack", controllers.Blackjack)
	router.POST("/8ball", controllers.Eightball)
	router.POST("/vasco", controllers.Vasco)
	router.POST("/preview", controllers.Preview)
	router.POST("/profile", controllers.Profile)

	log.Println("Listening and serving HTTP on :2080")

	log.Fatal(router.Run(":2080"))
}

func returnPing(c *gin.Context, startTime time.Time) {

	http := HttpPing{
		Uptime: time.Since(startTime).Milliseconds(),
	}

	returnData := PingStruct{
		Http: http,
	}

	c.JSON(200, returnData)
}
