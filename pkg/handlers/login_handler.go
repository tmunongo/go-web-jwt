package handlers

import (
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/tmunongo/go-web-jwt/pkg/middleware"
	"github.com/tmunongo/go-web-jwt/pkg/models"
)

func GenerateJWT(user models.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": user.Username,
		"userId": user.ID,
		"exp": time.Now().Add(time.Hour * 1).Unix(),
	})

	return token.SignedString(middleware.SecretKey)
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	username := r.PostFormValue("username")
	password := r.PostFormValue("password")

	user, err := models.GetUserByUsernameAndPassword(username, password)
	
	if err != nil {
		w.Write([]byte("Wrong username or password"))
		return
	}

	token, err := GenerateJWT(*user)

	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	w.Write([]byte(token))
}
