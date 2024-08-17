package server

import (
	"fmt"
	"log"
	"net/http"

	"go-auth/internal/config"
)

type Server struct {
	Config *config.Config
}

func NewServer(config *config.Config) *Server {
	return &Server{
		Config: config,
	}
}

func (s *Server) Start() {
	fmt.Printf("Starting server at port %s...\n", s.Config.Port)
	if err := http.ListenAndServe(":"+s.Config.Port, nil); err != nil {
		log.Fatal("Error starting server: ", err)
	}
}
