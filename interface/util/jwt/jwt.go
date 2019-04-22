package jwt

import (
	"github.com/dgrijalva/jwt-go"
	"os"
	"time"
)
type jwtUtil struct {}

type IJWTUtil interface {
	CreateAccessToken(claim Claim) (string, error)
}

func NewJWTUtil() IJWTUtil {
	return &jwtUtil{}
}

func (util *jwtUtil) CreateAccessToken(claim Claim) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["iss"] = claim.Issuer
	claims["sub"] = "AccessToken"
	claims["aud"] = ""
	claims["nbf"] = ""
	claims["iat"] = time.Now()
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	claims["jti"] = ""
	resultJwt, _ := token.SignedString([]byte(os.Getenv("PRIVATE_KEY")))
	return resultJwt, nil
}