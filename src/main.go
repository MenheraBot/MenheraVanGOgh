package main

import (
	"errors"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
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

	database, err := initializeDatabase()

	if err != nil {
		log.Panicf("Database error %s", err.Error())
	}

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

		expectedToken := os.Getenv("TOKEN")

		if expectedToken != "" && token != os.Getenv("TOKEN") {
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
			os.Remove(socketPath)
			socket, err := net.Listen("unix", socketPath)

			if err != nil {
				log.Println("Error to setup unix socket", err)
			} else {
				sigc := make(chan os.Signal, 1)
				signal.Notify(sigc, os.Interrupt, syscall.SIGTERM)

				go func(c chan os.Signal) {
					sig := <-c
					log.Printf("Caught signal %s: Closing socket.", sig)
					socket.Close()
					os.Exit(0)
				}(sigc)

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

func initializeDatabase() (*database.Database, error) {
	databaseNumber, err := strconv.Atoi(os.Getenv("REDIS_DB"))

	if err != nil {
		log.Print(err)
		databaseNumber = -1
	}

	address := os.Getenv("REDIS_ADDRESS")

	if address == "" {
		log.Print(errors.New("REDIS_ADDRESS is not set"))
	}

	return database.NewDatabase(address, databaseNumber)
}

func appendVanGOghRoutes(router *gin.Engine, db *database.Database) *gin.Engine {

	router.POST("/philo", controllers.Philo)

	router.POST("/astolfo", func(c *gin.Context) {
		controllers.Astolfo(c, db.ImageCache)
	})

	router.POST("/blackjack", func(c *gin.Context) {
		controllers.Blackjack(c, db.ImageCache)
	})

	router.POST("/8ball", func(c *gin.Context) {
		controllers.Eightball(c, db.ImageCache)
	})

	router.POST("/preview", func(c *gin.Context) {
		controllers.Preview(c, db.ImageCache)
	})

	router.POST("/ship", func(c *gin.Context) {
		controllers.Ship(c, db)
	})

	router.POST("/trisal", func(c *gin.Context) {
		controllers.Trisal(c, db)
	})

	router.POST("/roulette", func(c *gin.Context) {
		controllers.Roulette(c, db)
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
