package middlewares

import (
	"github.com/Levor/birthday/internal/config"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func CorsMiddleware(config *config.Config) gin.HandlerFunc {
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowMethods = config.Cors.Methods
	for _, origin := range config.Cors.Origins {
		if origin == "*" {
			corsConfig.AllowAllOrigins = true
		}
	}
	if !corsConfig.AllowAllOrigins {
		corsConfig.AllowOrigins = config.Cors.Origins
	}
	corsConfig.AllowHeaders = config.Cors.Headers

	return cors.New(corsConfig)
}
