package auth

import (
	"log"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
)

type Service interface {
	GenerateToken(userID int) (string, error)
}

type jwtService struct {
}

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err.Error())
	}
}

// var SECRET_KEY = os.Getenv("SECRET_KEY")

var SECRET_KEY = []byte(os.Getenv("SECRET_KEY"))

func NewService() *jwtService {
	return &jwtService{}
}

func (s *jwtService) GenerateToken(userID int) (string, error) {
	claim := jwt.MapClaims{}
	claim["user_id"] = userID

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	signedToken, err := token.SignedString(SECRET_KEY)
	if err != nil {
		return signedToken, err
	}

	return signedToken, nil
}
