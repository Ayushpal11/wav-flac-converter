package websocket

import (
	"log"
	"net/http"

	"github.com/ayushpal11/wav-flac-converter/pkg/ffmpeg"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func HandleConnection(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Failed to set websocket upgrade:", err)
		return
	}
	defer conn.Close()

	_, wavData, err := conn.ReadMessage()
	if err != nil {
		log.Println("Error reading WAV data:", err)
		return
	}

	flacData, err := ffmpeg.ConvertToFLAC(wavData)
	if err != nil {
		log.Println("Error converting to FLAC:", err)
		return
	}

	conn.WriteMessage(websocket.BinaryMessage, flacData)
}
