package services

import (
	"errors"

	"github.com/Akorex/caronago/internal/dto"
	"github.com/Akorex/caronago/internal/models"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthService struct {
	DB *gorm.DB

}

func NewAuthService(db *gorm.DB) *AuthService{
	return &AuthService{
		DB: db,
	}

}

func (s *AuthService) RegisterUser(payload dto.RegisterUser) (*models.User, error){
	var existingUser models.User

	if err := s.DB.Where("email = ?", payload.Email).First(&existingUser).Error; err == nil{
		return nil, errors.New("user already exists")
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(payload.Password), bcrypt.DefaultCost )
	if err != nil{
		return nil, err
	}

	

	user := models.User{
		ID:        uuid.New(),
		FirstName: payload.FirstName,
		LastName: payload.LastName,
		Email: payload.Email,
		Password: string(hashed),
		IsVerified: false,
	}

	if err := s.DB.Create(&user).Error; err != nil{
		return nil, err
	}

	return &user, nil
}