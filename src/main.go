package main

import (
	"log"
	"os"
	"time"

	"github.com/MenheraBot/MenheraVanGOgh/src/controllers"
	"github.com/MenheraBot/MenheraVanGOgh/src/utils"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
)

type Http struct {
	Uptime int64 `json:"uptime"`
}

type Ws struct {
	Id     uint8 `json:"id"`
	Uptime int64 `json:"uptime"`
}

type PingStruct struct {
	Http Http `json:"http"`
	Ws   []Ws `json:"ws"`
}

func main() {

	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(cors.Default())
	router.Use(gzip.Gzip(gzip.BestSpeed))

	httpStartTime := time.Now()

	websocketConnections := make(map[uint8]time.Time)

	router.GET("/ping", func(c *gin.Context) {
		returnPing(c, httpStartTime, &websocketConnections)
	})

	router.Use(func(c *gin.Context) {
		token := c.GetHeader("Authorization")

		if token != os.Getenv("TOKEN") {
			c.Status(401)
			c.Abort()
			return
		}

		c.Header("Content-Type", "image/png")

		c.Next()
	})

	router.GET("/ws", func(c *gin.Context) {
		utils.ServeHTTP(c.Writer, c.Request, &websocketConnections)
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
	router.POST("/profile", controllers.Profile)

	log.Println("Listening and serving HTTP on :2080")

	log.Fatal(router.Run(":2080"))
}

func returnPing(c *gin.Context, startTime time.Time, ws *map[uint8]time.Time) {
	now := time.Now()

	toSend := make([]Ws, 0)

	for k, v := range *ws {
		uptime := int64(now.Sub(v).Milliseconds())

		toSend = append(toSend, Ws{Id: k, Uptime: uptime})
	}

	http := Http{
		Uptime: now.Sub(startTime).Milliseconds(),
	}

	returnData := PingStruct{
		Http: http,
		Ws:   toSend,
	}

	c.JSON(200, returnData)
}
