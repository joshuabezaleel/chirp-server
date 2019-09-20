package auth

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"

	"golang.org/x/crypto/bcrypt"
)

// Service provides basic operations on for authentication service.
type Service interface {
	GetStoredPasswordByUsername(username string) (string, error)
	ComparePassword(incomingPassword, storedPassword string) (bool, error)
	AuthenticationMiddleware(next http.HandlerFunc) http.HandlerFunc
}

type service struct {
	auth Repository
}

// NewService creates an instance of the service for user domain model with necessary dependencies.
func NewService(auth Repository) Service {
	return &service{
		auth: auth,
	}
}

func (s *service) GetStoredPasswordByUsername(username string) (string, error) {
	password, err := s.auth.GetPassword(username)
	if err != nil {
		return "", err
	}
	return password, nil
}

func (s *service) ComparePassword(incomingPassword, storedPassword string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(storedPassword), []byte(incomingPassword))
	if err != nil {
		return false, err
	}

	return true, nil
}

func (s *service) AuthenticationMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		claims := jwt.MapClaims{}
		authorizationHeader := r.Header.Get("Authorization")

		if authorizationHeader != "" {
			bearerToken := strings.Split(authorizationHeader, " ")
			fmt.Println(bearerToken)

			if len(bearerToken) == 2 {
				token, err := jwt.ParseWithClaims(bearerToken[1], claims, func(token *jwt.Token) (interface{}, error) {
					if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
						return nil, fmt.Errorf("There was an error")
					}
					return SecretKey, nil
				})
				if err != nil {
					json.NewEncoder(w).Encode(err)
					return
				}

				if token.Valid {
					next(w, r)
				} else {
					json.NewEncoder(w).Encode("Invalid authorization token.")
				}
			}
		} else {
			json.NewEncoder(w).Encode("An authorization header is required.")
		}
	})
}
