package api

import (
	"errors"
	"net/http"
	"strconv"

	"aevitas.dev/veiled/messaging"
	"aevitas.dev/veiled/models"
	"aevitas.dev/veiled/names"
	"aevitas.dev/veiled/rng"
	"github.com/gin-gonic/gin"
)

func (s *Server) GetSeededName(ctx *gin.Context) {
	seed, err := strconv.Atoi(ctx.Param("seed"))

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
	num, err := strconv.Atoi(ctx.Query("count"))

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
	err := ProcessInboundEmail(ctx, s.Cache)

	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusOK, nil)
}

func (s *Server) GetInbox(ctx *gin.Context) {
	seed, err := strconv.Atoi(ctx.Param("seed"))

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
		ctx.JSON(http.StatusOK, messaging.EmptyInbox())
		return
	}

	ctx.JSON(http.StatusOK, ib)
}

func (s *Server) GetMessage(ctx *gin.Context) {
	seed, _ := strconv.Atoi(ctx.Param("seed"))
	id, _ := strconv.Atoi(ctx.Param("id"))

	if seed < 0 || id < 0 {
		ctx.AbortWithError(http.StatusBadRequest, errors.New("seed and message id can not be negative"))
		return
	}

	p := names.GeneratePerson(seed)

	inbox := messaging.EmptyInbox()
	ib, f := s.Cache.Get(p.EmailAddress)

	if f {
		inbox = ib.(messaging.Inbox)
	}

	m := inbox.GetMessage(uint64(id))

	if m == nil {
		ctx.AbortWithStatus(http.StatusNotFound)
		return
	}

	ctx.JSON(http.StatusOK, m)
}
