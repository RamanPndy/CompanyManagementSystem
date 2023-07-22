package auth

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	jwt "github.com/golang-jwt/jwt"
)

func GenerateToken(user_id int) (string, error) {
	tokenLifeSpan := os.Getenv("TOKEN_HOUR_LIFESPAN")
	if tokenLifeSpan == "" {
		return "", errors.New("empty token life span")
	}

	apiSecret := os.Getenv("API_SECRET")
	if apiSecret == "" {
		return "", errors.New("empty API secret")
	}

	token_lifespan, err := strconv.Atoi(tokenLifeSpan)

	if err != nil {
		return "", err
	}

	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["user_id"] = user_id
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(token_lifespan)).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(apiSecret))

}

func TokenValid(token, bearerToken string) error {
	apiSecret := os.Getenv("API_SECRET")
	if apiSecret == "" {
		return errors.New("empty API secret")
	}

	tokenString := ExtractToken(token, bearerToken)
	_, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(apiSecret), nil
	})
	if err != nil {
		return err
	}
	return nil
}

func ExtractToken(token, bearerToken string) string {
	if token != "" {
		return token
	}
	if len(strings.Split(bearerToken, " ")) == 2 {
		return strings.Split(bearerToken, " ")[1]
	}
	return ""
}
