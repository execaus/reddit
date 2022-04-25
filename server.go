package main

import (
	"net/http"
	"time"
)

type Server struct {
	httpServer *http.Server
}

func (s *Server) Run(port string, router http.Handler) error {
	s.httpServer = &http.Server{
		Handler:        router,
		Addr:           ":" + port,
		MaxHeaderBytes: 1 << 20, // 1MB
		ReadTimeout:    time.Second * 10,
		WriteTimeout:   time.Second * 10,
	}
	return s.httpServer.ListenAndServe()
}
