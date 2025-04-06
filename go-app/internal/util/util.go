package util

import (
	"crypto/rand"
	"encoding/hex"
	"log"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

var jwtLifetime time.Duration
var jwtRefreshLifetime time.Duration

func Init() {
	if os.Getenv("GO_MAIL_MANAGER_APP_MOD") == "PROD" {
		log.Println("GO_MAIL_MANAGER_APP_MOD is PROD")
		gin.SetMode(gin.ReleaseMode)
		jwtLifetime = GetXHoursDuration(24) // 1 day
	} else if os.Getenv("GO_MAIL_MANAGER_APP_MOD") == "DEV" {
		log.Println("GO_MAIL_MANAGER_APP_MOD is DEV")
		gin.SetMode(gin.TestMode)
		jwtLifetime = GetXHoursDuration(12)
	} else {
		log.Println("GO_MAIL_MANAGER_APP_MOD is DEV_LOCAL")
		jwtLifetime = GetXMinutesDuration(10)
	}
}

func GetOSEnv() string {
	return os.Getenv("GO_MAIL_MANAGER_APP_MOD")
}

func GetJWTRefreshLifetime() time.Duration {
	return jwtRefreshLifetime
}

func GetJWTLifetime() time.Duration {
	return jwtLifetime
}

func ExtractBearerToken(token string) string {
	return strings.Split(token, " ")[1]
}
func EncodedStringGenerator(size int) string {
	bytes := make([]byte, size)
	if _, err := rand.Read(bytes); err != nil {
		log.Println(err.Error())
	}
	//fmt.Printf("%x", md5.Sum(bytes))
	return hex.EncodeToString(bytes)
}
func GenerateHexEncodedKey() string {
	return EncodedStringGenerator(20)
}
func GeneratePassword() string {
	return EncodedStringGenerator(4)
}

func GetXMilliSecondsDuration(x int) time.Duration {
	return time.Duration(time.Duration(x) * time.Millisecond)
}

func GetXSecondsDuration(x int) time.Duration {
	return time.Duration(time.Duration(x) * time.Second)
}

func GetXMinutesDuration(x int) time.Duration {
	return time.Duration(time.Duration(x) * time.Minute)
}

func GetXHoursDuration(x int) time.Duration {
	return time.Duration(time.Duration(x) * time.Hour)
}
