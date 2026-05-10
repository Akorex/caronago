package handlers

import (
	"net/http"

	"github.com/Akorex/caronago/internal/dto"
	"github.com/Akorex/caronago/internal/services"
	"github.com/Akorex/caronago/internal/utils"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	Service *services.AuthService
}

func NewAuthHandler(service *services.AuthService) *AuthHandler{
	return &AuthHandler{
		Service: service,
	}
}

func (h *AuthHandler) Register(c *gin.Context){
	var payload dto.RegisterUser

	if err := c.ShouldBindJSON(&payload); err != nil {
        utils.SendError(c, http.StatusBadRequest, "Invalid request payload")
        return
    }

	user, err := h.Service.RegisterUser(payload)
	if err != nil{
		utils.SendError(c, http.StatusBadRequest, err.Error())
		return
	}

	utils.SendSuccess(c, http.StatusCreated, "User registered successfully", user)

}