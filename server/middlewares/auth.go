package middlewares

import (
	"dbo/helper"
	"dbo/service"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware(authService service.Auth) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if !strings.Contains(authHeader, "Bearer") {
			response := helper.ApiResponse("Unauthorization", http.StatusUnauthorized, "error", nil)
			c.JSON(http.StatusUnauthorized, response)
			return
		}

		tokenString := ""
		arrayToken := strings.Split(authHeader, " ")
		if len(arrayToken) == 2 {
			tokenString = arrayToken[1]
		}

		token, err := authService.ValidateToken(tokenString)

		if err != nil {
			response := helper.ApiResponse("Unauthorization", http.StatusUnauthorized, "error", nil)
			c.JSON(http.StatusUnauthorized, response)
			return
		}

		claim, ok := token.Claims.(jwt.MapClaims)

		if !ok || !token.Valid {
			response := helper.ApiResponse("Unauthorization", http.StatusUnauthorized, "error", nil)
			c.JSON(http.StatusUnauthorized, response)
			return
		}

		uid := int(claim["user_id"].(float64))
		name := claim["name"]
		email := claim["email"]

		c.Set("uid", uid)
		c.Set("name", name)
		c.Set("email", email)
		response := helper.ApiResponse("Authorized", http.StatusOK, "success", nil)
		c.JSON(http.StatusOK, response)
		return
	}
}
