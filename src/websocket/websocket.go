package websocket

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/MenheraBot/MenheraVanGOgh/src/controllers"
	"github.com/MenheraBot/MenheraVanGOgh/src/utils"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
	EnableCompression: true,
	Error: func(w http.ResponseWriter, r *http.Request, status int, reason error) {
		w.Header().Set("Content-Type", "text")
		w.WriteHeader(426)
		w.Write([]byte("Upgrade Required"))
	},
}

type ReceivedObject struct {
	Type string `json:"type"`
	Id   string `json:"id"`
}

type ToSendObject struct {
	Id  string `json:"id"`
	Res string `json:"res"`
}

func ServeHTTP(w http.ResponseWriter, r *http.Request, connections *map[uint8]time.Time, Utilities *utils.Utils) {

	id := r.URL.Query().Get("id")
	w.Header().Set("Content-Type", "text")

	if _, err := strconv.Atoi(id); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("An invalid ID was provided!"))
		return
	}

	auth := r.URI.Query().Get("auth")

	if auth != os.Getenv("TOKEN") {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Only Menhera client can access this!"))
		return
	}

	idNumber, _ := strconv.ParseInt(id, 10, 8)

	if _, found := (*connections)[uint8(idNumber)]; found {
		w.WriteHeader(http.StatusConflict)
		w.Write([]byte("A connection with this ID already exists!"))
		return
	}

	c, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		w.WriteHeader(http.StatusUpgradeRequired)
		return
	}

	defer c.Close()
	defer delete(*connections, uint8(idNumber))

	(*connections)[uint8(idNumber)] = time.Now()

	for {
		_, msg, err := c.NextReader()
		if err != nil {
			c.WriteJSON(map[string]bool{"error": true})
			break
		}

		data, _ := ioutil.ReadAll(msg)

		received := &ReceivedObject{}
		json.NewDecoder(bytes.NewReader(data)).Decode(received)

		res := controllers.HandleWebsocketRequest(received.Type, bytes.NewReader(data), Utilities)

		toSend := &ToSendObject{}

		toSend.Id = received.Id
		toSend.Res = *res

		err = c.WriteJSON(toSend)

		if err != nil {
			break
		}
	}
}
