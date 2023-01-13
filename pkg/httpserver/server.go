package httpserver

import (
	"net/http"
	//"fmt"
)

type Server struct {
	server *http.Server
}

func (s *Server) Start(port string, handler http.Handler) error {
	s.server = &http.Server{
		Addr:    ":" + port,
		Handler: handler,
	}

	return s.server.ListenAndServe()
}