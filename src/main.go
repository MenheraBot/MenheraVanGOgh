package main

import (
	"log"
	"os"
	"time"

	"github.com/MenheraBot/MenheraVanGOgh/src/controllers"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
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

	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(cors.Default())
	router.Use(gzip.Gzip(gzip.BestSpeed))

	//router.GET("/ws")
	// app.Get("/ws", websocket.New(setupWebsocket, websocket.Config{EnableCompression: true}))

	startTime := time.Now()

	router.GET("/ping", func(c *gin.Context) {
		returnPing(c, startTime)
	})

	router.Any("/", func(c *gin.Context) {
		token := c.GetHeader("Authorization")

		if token != os.Getenv("TOKEN") {
			c.Status(401)
			c.Abort()
			return
		}

		c.Header("Content-Type", "image/png")

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
	router.POST("/profile", controllers.Profile)

	log.Println("Server Running Port 2080")

	log.Fatal(router.Run(":2080"))
}

func returnPing(c *gin.Context, startTime time.Time) {
	now := time.Now()

	ws := []Ws{}
	http := Http{
		Uptime: now.Sub(startTime).Milliseconds(),
	}

	returnData := PingStruct{
		Http: http,
		Ws:   ws,
	}

	c.JSON(200, returnData)
}

/*
func setupWebsocket(c *websocket.Conn) {
	c.EnableWriteCompression(true)
	c.SetCompressionLevel(1)

	if _, err := strconv.Atoi(c.Query("id")); err != nil {
		c.WriteMessage(8, []byte("Invalid ID"))
		print("Erro %s", err)
	}

	var (
		mt  int
		msg []byte
		err error
	)
	for {
		if mt, msg, err = c.ReadMessage(); err != nil {
			log.Println("read:", err)
			break
		}

		log.Printf("recv: %s", msg)

		b, _ := json.Marshal(msg)

		println(string(b))

		c.WriteJSON(b)

		if err = c.WriteMessage(mt, []byte("filhos da puta")); err != nil {
			log.Println("write:", err)
			break
		}
	}
}
*/
