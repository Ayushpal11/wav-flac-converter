package server

import (
	"github.com/ayushpal11/wav-flac-converter/pkg/websocket"
	"github.com/gin-gonic/gin"
)

type Server struct {
	router *gin.Engine
}

func NewServer() *Server {
	s := &Server{router: gin.Default()}
	s.router.GET("/ws", s.handleWebSocket)
	return s
}

func (s *Server) Start(addr string) error {
	return s.router.Run(addr)
}

func (s *Server) handleWebSocket(c *gin.Context) {
	websocket.HandleWebSocket(c.Writer, c.Request)
}
