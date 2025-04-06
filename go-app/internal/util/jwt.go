package util

import (
	"github.com/kowaslki2019/mailmanager/config"
	"log"

	"github.com/golang-jwt/jwt/v4"

	"time"
)


// generate secret with nodejs require('crypto').randomBytes(64).toString('hex')

var jwt_key = config.AppConfig.JWTSecret
var secretKey = []byte(jwt_key)

func RetrieveJWTClaims(token *jwt.Token) (jwt.MapClaims, bool) {
	claims, ok := token.Claims.(jwt.MapClaims)
	return claims, ok
}

func GetPermissionFromClaims(claims jwt.MapClaims) int16 {
	var perms int16 = 1
	return perms
}

func GenerateJWT(user string, perms int16, user_company string, refreshToken bool) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["iat"] = time.Now().UTC().Unix()
	if refreshToken {
		claims["exp"] = time.Now().Add(GetJWTRefreshLifetime()).Unix()
	} else {
		claims["exp"] = time.Now().Add(GetJWTLifetime()).Unix()
	}
	claims["iss"] = "dacite/go-mail-manager/auth"

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		log.Println("Error by SignedString: ", err.Error())
		return "", err
	}

	return tokenString, nil
}