package helpers

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
	"github.com/theerudito/istla/model/dto"
)

type CustomClaims struct {
	UserId int    `json:"user_id"`
	Name   string `json:"name"`
	jwt.RegisteredClaims
}

func getSecretKey() []byte {

	_ = godotenv.Load()

	key := os.Getenv("Secret_Key")
	if key == "" {
		fmt.Println("⚠️  WARNING: SECRET_KEY no está definido en las variables de entorno.")
	}
	return []byte(key)
}

var accessTokenDuration = 10 * time.Hour

func GenerateToken(user dto.UsuarJWT) (string, error) {

	var jwtSecret = []byte(os.Getenv("Secret_Key"))

	claims := CustomClaims{
		UserId: user.UsuarioId,
		Name:   user.Nombres,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(accessTokenDuration)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    os.Getenv("URL"),
			Audience:  jwt.ClaimStrings{os.Getenv("URL")},
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ParseAndVerifyToken(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("algoritmo de firma inválido")
		}
		return getSecretKey(), nil
	})

	if err != nil {
		return nil, fmt.Errorf("token inválido: %w", err)
	}

	claims, ok := token.Claims.(*CustomClaims)
	if !ok || !token.Valid {
		return nil, fmt.Errorf("claims inválidos o token inválido")
	}

	return claims, nil
}
