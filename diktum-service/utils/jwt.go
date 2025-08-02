package utils

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
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
		// Debug: cetak algoritma token
		fmt.Println("🔍 Token alg:", token.Header["alg"])

		// Pastikan RS256
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return publicKey, nil
	})

	if err != nil {
		fmt.Println("❌ JWT Parse error:", err)
		return nil, fmt.Errorf("JWT parse error: %w", err)
	}

	if !token.Valid {
		fmt.Println("❌ JWT invalid (but parsed)")
		return nil, errors.New("token parsed but invalid")
	}

	fmt.Println("✅ JWT Valid!")
	fmt.Println("👤 Subject:", claims.Subject)
	fmt.Println("📅 Exp:", claims.ExpiresAt)
	fmt.Println("📅 Iat:", claims.IssuedAt)

	return claims, nil
}
