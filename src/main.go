package main

import (
	"log"
	"os"
	"time"

	"github.com/MenheraBot/MenheraVanGOgh/src/controllers"
	"github.com/MenheraBot/MenheraVanGOgh/src/utils"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type HttpPIng struct {
	Uptime int64 `json:"uptime"`
}

type PingStruct struct {
	Http HttpPIng `json:"http"`
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

	Utilities := utils.New()

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

	router.POST("/astolfo", func(ctx *gin.Context) { controllers.Astolfo(ctx, &Utilities) })
	router.POST("/philo", func(ctx *gin.Context) { controllers.Philo(ctx, &Utilities) })
	router.POST("/ship", func(ctx *gin.Context) { controllers.Ship(ctx, &Utilities) })
	router.POST("/trisal", func(ctx *gin.Context) { controllers.Trisal(ctx, &Utilities) })
	router.POST("/gado", func(ctx *gin.Context) { controllers.Gado(ctx, &Utilities) })
	router.POST("/macetava", func(ctx *gin.Context) { controllers.Macetava(ctx, &Utilities) })
	router.POST("/blackjack", func(ctx *gin.Context) { controllers.Blackjack(ctx, &Utilities) })
	router.POST("/8ball", func(ctx *gin.Context) { controllers.Eightball(ctx, &Utilities) })
	router.POST("/vasco", func(ctx *gin.Context) { controllers.Vasco(ctx, &Utilities) })
	router.POST("/preview", func(ctx *gin.Context) { controllers.Preview(ctx, &Utilities) })
	router.POST("/profile", func(ctx *gin.Context) { controllers.Profile(ctx, &Utilities) })

	log.Println("Listening and serving HTTP on :2080")

	log.Fatal(router.Run(":2080"))
}

func returnPing(c *gin.Context, startTime time.Time) {

	http := HttpPIng{
		Uptime: time.Since(startTime).Milliseconds(),
	}

	returnData := PingStruct{
		Http: http,
	}

	c.JSON(200, returnData)
}
