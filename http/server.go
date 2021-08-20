package http

import (
	"fmt"
	"net/http"
)

var (
	serv *Server
)

type Server struct {
	handler *Handler
	config  *Config
}

func Init(config *Config) {
	serv = NewServer(config)
}

func GetConfig() *Config {
	return serv.Config()
}

func NewServer(config *Config) *Server {
	return &Server{config: config, handler: nil}
}

func (s *Server) SetHandler(handler *Handler) {
	s.handler = handler
}

func (s *Server) Config() *Config {
	return s.config
}

func (s *Server) Start() {
	if s.handler == nil {
		panic("handler is nil")
	}

	http.ListenAndServe(fmt.Sprintf("%s:%d", s.config.Server.Addr.Host, s.config.Server.Addr.Port), s.handler)
}
