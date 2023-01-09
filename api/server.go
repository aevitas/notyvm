package api

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Server struct {
	Router *gin.Engine
}

func (s *Server) Init() {
	s.Router = gin.Default()

	s.Router.GET("v1/persons/random", s.GenerateRandomNames)
	s.Router.GET("v1/persons/:seed", s.GetSeededName)
}

func (s *Server) Start(ep string) {

	if !s.Ready() {
		log.Fatal("server isnt ready - make sure to init first")
	}

	err := http.ListenAndServe(ep, s.Router.Handler())

	if err != nil {
		log.Fatal(err)
	}
}

func (s *Server) Ready() bool {
	return s.Router != nil
}
