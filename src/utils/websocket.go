package utils

import (
	"net/http"
	"os"
	"strconv"
	"time"

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

func ServeHTTP(w http.ResponseWriter, r *http.Request, connections *map[uint8]time.Time) {

	id := r.URL.Query().Get("id")
	w.Header().Set("Content-Type", "text")

	if _, err := strconv.Atoi(id); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("An invalid ID was provided!"))
		return
	}

	if r.Header.Get("Authorization") != os.Getenv("TOKEN") {
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
		mt, message, err := c.ReadMessage()
		if err != nil {
			c.WriteJSON(map[string]bool{"error": true})
			break
		}

		err = c.WriteMessage(mt, message)
		if err != nil {
			break
		}
	}
}
