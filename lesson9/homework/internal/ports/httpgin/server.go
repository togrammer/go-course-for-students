package httpgin

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"homework9/internal/app"
)

type Server struct {
	port string
	app  *gin.Engine
}

func NewHTTPServer(port string, a app.App) *http.Server {
	gin.SetMode(gin.ReleaseMode)
	eng := gin.New()
	eng.Use(gin.Recovery())
	eng.Use(CustomMW)
	api := eng.Group("/api/v1")
	AppRouter(api, a)
	s := &http.Server{Addr: port, Handler: eng}
	return s
}

func (s *Server) Listen() error {
	return s.app.Run(s.port)
}

func (s *Server) Handler() http.Handler {
	return s.app
}
