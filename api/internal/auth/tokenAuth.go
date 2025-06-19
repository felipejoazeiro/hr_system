package auth

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte(os.Getenv("JWT_SECRET_KEY"))

type Claims struct {
	Login 		string 		`json:"login"`
	Position	string		`json:"position"`
	Role		string		`json:"role"`
	jwt.RegisteredClaims
}

func GenerateToken(login,position string)(string, error){
	expirationTime := time.Now().Add(24 *time.Hour)

	claims := &Claims{
		Login: login,
		Position: position,
		Role: determineRole(position),
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	return token.SignedString(jwtKey)
}


func determineRole(position string) string{
	switch position{
		case "Gerente":
			return "admin"
		case "Supervisor":
			return "manager"
		default:
			return "user"
	}
}