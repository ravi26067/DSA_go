package helper

import (
	"dsa-go/LLD/MovieReview/database"
	"log"
	"os"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/mongo"
)

// JwtSignedDetails used to store the token details
type JwtSignedDetails struct {
	Email    string
	Name     string
	Username string
	UID      string
	UserType string
	jwt.StandardClaims
}

var userCollection *mongo.Collection = database.OpenCollection(database.Client, "user")

// SecretKey fetch the secret key from env
var SecretKey string = os.Getenv("SECRET_KEY")

// GenerateAllTokens used to generate tokens
func GenerateAllTokens(email string, name string, username string, userType string, UID string) (signedToken string, signedRefreshToken string, err error) {
	claims := &JwtSignedDetails{
		Email:    email,
		Name:     name,
		Username: username,
		UID:      UID,
		UserType: userType,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(12)).Unix(),
		},
	}

	refreshClaims := &JwtSignedDetails{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(100)).Unix(),
		},
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodES256, claims).SignedString([]byte(SecretKey))
	refreshToken, err := jwt.NewWithClaims(jwt.SigningMethodES256, refreshClaims).SignedString([]byte(SecretKey))

	if err != nil {
		log.Panic(err)
		return
	}

	return token, refreshToken, err
}
