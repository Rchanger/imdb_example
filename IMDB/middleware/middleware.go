package middleware

import (
	"errors"
	"fynd/IMDB/models"
	"fynd/IMDB/modules/common"
	"net/http"
	"strings"
)

//validate :
func Validate(r *http.Request) error {
	if r.Method != models.GetRequest && !strings.Contains(r.URL.Path, "/auth/login") {
		token := r.Header.Get("authorization")
		return ValidateToken(token)
	}
	return nil

}

// ValidateToken: validate token for expiry and accessiblity of user
func ValidateToken(token string) error {
	if token == "" {
		if strings.TrimSpace(token) == "" {
			return errors.New("Unauthorized Access: Token is Invalid")
		}
	}
	tokenData, err := common.DecryptToken(token)
	if err != nil {
		return errors.New("Unauthorized Access: Token is Invalid")
	}
	if !tokenData.Valid {
		return errors.New("Unauthorized Access: Token is Invalid")
	}
	user := tokenData.Claims.(*models.JwtCustomClaims)
	if user.LoginID == "" || user.Role != models.AdminRole {
		return errors.New("Unauthorized Access: Token is Invalid")
	}
	return nil
}
