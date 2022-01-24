package main

import (
	"github.com/Levor/birthday/internal/config"
	"github.com/Levor/birthday/internal/di"
	"github.com/gin-gonic/gin"
)

func main() {
	c := di.Container

	err := c.Invoke(func(
		api *gin.Engine,
		config *config.Config,
	) {
		err := api.Run(":" + config.ServerPort)
		if err != nil {
			panic(err)
		}
	})
	if err != nil {
		panic(err)
	}
}
