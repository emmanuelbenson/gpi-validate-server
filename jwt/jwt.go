package jwt

import (
	"errors"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

const (
	secret = "irianonaisebuozizaajujua.45istheAnswer"
)

// JClaims struct
type JClaims struct {
	jwt.StandardClaims
	CustomClaims map[string]string `json:"custom,omitempty"`
}

// GenerateToken receiver generates token for a successful user who has logged in
func (jc *JClaims) GenerateToken(userID string) string {
	jc.StandardClaims = jwt.StandardClaims{ExpiresAt: time.Now().Add(time.Hour).Unix()}
	jc.CustomClaims = map[string]string{"userID": userID}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jc)
	tokenString, err := token.SignedString([]byte(secret))

	if err != nil {
		return ""
	}

	return tokenString
}

// VerifyToken func takes token string as parameters
func (jc *JClaims) VerifyToken(jwtToken string) bool {
	token, err := jwt.Parse(
		jwtToken,
		func(token *jwt.Token) (interface{}, error) {
			if jwt.SigningMethodHS256 != token.Method {
				return nil, errors.New("Invalid signing algorithm")
			}
			return []byte(secret), nil
		},
	)

	if err == nil && token.Valid {
		return true
	}

	return false
}
