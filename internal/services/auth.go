package services

import (
	"time"

	"github.com/Akorex/caronago/internal/config"
	"github.com/Akorex/caronago/internal/dto"
	appError "github.com/Akorex/caronago/internal/errors"
	"github.com/Akorex/caronago/internal/models"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"

	"gorm.io/gorm"
)

type AuthService struct {
	DB *gorm.DB
	Config *config.Config

}

func NewAuthService(db *gorm.DB, config *config.Config) *AuthService{
	return &AuthService{
		DB: db,
		Config: config,
	}

}

func (s *AuthService) RegisterUser(payload dto.RegisterUser) (*models.User, error){
	var existingUser models.User

	if err := s.DB.Where("email = ?", payload.Email).First(&existingUser).Error; err == nil{
		return nil, appError.ErrConflict("User")
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(payload.Password), bcrypt.DefaultCost )
	if err != nil{
		return nil, appError.ErrInternalServerError(err)
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
		return nil, appError.ErrInternalServerError(err)
	}

	return &user, nil
}



func (s *AuthService) LoginUser(payload dto.LoginUser) (string, error){
	var user models.User

	if err := s.DB.Where("email = ?", payload.Email).First(&user).Error; err != nil{
		return "", appError.ErrBadRequest("Invalid email or password")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(payload.Password)); err != nil{
		return "", appError.ErrBadRequest("Invalid email or password")
	}

	token, err := s.generateToken(user.ID)
	if err != nil{
		return "", appError.ErrInternalServerError(err)
	}

	return token, nil

}


func (s *AuthService) GetUserByID(id uuid.UUID) (*models.User, error){
	var user models.User

	if err := s.DB.Where("id = ?", id).First(&user).Error; err != nil{
		return nil, appError.ErrNotFound("User")
	}

	return &user, nil
}





func (s *AuthService) generateToken(userID uuid.UUID) (string, error){
	claims := jwt.MapClaims{
		"sub": userID,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
		"iat": time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secret := s.Config.JWTSecret

	return token.SignedString([]byte(secret))
	
}