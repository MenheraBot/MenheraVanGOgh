package main

import (
	"log"
	"os"
	"runtime"
	"time"
	"unsafe"

	"github.com/MenheraBot/MenheraVanGOgh/src/controllers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type HttpPIng struct {
	Uptime int64 `json:"uptime"`
	Ping   uint  `json:"ping"`
}

type MemoryProfiler struct {
	Sys       uint64 `json:"alloc"`
	Timestamp int64  `json:"timestamp"`
}

type PingStruct struct {
	Http           HttpPIng         `json:"http"`
	MemoryProfiler []MemoryProfiler `json:"memory"`
}

var memory = make([]MemoryProfiler, 0)

func main() {
	gin.SetMode(gin.ReleaseMode)

	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(cors.Default())

	httpStartTime := time.Now()
	requestsMade := 0

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

		requestsMade++

		if requestsMade >= 2 {
			var m runtime.MemStats
			runtime.ReadMemStats(&m)
			memory = append(memory, MemoryProfiler{
				Sys:       m.Sys / 1024,
				Timestamp: time.Now().Unix(),
			})
			requestsMade = 0
		}

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

	router.DELETE("/memory", func(c *gin.Context) {
		memorySize := unsafe.Sizeof(memory)
		memoryLenArr := len(memory)
		memory = nil
		c.String(200, "Size of array: %d\nMemory size of array: %dMB", memoryLenArr, memorySize)
	})

	log.Println("Listening and serving HTTP on :2080")

	log.Fatal(router.Run(":2080"))
}

func returnPing(c *gin.Context, startTime time.Time) {

	http := HttpPIng{
		Uptime: time.Since(startTime).Milliseconds(),
	}

	returnData := PingStruct{
		Http:           http,
		MemoryProfiler: memory,
	}

	c.JSON(200, returnData)
}
