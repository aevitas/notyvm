package api

import (
	"errors"
	"net/http"
	"strconv"

	"aevitas.dev/veiled/inbound"
	"aevitas.dev/veiled/models"
	"aevitas.dev/veiled/names"
	"aevitas.dev/veiled/rng"
	"github.com/gin-gonic/gin"
)

func (s *Server) GetSeededName(ctx *gin.Context) {
	arg := ctx.Param("seed")
	seed, err := strconv.Atoi(arg)

	if seed < 0 {
		ctx.AbortWithError(http.StatusBadRequest, errors.New("seed can not be negative"))
		return
	}

	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	p := names.GeneratePerson(seed)

	ctx.JSON(http.StatusOK, p)
}

func (s *Server) GenerateRandomNames(ctx *gin.Context) {
	count := ctx.Query("count")
	num, err := strconv.Atoi(count)

	if err != nil {
		num = 1
	}

	var ret []models.Person
	for i := 0; i < num; i++ {
		seed := rng.RandomNumber()
		p := names.GeneratePerson(seed)

		ret = append(ret, p)
	}

	ctx.JSON(http.StatusOK, ret)
}

func (s *Server) HandleInbound(ctx *gin.Context) {
	err := inbound.ProcessInboundEmail(ctx, s.Cache)

	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusOK, nil)
}

func (s *Server) ListInboxMessages(ctx *gin.Context) {
	arg := ctx.Param("seed")
	seed, err := strconv.Atoi(arg)

	if seed < 0 {
		ctx.AbortWithError(http.StatusBadRequest, errors.New("seed can not be negative"))
		return
	}

	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	p := names.GeneratePerson(seed)

	ib, f := s.Cache.Get(p.EmailAddress)

	if !f {
		ctx.AbortWithStatus(http.StatusNotFound)
		return
	}

	ctx.JSON(http.StatusOK, ib)
}
