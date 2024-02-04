package healthcheck

import "github.com/gin-gonic/gin"

type HealthCheckModule struct {
	controller *HealthCheckController
}

func initRoutes(r *gin.RouterGroup, controller *HealthCheckController) {
	healthCheck := r.Group("/health-check")
	healthCheck.GET("/health", controller.HealthCheck)
}

func NewHealthCheckModule(r *gin.RouterGroup) *HealthCheckModule {
	controller := NewHealthCheckController()
	initRoutes(r, controller)
	return &HealthCheckModule{}
}
