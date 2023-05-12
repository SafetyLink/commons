package jwt

import (
	_ "embed"
	"github.com/SafetyLink/commons/errors"

	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"log"
	"time"
)

type Claims struct {
	UserId int64 `json:"user_id"`
	jwt.RegisteredClaims
}

// In Production, You should generate your own private key and public key
// You can use this command to generate private key and public key
// openssl genrsa -out private.pem 2048
//
//go:embed secret/id_rsa
var _prvKey []byte

//go:embed secret/id_rsa.pub
var _publicKey []byte

// GenerateJwt generates a jwt token with userID
func GenerateJwt(userID int64) string {
	var err error
	// Load the private key
	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(_prvKey)
	if err != nil {
		log.Println(err)
		return ""
	}

	// Create the claims
	claims := Claims{
		userID,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(100 * time.Hour)),
		},
	}
	// Sign the token
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	signedToken, err := token.SignedString(privateKey)
	if err != nil {
		log.Println(err)
		return ""
	}
	return signedToken
}

// VerifyJwt verifies a jwt token and returns the claims
func VerifyJwt(token string) (*Claims, error) {
	var err error

	// Load the public key
	parsedPublicKey, err := jwt.ParseRSAPublicKeyFromPEM(_publicKey)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	// Parse the token
	decodedToken, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return parsedPublicKey, nil
	})
	if err != nil {
		return nil, err
	}
	//check if the claims are valid if are they return the claims else return an error
	if claims, ok := decodedToken.Claims.(*Claims); ok && decodedToken.Valid {
		return claims, nil

	} else {
		return nil, errors.ErrInvalid
	}
}
