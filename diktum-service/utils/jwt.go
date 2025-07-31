package utils

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"io/ioutil"

	"github.com/golang-jwt/jwt/v5"
)

// Public key variable (loaded once)
var publicKey *rsa.PublicKey

// Claims - Customize sesuai kebutuhan
type Claims struct {
	jwt.RegisteredClaims
}

// init() reads public_key.pem at start
func init() {
	keyData, err := ioutil.ReadFile("keys/public_key.pem")
	if err != nil {
		panic("failed to read public key: " + err.Error())
	}

	block, _ := pem.Decode(keyData)
	if block == nil || block.Type != "PUBLIC KEY" {
		panic("failed to decode PEM block containing public key")
	}

	parsedKey, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		panic("failed to parse public key: " + err.Error())
	}

	var ok bool
	publicKey, ok = parsedKey.(*rsa.PublicKey)
	if !ok {
		panic("not RSA public key")
	}
}

// ValidateJWT validates a JWT signed with RS256
func ValidateJWT(tokenStr string) (*Claims, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		// Only allow RS256
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return publicKey, nil
	})

	if err != nil || !token.Valid {
		return nil, errors.New("invalid or expired token")
	}

	return claims, nil
}
