package handlers

import (
	"github.com/Levor/birthday/internal/config"
	"github.com/Levor/birthday/internal/db/models"
	"github.com/Levor/birthday/internal/db/repositories"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type WorkersHandler struct {
	cfg *config.Config
	wr  *repositories.WorkersRepository
}

type Id struct {
	ID int `json:"id"`
}

func NewWorkersHandler(cfg *config.Config, wr *repositories.WorkersRepository) *WorkersHandler {
	return &WorkersHandler{cfg: cfg, wr: wr}
}

func (h *WorkersHandler) GetWorkers(c *gin.Context) {
	list, err := h.wr.FindAll()
	if err != nil {
		_ = c.Error(err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"workers": list,
	})
}

func (h *WorkersHandler) Create(c *gin.Context) {
	var worker models.Worker
	c.BindJSON(&worker)
	h.wr.Create(&worker)
	c.JSON(http.StatusOK, "Worker created successful")
}

func (h *WorkersHandler) Delete(c *gin.Context) {
	var id Id
	c.BindJSON(&id)
	err := h.wr.Delete(id.ID)
	if err != nil {
		log.Println(err)
	}
	c.JSON(http.StatusOK, "Worker deleted successful")
}

func (h *WorkersHandler) Update(c *gin.Context) {
	var worker models.Worker
	c.BindJSON(&worker)
	w, err := h.wr.FindByUserId(worker.ID)
	if err != nil {
		log.Println(err)
	}
	log.Println(w[0])
	err = h.wr.Update(w[0], &worker)
	if err != nil {
		log.Println(err)
	}
	c.JSON(http.StatusOK, "Worker updated successful")
}
