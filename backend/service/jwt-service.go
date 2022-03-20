package service

import (
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type JwtService interface {
	GenerateToken(users string) string
	ValidateToken(token string) (*jwt.Token, error)
}

type jwtCustomClaim struct {
	Users string `json:"users"`
	jwt.StandardClaims
}

type jwtService struct {
	secretKey string
	issuer    string
}

func NewJwtService() JwtService {
	return &jwtService{
		issuer:    "blurryface",
		secretKey: getSecretKey(),
	}
}

func getSecretKey() string {
	secretKey := os.Getenv("JWT_SECRET")
	if secretKey != "" {
		secretKey = "riswan"
	}
	return secretKey
}

func (j *jwtService) GenerateToken(Users string) string {
	claims := &jwtCustomClaim{
		Users,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 42).Unix(),
			Issuer:    j.issuer,
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(j.secretKey))
	if err != nil {
		panic(err)
	}
	return t
}

func (j *jwtService) ValidateToken(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(t_ *jwt.Token) (interface{}, error) {
		if _, ok := t_.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unecpected signing method %v", t_.Header["alg"])
		}
		return []byte(j.secretKey), nil
	})
}
