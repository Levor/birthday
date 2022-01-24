package handlers

import (
	"github.com/Levor/birthday/internal/config"
	"github.com/gin-gonic/gin"
	"net/http"
)

type HealthCheckHandler struct {
	cfg *config.Config
}

func NewHealthCheckHandler(cfg *config.Config) *HealthCheckHandler {
	return &HealthCheckHandler{cfg: cfg}
}

func (h *HealthCheckHandler) GetHealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"Status":  http.StatusOK,
		"version": "1.0.0",
	})
}
