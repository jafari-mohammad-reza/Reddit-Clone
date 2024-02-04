package healthcheck

import "github.com/gin-gonic/gin"

type HealthCheckController struct {
}

func NewHealthCheckController() *HealthCheckController {
	return &HealthCheckController{}
}

func (c *HealthCheckController) HealthCheck(ctx *gin.Context) {
	ctx.JSON(200, map[string]string{
		"msg": "Api is healthy",
	})
}
