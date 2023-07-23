package utils

import (
	"html"
	"strings"

	"github.com/mitchellh/mapstructure"
	goErr "github.com/ralstan-vaz/go-errors"
	"golang.org/x/crypto/bcrypt"
)

// Bind ... uses mapstructure internally to bind input to output (does not use tags)
func Bind(input interface{}, output interface{}) error {
	err := mapstructure.Decode(input, output)
	if err != nil {
		return goErr.NewInternalError(err).SetCode("PKG.UTILS.DECODE_ERROR")
	}

	return nil
}

func Contains(slice []string, item string) bool {
	set := make(map[string]struct{}, len(slice))
	for _, s := range slice {
		set[s] = struct{}{}
	}

	_, ok := set[item]
	return ok
}

// GetPointerToString ...
func GetPointerToString(x string) *string {
	return &x
}

func CreateHash(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func VerifyHash(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func EscapeString(test string) string {
	return html.EscapeString(strings.TrimSpace(test))
}
