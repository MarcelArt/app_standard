package utils

import (
	"time"

	"github.com/MarcelArt/app_standard/config"
	"github.com/MarcelArt/app_standard/enums"
	"github.com/MarcelArt/app_standard/models"
	"github.com/golang-jwt/jwt/v5"
)

func GenerateTokenPair(user models.UserDTO, isRemember bool) (string, string, error) {
	accessToken, err := generateAccessToken(user)
	if err != nil {
		return "", "", err
	}

	refreshToken, err := generateRefreshToken(user, isRemember)
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}

func generateAccessToken(user models.UserDTO) (string, error) {
	claims := jwt.MapClaims{
		"username": user.Username,
		"userId":   user.ID,
		"exp":      time.Now().Add(time.Minute * 5).Unix(),
	}

	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	t, err := token.SignedString([]byte(config.Env.JwtSecret))

	return t, err
}

func generateRefreshToken(user models.UserDTO, isRemember bool) (string, error) {
	expireAt := time.Now().Add(enums.Day)
	if isRemember {
		expireAt = time.Now().Add(enums.Month)
	}

	claims := jwt.MapClaims{
		"userId":     user.ID,
		"isRemember": isRemember,
		"exp":        expireAt.Unix(),
	}

	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	t, err := token.SignedString([]byte(config.Env.JwtSecret))

	return t, err
}
