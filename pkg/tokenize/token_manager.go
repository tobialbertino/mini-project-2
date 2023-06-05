package tokenize

import (
	"log"
	"miniProject2/pkg/config"

	"github.com/golang-jwt/jwt/v5"
)

var AccessTokenKey string = config.GetKeyConfig("ACCESS_TOKEN_KEY")
var RefreshTokenKey string = config.GetKeyConfig("REFRESH_TOKEN_KEY")

type AccountClaims struct {
	jwt.RegisteredClaims
	ID         int64 `json:"jti,omitempty"`
	RoleID     int64 `json:"role_id"`
	IsVerified bool  `json:"is_verified"`
	IsActive   bool  `json:"is_active"`
	ExpiresAt  int64 `json:"exp,omitempty"`
}

// Create token
func GenerateAccessToken(claims jwt.Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte(AccessTokenKey))
	if err != nil {
		log.Printf("token.SignedString: %v", err)
		return "", err
	}

	return t, nil
}

func GenerateRefreshToken(claims jwt.Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	rt, err := token.SignedString([]byte(RefreshTokenKey))
	if err != nil {
		log.Printf("token.SignedString: %v", err)
		return "", err
	}

	return rt, nil
}

func VerifyRefreshToken(auth string) (*jwt.Token, error) {
	keyFunc := func(t *jwt.Token) (interface{}, error) {
		if t.Method.Alg() != "HS256" {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(RefreshTokenKey), nil
	}
	token, err := jwt.Parse(auth, keyFunc)
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, err
	}
	return token, nil
}

func VerifyAccessToken(auth string) (*jwt.Token, error) {
	keyFunc := func(t *jwt.Token) (interface{}, error) {
		if t.Method.Alg() != "HS256" {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(AccessTokenKey), nil
	}
	token, err := jwt.Parse(auth, keyFunc)
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, err
	}
	return token, nil
}

// helper get ID from Token AccessToken
func GetIdUserFromToken(token string) (string, error) {
	// validasi dari token signature
	tokenDetail, err := VerifyAccessToken(token)
	if err != nil {
		return "", err
	}

	// Cast data to map[string]interface{} and cast data["name"] to string
	claims := tokenDetail.Claims.(jwt.MapClaims)
	dataID := claims["ID"].(string)

	return dataID, nil
}
