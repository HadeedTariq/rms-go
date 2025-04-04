package utils

import (
	"fmt"
	"os"
	"rms-platform/models"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func GenerateTokens(existingUser models.User) (string, string, error) {
	fmt.Println(existingUser)
	access_secret := os.Getenv("ACCESS_SECRET")
	refresh_secret := os.Getenv("REFRESH_SECRET")
	accessClaims := jwt.MapClaims{
		"user_id":   existingUser.ID,
		"user_name": existingUser.Username,
		"email":     existingUser.Email,
		"role":      existingUser.Role,
		"exp":       time.Now().Add(time.Minute * 15).Unix(),
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)
	accessSignedToken, err := accessToken.SignedString([]byte(access_secret))
	if err != nil {
		return "", "", err
	}

	refreshClaims := jwt.MapClaims{
		"user_id": existingUser.ID,
		"exp":     time.Now().Add(time.Hour * 24 * 7).Unix(), // 1 week expiration
	}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
	refreshSignedToken, err := refreshToken.SignedString([]byte(refresh_secret))
	if err != nil {
		return "", "", err
	}

	return accessSignedToken, refreshSignedToken, nil
}
