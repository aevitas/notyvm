package api

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/patrickmn/go-cache"
)

type Server struct {
	Router *gin.Engine
	Cache  *cache.Cache
}

func (s *Server) Init() {
	s.Router = gin.Default()

	s.Cache = cache.New(60*time.Minute, 90*time.Minute)

	s.Router.GET("v1/persons/random", s.GenerateRandomNames)
	s.Router.GET("v1/persons/:seed", s.GetSeededName)
	s.Router.GET("v1/inbox/:seed", s.GetInbox)
	s.Router.GET("v1/inbox/:seed/:id", s.GetMessage)
	s.Router.POST("/inbound", s.HandleInbound)

	s.Router.GET("/healthz", func(ctx *gin.Context) { ctx.String(http.StatusOK, "healthy") })
}

func (s *Server) Start(ep string) {

	if !s.Ready() {
		log.Fatal("server isnt ready - make sure to init first")
	}

	e := make(chan error)

	e <- http.ListenAndServe(ep, s.Router.Handler())

	for err := range e {
		log.Fatal(err)
	}
}

func (s *Server) Ready() bool {
	return s.Router != nil
}
