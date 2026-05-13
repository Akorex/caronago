package handlers

import (
	"github.com/Akorex/caronago/internal/config"
	"github.com/Akorex/caronago/internal/services"
	"gorm.io/gorm"
)

type Handlers struct {
	Auth *AuthHandler
}

func RegisterHandlers(db *gorm.DB, config *config.Config) *Handlers{
	return &Handlers{
		Auth: NewAuthHandler(services.NewAuthService(db, config)),
	}
}