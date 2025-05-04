package controllers

import (
	"html/template"
	"net/http"
	"strings"

	"github.com/jaiminbhaduri/codeprac/utils"
)

func Dashboard(w http.ResponseWriter, r *http.Request) {
	// Get JWT from cookie or Authorization header
	tokenString := r.Header.Get("Authorization")
	if tokenString == "" {
		cookie, err := r.Cookie("token")
		if err == nil {
			tokenString = cookie.Value
		}
	}

	tokenString = strings.TrimPrefix(tokenString, "Bearer ")

	// Validate JWT
	claims, err := utils.ValidateJWT(tokenString)
	if err != nil {
		tmpl := template.Must(template.ParseFiles("templates/auth.html"))
		tmpl.Execute(w, nil)
		return
	}

	// If valid, show dashboard
	tmpl := template.Must(template.ParseFiles("templates/dashboard.html"))
	tmpl.Execute(w, map[string]interface{}{
		"Username": claims.Username, // assuming username in JWT
	})
}
