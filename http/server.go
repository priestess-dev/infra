package http

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type RouteHandler gin.HandlerFunc

type EndpointConfig struct {
	Path    string          `json:"path"`    // path to endpoint
	Method  string          `json:"method"`  // http method
	Handler gin.HandlerFunc `json:"handler"` // handler
}

type Server struct {
	Host   string
	Port   int
	Routes []EndpointConfig
	router *gin.Engine
}

func NewServer(host string, port int) *Server {
	return &Server{
		Host: host,
		Port: port,
	}
}

func (s *Server) AddRoutes(routes ...EndpointConfig) {
	if s.Routes == nil {
		s.Routes = make([]EndpointConfig, 0)
	}
	s.Routes = append(s.Routes, routes...)
}

func (s *Server) AddRoute(config EndpointConfig) {
	s.AddRoutes(config)
}

func (s *Server) Start() error {
	s.router = gin.Default()
	if s.Routes != nil {
		for _, route := range s.Routes {
			s.router.Handle(route.Method, route.Path, route.Handler)
		}
	} else {
		return fmt.Errorf("no routes")
	}
	return s.router.Run(fmt.Sprintf("%s:%d", s.Host, s.Port))
}
