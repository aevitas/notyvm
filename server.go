package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"aevitas.dev/notyvm/names"
	"aevitas.dev/notyvm/rng"
	"github.com/gin-gonic/gin"
)

type Server struct {
	Router *gin.Engine
}

func (s *Server) Init() {
	s.Router = gin.Default()

	s.Router.GET("v1/name/random", s.RandomName)
	s.Router.GET("v1/name/random/:count", s.RandomNameMany)
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

type Persona struct {
	Seed         int    `json:"seed"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	EmailAddress string `json:"email_address"`
}

func (s *Server) RandomName(ctx *gin.Context) {
	arg := ctx.Query("seed")
	seed, err := strconv.Atoi(arg)

	if seed < 0 {
		ctx.AbortWithError(http.StatusBadRequest, errors.New("seed can not be negative"))
		return
	}

	if err != nil {
		seed = rng.RandomNumber()
	}

	n := names.RandomName(seed)

	p := Persona{Seed: seed, FirstName: n.FirstName, LastName: n.LastName, EmailAddress: strings.ToLower(fmt.Sprintf("%s.%s@notyvm.com", n.FirstName, n.LastName))}

	ctx.JSON(http.StatusOK, p)
}

func (s *Server) RandomNameMany(ctx *gin.Context) {
	count := ctx.Param("count")
	num, err := strconv.Atoi(count)

	if err != nil {
		num = 1
	}
	var ret []Persona
	for i := 0; i < num; i++ {
		seed := rng.RandomNumber()
		n := names.RandomName(seed)

		p := Persona{Seed: seed, FirstName: n.FirstName, LastName: n.LastName, EmailAddress: strings.ToLower(fmt.Sprintf("%s.%s@notyvm.com", n.FirstName, n.LastName))}
		ret = append(ret, p)
	}

	ctx.JSON(http.StatusOK, ret)
}
