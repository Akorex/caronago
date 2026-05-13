package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/Akorex/caronago/internal/config"
	"github.com/Akorex/caronago/internal/utils"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

func AuthMiddleware(cfg *config.Config) gin.HandlerFunc{
	return func(c *gin.Context){
		authHeader := c.GetHeader("Authorization")

		if authHeader == ""{
			utils.SendError(c, http.StatusUnauthorized, "Authorization header is required")
			c.Abort()
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == authHeader{
			utils.SendError(c, http.StatusUnauthorized, "Invalid authorization header")
			c.Abort()
			return
		}

		token, err := jwt.Parse(tokenString, func (token *jwt.Token) (any, error){
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok{
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(cfg.JWTSecret), nil
		})

		if err != nil{
			utils.SendError(c, http.StatusUnauthorized, "Invalid token")
			c.Abort()
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid{
			subStr, ok := claims["sub"].(string)
    		if !ok {
        		utils.SendError(c, http.StatusUnauthorized, "Invalid token claims: sub missing")
        		c.Abort()
        		return
    		}

			userId, err := uuid.Parse(subStr)
			if err != nil {
				utils.SendError(c, http.StatusUnauthorized, "Invalid user ID format in token")
				c.Abort()
				return
			}
			c.Set("userId", userId)
			c.Next()
			return	
		}
	
	}

}