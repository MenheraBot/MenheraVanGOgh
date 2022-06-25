package websocket

import (
	"encoding/json"
	"net/http"
	"os"
	"time"

	"github.com/MenheraBot/MenheraVanGOgh/src/controllers"
	"github.com/MenheraBot/MenheraVanGOgh/src/utils"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
	EnableCompression: true,
}

func ServeHTTP(ctx *gin.Context, connections *map[uint8]time.Time, Utilities *utils.Utils) {

	var connectionDetails struct {
		Id   *uint8  `form:"id" binding:"required"`
		Auth *string `form:"auth" binding:"required"`
	}

	err := ctx.ShouldBindQuery(&connectionDetails)

	if err != nil {
		if connectionDetails.Id == nil {
			ctx.String(http.StatusBadRequest, "An invalid ID was provided!")
			return
		}

		if connectionDetails.Auth == nil {
			ctx.String(http.StatusBadRequest, "No Auth provided")
			return
		}

		return
	}

	if *connectionDetails.Auth != os.Getenv("TOKEN") {
		ctx.String(http.StatusUnauthorized, "Only Menhera client can access this!")
		return
	}

	if _, found := (*connections)[*connectionDetails.Id]; found {
		ctx.String(http.StatusConflict, "A connection with this ID already exists!")
		return
	}

	c, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)

	if err != nil {
		ctx.Status(http.StatusUpgradeRequired)
		return
	}

	defer c.Close()
	defer delete(*connections, *connectionDetails.Id)

	(*connections)[*connectionDetails.Id] = time.Now()

	nome := make(chan []byte)

	go (func() {
		for {
			msg, ok := <-nome

			if !ok {
				continue
			}

			go receiveMessage(msg, c, Utilities)
		}
	})()

	for {
		_, msg, err := c.ReadMessage()
		if err != nil {
			c.WriteJSON(map[string]bool{"error": true})
			break
		}

		nome <- msg
	}
}

func receiveMessage(msg []byte, c *websocket.Conn, Utilities *utils.Utils) {
	var received struct {
		RequestType string `json:"requestType"`
		Id          string `json:"id"`
	}

	json.Unmarshal(msg, &received)

	res := controllers.HandleWebsocketRequest(received.RequestType, msg, Utilities)

	var toSend struct {
		Id  string `json:"id"`
		Res string `json:"res"`
	}

	toSend.Id = received.Id
	toSend.Res = *res

	c.WriteJSON(toSend)
}
