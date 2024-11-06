package websocket

import (
	"log"
	"net/http"

	"github.com/ayushpal11/wav-flac-converter/internal/converter"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func HandleWebSocket(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Failed to upgrade connection:", err)
		return
	}
	defer conn.Close()

	for {
		_, audioData, err := conn.ReadMessage()
		if err != nil {
			log.Println("Error reading message:", err)
			break
		}

		flacData, err := converter.ConvertToFLAC(audioData)
		if err != nil {
			log.Println("Conversion error:", err)
			conn.WriteMessage(websocket.TextMessage, []byte("Conversion failed"))
			continue
		}

		if err := conn.WriteMessage(websocket.BinaryMessage, flacData); err != nil {
			log.Println("Error sending message:", err)
			break
		}
	}
}
