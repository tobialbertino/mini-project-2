package tokenize

import (
	"log"
	"miniProject2/pkg/config"

	"github.com/golang-jwt/jwt/v5"
)

var AccessTokenKey string = config.GetKeyConfig("ACCESS_TOKEN_KEY")

type AccountData struct {
	IDNum      int64
	RoleID     int64
	IsVerified bool
	IsActive   bool
	ExpiresAt  int64
}

type AccountClaims struct {
	jwt.RegisteredClaims
	IDNum      int64 `json:"id_actor,omitempty"`
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
func GetDataUserFromToken(token string) (AccountData, error) {
	// validasi dari token signature
	tokenDetail, err := VerifyAccessToken(token)
	if err != nil {
		return AccountData{}, err
	}

	// Cast data to map[string]interface{} and cast data["name"] to string
	claims := tokenDetail.Claims.(jwt.MapClaims)
	dataID := claims["id_actor"].(float64)
	RoleID := claims["role_id"].(float64)
	IsVerified := claims["is_verified"].(bool)
	IsActive := claims["is_active"].(bool)

	data := AccountData{
		IDNum:      int64(dataID),
		RoleID:     int64(RoleID),
		IsVerified: IsVerified,
		IsActive:   IsActive,
	}

	return data, nil
}
