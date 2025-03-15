package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte("secret_key")

var users = map[string]string {
	"user1": "password1",
	"user2": "password2",
}

type Credentials struct {
	Username string `json: "username"`
	Password string `json: "password"`
}

type Claims struct {
	Username string `json: "username"`
	jwt.StandardClaims
}

func Login(w http.ResponseWriter, r *http.Request) {
	var credentials Credentials
	err := json.NewDecoder(r.Body).Decode(&credentials)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return 
	}
	
	expectedPassword, ok := users[credentials.Username]

	if !ok || expectedPassword != credentials.Password {
		w.WriteHeader(http.StatusUnauthorized)
		return 
	}

	//De revizuit putin codul pe maine
	experationTime := time.Now().Add(time.Minute * 5)

	claims := &Claims{
		Username: credentials.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: experationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
}

func Home(w http.ResponseWriter, r *http.Request) {
	
}

func Refresh(w http.ResponseWriter, r *http.Request) {
	
}
