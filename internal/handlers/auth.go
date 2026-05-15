package handlers

import (
	"net/http"

	"github.com/Akorex/caronago/internal/dto"
	"github.com/Akorex/caronago/internal/services"
	"github.com/Akorex/caronago/internal/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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
		c.Error(err)
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
		c.Error(err)
		return
	}

	utils.SendSuccess(c, http.StatusOK, "Login successful", gin.H{"token": token})
}

func (h *AuthHandler) GetMe(c *gin.Context){
	userID, ok := c.Get("userId")

	if !ok{
		utils.SendError(c, http.StatusUnauthorized, "Unauthorized")
		return
	}

	user, err := h.Service.GetUserByID(userID.(uuid.UUID))

	if err != nil{
		c.Error(err)
		return
	}

	utils.SendSuccess(c, http.StatusOK, "Profile successfully retrieved", user)

}