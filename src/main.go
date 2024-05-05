package main

import (
	"errors"
	"log"
	"net"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/MenheraBot/MenheraVanGOgh/src/controllers"
	"github.com/MenheraBot/MenheraVanGOgh/src/database"
	"github.com/MenheraBot/MenheraVanGOgh/src/middleware"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type HttpPing struct {
	Uptime int64 `json:"uptime"`
}

type PingStruct struct {
	Http  HttpPing `json:"http"`
	Redis int64    `json:"redis"`
}

func main() {
	gin.SetMode(gin.ReleaseMode)

	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(cors.Default())
	router.Use(middleware.MetricsMiddleware())

	httpStartTime := time.Now()

	database := initializeDatabase()

	router.HEAD("/ping", func(c *gin.Context) {
		c.Header("Content-Type", "application/json; charset=utf-8")
		c.Status(http.StatusOK)
	})

	router.GET("/ping", func(c *gin.Context) {
		returnPing(c, httpStartTime, database)
	})

	router.Any("/metrics", gin.WrapH(promhttp.HandlerFor(middleware.GetCustomRegistry(), promhttp.HandlerOpts{})))

	router.Use(func(c *gin.Context) {
		token := c.GetHeader("Authorization")

		if token != os.Getenv("TOKEN") {
			c.Status(http.StatusUnauthorized)
			c.Abort()
			return
		}

		c.Header("Content-Type", "text/plain")

		c.Next()
	})

	router = appendVanGOghRoutes(router, database)

	socketPath := os.Getenv("SOCKET_PATH")

	if socketPath != "" {
		go func() {
			socket, err := net.Listen("unix", socketPath)

			if err != nil {
				log.Println("Error to setup unix socket")
				log.Panic(err)
			} else {
				defer socket.Close()

				log.Printf("Running unix socket on %s\n", socketPath)
				http.Serve(socket, router.Handler())
			}
		}()
	} else {
		log.Println("No SOCKET_PATH defined to setup unix socket server")
	}

	log.Println("Listening and serving HTTP on :2080")

	log.Fatal(router.Run(":2080"))
}

func initializeDatabase() *database.Database {
	databaseNumber, err := strconv.Atoi(os.Getenv("REDIS_DB"))

	if err != nil {
		log.Print(err)
		return &database.Database{Client: nil}
	}

	if os.Getenv("REDIS_ADDRESS") == "" {
		log.Print(errors.New("REDIS_ADDRESS is not set"))
		return &database.Database{Client: nil}
	}

	redis, err := database.NewDatabase(os.Getenv("REDIS_ADDRESS"), databaseNumber)

	if err != nil {
		log.Print(err)
		return &database.Database{Client: nil}
	}

	return redis

}

func appendVanGOghRoutes(router *gin.Engine, db *database.Database) *gin.Engine {

	router.POST("/astolfo", controllers.Astolfo)
	router.POST("/philo", controllers.Philo)
	router.POST("/blackjack", controllers.Blackjack)
	router.POST("/8ball", controllers.Eightball)
	router.POST("/preview", controllers.Preview)

	router.POST("/ship", func(c *gin.Context) {
		controllers.Ship(c, db)
	})

	router.POST("/trisal", func(c *gin.Context) {
		controllers.Trisal(c, db)
	})

	router.POST("/gado", func(c *gin.Context) {
		controllers.Gado(c, db)
	})

	router.POST("/macetava", func(c *gin.Context) {
		controllers.Macetava(c, db)
	})

	router.POST("/vasco", func(c *gin.Context) {
		controllers.Vasco(c, db)
	})

	router.POST("/profile", func(c *gin.Context) {
		controllers.Profile(c, db)
	})

	router.POST("/poker", func(c *gin.Context) {
		controllers.Poker(c, db)
	})

	return router
}

func returnPing(c *gin.Context, startTime time.Time, db *database.Database) {

	response := HttpPing{
		Uptime: time.Since(startTime).Milliseconds(),
	}

	var redisPing int64

	if db.Client == nil {
		redisPing = -1
	} else {
		startTime := time.Now()
		ctx, finishCtx := database.RedisContext()

		_, err := db.Client.Ping(ctx).Result()
		finishCtx()

		if err != nil {
			redisPing = -1
		} else {
			redisPing = time.Since(startTime).Milliseconds()
		}
	}

	returnData := PingStruct{
		Http:  response,
		Redis: redisPing,
	}

	c.JSON(http.StatusOK, returnData)
}
