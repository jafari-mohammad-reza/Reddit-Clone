package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/reddit-clone/src/share/config"
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