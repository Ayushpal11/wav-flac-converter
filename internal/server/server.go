package server

import (
	"net/http"

	"github.com/ayushpal11/wav-flac-converter/pkg/websocket"
)

func HandleWebSocket(w http.ResponseWriter, r *http.Request) {
	websocket.HandleConnection(w, r)
}

// type Server struct {
// 	router *gin.Engine
// }

// func NewServer() *Server {
// 	s := &Server{router: gin.Default()}
// 	s.router.GET("/", s.mainPage)
// 	s.router.GET("/ws", s.handleWebSocket)
// 	return s
// }

// func (s *Server) Start(addr string) error {
// 	return s.router.Run(addr)
// }

// func (s *Server) handleWebSocket(c *gin.Context) {
// 	websocket.HandleWebSocket(c.Writer, c.Request)
// }

// func (s *Server) mainPage(c *gin.Context) {
// 	c.JSON(200, gin.H{
// 		"message": "Hello, World!",
// 	})
// }
