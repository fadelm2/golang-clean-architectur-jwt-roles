package util

import (
	"context"
	"errors"
	"fmt"
	"golang-clean-architecture/internal/model"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

type TokenUtil struct {
	SecretKey string
}

func NewTokenUtil(secretKey string) *TokenUtil {
	return &TokenUtil{
		SecretKey: secretKey,
	}
}

func (t TokenUtil) CreateToken(ctx context.Context, auth *model.Auth) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":     auth.ID,
		"role":   auth.RoleID,
		"iat":    time.Now().Unix(),
		"expire": time.Now().Add(24 * time.Hour).Unix(),
	})

	jwtToken, err := token.SignedString([]byte(t.SecretKey))
	if err != nil {
		return "", err
	}

	return jwtToken, nil
}

func (t TokenUtil) ParseToken(ctx context.Context, jwtToken string) (*model.Auth, error) {
	token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
		return []byte(t.SecretKey), nil
	})
	if err != nil {
		return nil, fiber.ErrUnauthorized
	}
	claims := token.Claims.(jwt.MapClaims)

	expire := claims["expire"].(float64)
	if int64(expire) < time.Now().UnixMilli() {
		return nil, fiber.ErrUnauthorized
	}

	id := claims["id"].(string)
	roleId := claims["role"].(int)
	auth := &model.Auth{
		ID:     id,
		RoleID: roleId,
	}
	return auth, nil
}

func (t TokenUtil) getToken(ctx *fiber.Ctx) (*jwt.Token, error) {
	tokenString, err := t.getTokenFromRequest(ctx)
	if err != nil {
		return nil, err
	}
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return t.SecretKey, nil
	})
	return token, err
}

func (t TokenUtil) getTokenFromRequest(ctx *fiber.Ctx) (string, error) {
	var tokenString string
	authorization := ctx.Get("Authorization")

	if strings.HasPrefix(authorization, "Bearer ") {
		tokenString = strings.TrimPrefix(authorization, "Bearer ")
	} else if ctx.Cookies("token") != "" {
		tokenString = ctx.Cookies("token")
	}

	if tokenString == "" {
		return "", fiber.ErrUnauthorized
	}
	return tokenString, nil
}

func (t TokenUtil) ValidateJWT(context *fiber.Ctx) error {
	token, err := t.getToken(context)
	if err != nil {
		return err
	}
	_, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		return nil
	}
	return errors.New("invalid token provided")
}

func (t TokenUtil) ValidateAdminRoleJWT(context *fiber.Ctx) error {
	token, err := t.getToken(context)
	if err != nil {
		return err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	userRole := uint(claims["role"].(int))
	if ok && token.Valid && userRole == 1 {
		return nil
	}
	return errors.New("invalid admin token provided")
}

func (t TokenUtil) ValidateCustomerRoleJWT(context *fiber.Ctx) error {
	token, err := t.getToken(context)
	if err != nil {
		return err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	userRole := uint(claims["role"].(float64))
	if ok && token.Valid && userRole == 2 || userRole == 1 {
		return nil
	}

	return errors.New("invalid customer or admin token provided")
}

func (t TokenUtil) ValidateSuperAdminRoleJWT(context *fiber.Ctx) error {
	token, err := t.getToken(context)
	if err != nil {
		return err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	userRole := uint(claims["role"].(float64))
	if ok && token.Valid && userRole == 99 {
		return nil
	}

	return errors.New("invalid customer or admin token provided")
}

func (t TokenUtil) ValidateDriverRoleJWT(context *fiber.Ctx) error {
	token, err := t.getToken(context)
	if err != nil {
		return err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	userRole := uint(claims["role"].(float64))
	if ok && token.Valid && userRole == 4 {
		return nil
	}

	return errors.New("invalid customer or admin token provided")
}
