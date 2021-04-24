package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func LoginHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		username, password, _ := c.Request.BasicAuth()
		var authService AuthServiceInterface = AuthService{}
		user, err := authService.LoginService(username, password)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"status":  "error",
				"message": "not authorized",
			})
			c.Abort()
		}
		c.Set("Role", user.Role)
	}
}
