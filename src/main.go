package main

import (
	"log"
	"os"
	"time"

	"github.com/MenheraBot/MenheraVanGOgh/src/controllers"
	"github.com/MenheraBot/MenheraVanGOgh/src/utils"
	"github.com/MenheraBot/MenheraVanGOgh/src/websocket"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type HttpPIng struct {
	Uptime int64 `json:"uptime"`
}

type WsPing struct {
	Id     uint8 `json:"id"`
	Uptime int64 `json:"uptime"`
	Ping   uint  `json:"ping"`
}

type PingStruct struct {
	Http HttpPIng `json:"http"`
	Ws   []WsPing `json:"ws"`
}

func main() {
	gin.SetMode(gin.ReleaseMode)

	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(cors.Default())

	httpStartTime := time.Now()
	websocketConnections := make(map[uint8]websocket.WsConnection)

	router.GET("/ping", func(c *gin.Context) {
		returnPing(c, httpStartTime, &websocketConnections)
	})

	Utilities := utils.New()

	router.GET("/ws", func(c *gin.Context) {
		websocket.ServeHTTP(c, &websocketConnections, &Utilities)
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

	utils.SetInterval(func() {
		for id, skt := range websocketConnections {
			if !skt.IsAlive {
				skt.Socket.Close()
				continue
			}

			skt.IsAlive = false
			skt.LastPingAt = time.Now()
			skt.Socket.WriteControl(9, []byte{}, time.Now().Add(time.Second*3))
			websocketConnections[id] = skt
		}
	}, 15000)

	log.Fatal(router.Run(":2080"))
}

func returnPing(c *gin.Context, startTime time.Time, ws *map[uint8]websocket.WsConnection) {
	toSend := make([]WsPing, 0)

	for k, v := range *ws {
		toSend = append(toSend, WsPing{
			Id:     k,
			Uptime: time.Since(v.Uptime).Milliseconds(),
			Ping:   v.Ping,
		})
	}

	http := HttpPIng{
		Uptime: time.Since(startTime).Milliseconds(),
	}

	returnData := PingStruct{
		Http: http,
		Ws:   toSend,
	}

	c.JSON(200, returnData)
}
