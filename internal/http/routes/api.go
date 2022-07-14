package routes

import (
	"github.com/Levor/birthday/internal/config"
	"github.com/Levor/birthday/internal/http/handlers"
	"github.com/Levor/birthday/internal/http/middlewares"
	"github.com/Levor/birthday/internal/services"
	"github.com/gin-gonic/gin"
)

func API(
	cfg *config.Config,
	handlerHealthCheck *handlers.HealthCheckHandler,
	workerHandler *handlers.WorkersHandler,
	userHandler *handlers.UserHandler,
	tokenService *services.Token,
) *gin.Engine {
	r := gin.New()

	r.Use(
		gin.Recovery(),
		gin.Logger(),
		middlewares.CorsMiddleware(cfg),
	)

	r.GET("/healthcheck", handlerHealthCheck.GetHealthCheck)
	privateGroup := r.Group("/private")
	{
		v1Group := privateGroup.Group("/v1")
		{
			workerGroup := v1Group.Group("/worker", middlewares.GetToken(tokenService))
			{
				workerGroup.GET("/get_all", workerHandler.GetWorkers)
				workerGroup.POST("/new", workerHandler.Create)
				workerGroup.DELETE("/delete", workerHandler.Delete)
				workerGroup.POST("/update", workerHandler.Update)
			}
		}
	}
	userGroup := r.Group("user")
	{
		userGroup.POST("/login", userHandler.Login)
		userGroup.POST("/logout", userHandler.Logout)
		userGroup.POST("/signup", userHandler.CreateNewUser)
		userGroup.DELETE("/delete", userHandler.DeleteUser)
		userGroup.POST("/changepassword", userHandler.ChangePassword)
	}

	return r
}
