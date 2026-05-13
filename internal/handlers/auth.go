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

	if !utils.BindAndValidate(c, &payload){
		return
	}


	user, err := h.Service.RegisterUser(payload)
	if err != nil{
		utils.SendError(c, http.StatusBadRequest, err.Error())
		return
	}

	utils.SendSuccess(c, http.StatusCreated, "User registered successfully", user)

}


func (h *AuthHandler) Login(c *gin.Context){
	var payload dto.LoginUser

	if !utils.BindAndValidate(c, &payload){
		return
	}

	token, err := h.Service.LoginUser(payload)

	if err != nil{
		utils.SendError(c, http.StatusUnauthorized, err.Error())
		return
	}

	utils.SendSuccess(c, http.StatusOK, "Login successful", gin.H{"token": token})
}