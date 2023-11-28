package http

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type RouteHandler gin.HandlerFunc

type Server struct {
	Host     string
	Port     int
	RouteMap map[string]gin.HandlerFunc
	router   *gin.Engine
}

func NewServer(host string, port int) *Server {
	return &Server{
		Host: host,
		Port: port,
	}
}

func (s *Server) AddRoutes(routeMap map[string]gin.HandlerFunc) {
	if s.RouteMap == nil {
		s.RouteMap = make(map[string]gin.HandlerFunc)
	}
	for path, handler := range routeMap {
		s.RouteMap[path] = handler
	}
}

func (s *Server) AddRoute(path string, handler gin.HandlerFunc) {
	if s.RouteMap == nil {
		s.RouteMap = make(map[string]gin.HandlerFunc)
	}
	s.RouteMap[path] = handler
}

func (s *Server) Start() error {
	s.router = gin.Default()
	if s.RouteMap == nil {
		return fmt.Errorf("route map is nil")
	}
	for path, handler := range s.RouteMap {
		s.router.GET(path, handler)
	}
	return s.router.Run(fmt.Sprintf("%s:%d", s.Host, s.Port))
}
