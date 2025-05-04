package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/jaiminbhaduri/codeprac/db"
	"github.com/jaiminbhaduri/codeprac/utils"
	"golang.org/x/crypto/bcrypt"
)

type AuthPayload struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Register(w http.ResponseWriter, r *http.Request) {
	var creds AuthPayload
	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	conn, err := db.DB()
	if err != nil {
		http.Error(w, "DB error", http.StatusInternalServerError)
		return
	}
	defer conn.Close()

	// Check if user exists
	var exists int
	err = conn.QueryRow("SELECT COUNT(*) FROM users WHERE username = ?", creds.Username).Scan(&exists)
	if err != nil || exists > 0 {
		http.Error(w, "User already exists", http.StatusBadRequest)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(creds.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Server error", http.StatusInternalServerError)
		return
	}

	_, err = conn.Exec("INSERT INTO users(username, email, password) VALUES (?, ?, ?)", creds.Username, creds.Email, hashedPassword)
	if err != nil {
		http.Error(w, "Insert error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

type User struct {
	ID       int
	Username string
	Email    string
	Password string
}

func Login(w http.ResponseWriter, r *http.Request) {
	var creds AuthPayload

	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	conn, err := db.DB()
	if err != nil {
		http.Error(w, "DB error", http.StatusInternalServerError)
		return
	}
	defer conn.Close()

	var user User
	err = conn.QueryRow("SELECT * FROM users WHERE username = ? AND email = ?", creds.Username, creds.Email).Scan(&user.ID, &user.Username, &user.Email, &user.Password)
	if err != nil || user.Password != creds.Password {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	// Compare hash with bcrypt
	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(creds.Password)) != nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	// Generate JWT
	token, err := utils.GenerateJWT(creds.Username, user.ID)
	if err != nil {
		http.Error(w, "JWT error", http.StatusInternalServerError)
		return
	}

	// Set token in cookie
	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Value:    token,
		Path:     "/",
		HttpOnly: true,
	})
}

type ExecutePayload struct {
	Code     string `json:"code"`
	Language string `json:"language"`
}

func ExecuteCode(w http.ResponseWriter, r *http.Request) {
	var payload ExecutePayload
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// STUB: Replace with actual execution logic
	output := "Received code in " + payload.Language + ":\n" + payload.Code

	json.NewEncoder(w).Encode(map[string]string{
		"output": output,
	})
}
