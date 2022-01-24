package routes

import (
	"github.com/Levor/birthday/internal/config"
	"github.com/Levor/birthday/internal/http/handlers"
	"github.com/gin-gonic/gin"
)

func API(
	cfg *config.Config,
	handlerHealthCheck *handlers.HealthCheckHandler,
	workerHandler *handlers.WorkersHandler,
) *gin.Engine {
	r := gin.New()

	r.Use(
		gin.Recovery(),
		gin.Logger(),
	)

	r.GET("/healthcheck", handlerHealthCheck.GetHealthCheck)

	workerGroup := r.Group("worker")
	{
		workerGroup.GET("/get", workerHandler.GetWorkers)
		workerGroup.POST("/creat", workerHandler.Create)
		workerGroup.DELETE("/delete", workerHandler.Delete)
		workerGroup.POST("/update", workerHandler.Update)
	}

	return r
}
