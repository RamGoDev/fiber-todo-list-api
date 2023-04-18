package helpers

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
	"todo-list/app/models"
	"todo-list/configs"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

// Compare hash and a plain text, will returns true if hash is match
//
// param	hashed	string
// param	plain	string
// return	bool
func CompareHash(hashed string, text string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(text))
	return err == nil
}

// Get JWT Signing Method for manipulate JWT Token
//
// return jwt.SigningMethod
func GetJwtSigningMethod() jwt.SigningMethod {
	return jwt.SigningMethodHS256
}

// Generate JWT Token
//
// param user models.User
// return models.User, string, error
func GenerateJwt(user *models.User) (*models.User, string, error) {
	hours, errHour := strconv.Atoi(configs.GetEnv("JWT_EXPIRED_HOUR"))
	if errHour != nil {
		return user, "", errors.New("failed to get hours of JWT_EXPIRED_HOUR")
	}

	claims := jwt.MapClaims{
		"id":    user.ID,
		"email": user.Email,
		"exp":   time.Now().Add(time.Hour * time.Duration(hours)).Unix(),
	}
	secret := []byte(configs.GetEnv("JWT_KEY"))

	token := jwt.NewWithClaims(GetJwtSigningMethod(), claims)
	tokenString, err := token.SignedString(secret)

	if err != nil {
		return user, "", errors.New("failed Generate JWT")
	}

	return user, tokenString, nil
}

// Validate JWT Token from Headers
//
// param c fiber.Ctx
// return bool, error
func ValidateJWT(c *fiber.Ctx) (bool, error) {
	authheader := c.Get("Authorization")

	if !strings.Contains(authheader, "Bearer") {
		return false, errors.New("invalid token")
	}

	authString := strings.Split(authheader, "Bearer ")
	tokenString := authString[1]

	secret := []byte(configs.GetEnv("JWT_KEY"))
	signingMethod := GetJwtSigningMethod()

	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if t.Method != signingMethod {
			return nil, fmt.Errorf("unexpected signing method")
		}
		return secret, nil
	})

	if err != nil {
		return false, errors.New("unauthorized")
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// Set Locals based on token claims
		c.Locals("user_id", claims["id"])
		c.Locals("email", claims["email"])
	}

	if !token.Valid {
		return false, errors.New("invalid token")
	}

	return true, nil
}
