package handlers

import (
	"github.com/Levor/birthday/internal/config"
	"github.com/Levor/birthday/internal/db/models"
	"github.com/Levor/birthday/internal/db/repositories"
	"github.com/Levor/birthday/internal/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserHandler struct {
	cfg         *config.Config
	ur          *repositories.UserRepository
	userService *services.UserService
}

func NewUserHandler(cfg *config.Config, ur *repositories.UserRepository, userService *services.UserService) *UserHandler {
	return &UserHandler{cfg: cfg, ur: ur, userService: userService}
}

func (h *UserHandler) Login(c *gin.Context) {
	var req models.User
	err := c.BindJSON(&req)
	if err != nil {
		HandleError(err, c)
	}
	token, err := h.userService.Login(req)
	if err != nil {
		HandleError(err, c)
	} else {
		c.JSON(http.StatusOK, gin.H{
			"token": token,
		})
	}

}
func (h *UserHandler) Logout(c *gin.Context) {
	var req models.User
	c.BindJSON(&req)
	err := h.userService.Logout(req)
	if err != nil {
		HandleError(err, c)
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status": http.StatusOK,
		})
	}

}

func (h *UserHandler) CreateNewUser(c *gin.Context) {
	var user models.User
	c.BindJSON(&user)
	newUser, err := h.userService.SignUp(&user)
	if err != nil {
		HandleError(err, c)
	}
	c.JSON(http.StatusOK, gin.H{
		"newUser": newUser,
	})
}

func (h *UserHandler) DeleteUser(c *gin.Context) {
	var id struct {
		ID int `json:"id"`
	}
	c.BindJSON(&id)
	err := h.ur.Delete(id.ID)
	if err != nil {
		HandleError(err, c)
	}
	c.JSON(http.StatusOK, "User deleted successful")
}

func (h *UserHandler) ChangePassword(c *gin.Context) {
	var user *models.User
	c.BindJSON(&user)
	err := h.userService.ChangePassword(user.Login, user.Password)
	if err != nil {
		HandleError(err, c)
	}
	c.JSON(http.StatusOK, "Password changed successful")
}
