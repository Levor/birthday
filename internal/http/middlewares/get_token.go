package middlewares

import (
	"errors"
	"github.com/Levor/birthday/internal/http/handlers"
	"github.com/Levor/birthday/internal/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetToken(services *services.Token) gin.HandlerFunc {
	return func(c *gin.Context) {
		signedToken := c.Request.Header.Get("Authorization")
		_, err := services.Verify(signedToken)
		if err != nil {
			handlers.HandleError(errors.New("token verification failed"), c)
			c.Abort()
			c.Status(http.StatusUnauthorized)
			return
		}
		c.Set("AccessToken", signedToken)
	}
}
