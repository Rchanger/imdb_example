package common

import (
	"errors"
	"fynd/IMDB/models"
	"log"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

// GenerateToken generates JWT token from Login object
func GenerateToken(user models.User) (string, error) {
	claims := models.JwtCustomClaims{
		user.Username,
		user.Role,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * models.JWT_EXPIRY_DURATION_IN_HOURS).Unix(),
		},
	}
	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte(models.Config.JWTSecret))
	if err != nil {
		log.Println("GenerateToken SignedString() Error: ", err)
		return "", errors.New("JWTUTILS_FAILED_TO_GENERATE_TOKEN")
	}
	return t, nil
}

//DecryptToken to get Login object from token
func DecryptToken(tknStr string) (*jwt.Token, error) {
	tknstrList := strings.Split(tknStr, models.BLANKSPACE)
	claims := &models.JwtCustomClaims{}
	tokenObj, err := jwt.ParseWithClaims(tknstrList[len(tknstrList)-1], claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(models.Config.JWTSecret), nil
	})
	if err != nil {
		log.Println("DecryptToken : error in token encryption : ", err, tokenObj)
	}
	return tokenObj, err
}
