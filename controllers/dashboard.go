package controllers

import (
	"html/template"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/jaiminbhaduri/codeprac/utils"
)

func Dashboard(c *gin.Context) {
	// Get JWT from Authorization header or cookie
	tokenString := c.GetHeader("Authorization")
	if tokenString == "" {
		if cookie, err := c.Cookie("token"); err == nil {
			tokenString = cookie
		}
	}
	tokenString = strings.TrimPrefix(tokenString, "Bearer ")

	// Validate JWT
	claims, err := utils.ValidateJWT(tokenString)
	if err != nil {
		tmpl := template.Must(template.ParseFiles("templates/auth.html"))
		c.Status(http.StatusUnauthorized)
		tmpl.Execute(c.Writer, nil)
		return
	}

	// Render dashboard with username
	tmpl := template.Must(template.ParseFiles("templates/dashboard.html"))
	c.Status(http.StatusOK)
	tmpl.Execute(c.Writer, map[string]interface{}{
		"Username": claims.Username,
	})
}
