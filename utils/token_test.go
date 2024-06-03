package utils

import (
	"github.com/dgrijalva/jwt-go"
	"testing"
)

func TestGenerateToken(t *testing.T) {
	userID := uint(1)
	token, err := GenerateToken(userID)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	claims := &Claims{}
	jwtToken, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil || !jwtToken.Valid {
		t.Fatalf("Expected valid token, got invalid token")
	}

	if claims.UserID != userID {
		t.Fatalf("Expected user ID %d, got %d", userID, claims.UserID)
	}
}

func TestParseToken(t *testing.T) {
	userID := uint(1)
	token, err := GenerateToken(userID)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	parsedUserID, err := ParseToken(token)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if parsedUserID != userID {
		t.Fatalf("Expected user ID %d, got %d", userID, parsedUserID)
	}
}

func TestParseInvalidToken(t *testing.T) {
	_, err := ParseToken("invalid.token.string")
	if err == nil {
		t.Fatalf("Expected error, got none")
	}
}
