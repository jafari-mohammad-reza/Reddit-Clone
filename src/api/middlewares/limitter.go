package middlewares

import (
	"Reddit-Clone/src/share/config"
	"github.com/gin-gonic/gin"
)

type LimitterMiddeware struct {
	cfg *config.Config
}

func NewLimmiterMiddlware(cfg *config.Config) *LimitterMiddeware {
	return &LimitterMiddeware{cfg}
}
func (l *LimitterMiddeware) RateLimiter() gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}
