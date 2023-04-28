package handler

import (
	"github.com/bootcamp-go/desafio-cierre-db.git/internal/challenges"
	"github.com/gin-gonic/gin"
)

type Challenges struct {
	s challenges.Service
}

func NewHandlerChallenges(s challenges.Service) *Challenges {
	return &Challenges{s}
}

func (c *Challenges) GetTotalsByCustomerCondition() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		result, err := c.s.GetTotalsByCustomerCondition()
		if err != nil {
			ctx.JSON(500, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(200, result)
	}
}

func (c *Challenges) GetTopSoldProducts() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		result, err := c.s.GetTopSoldProducts()
		if err != nil {
			ctx.JSON(500, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(200, result)
	}
}

func (c *Challenges) GetTopActiveCustomersSpent() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		result, err := c.s.GetTopActiveCustomersSpent()
		if err != nil {
			ctx.JSON(500, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(200, result)
	}
}
