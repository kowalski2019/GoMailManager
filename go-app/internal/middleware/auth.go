package middleware

import (
	"net/http"
	"strings"

	"github.com/kowaslki2019/mailmanager/config"
	"github.com/kowaslki2019/mailmanager/internal/util"

	"github.com/gin-gonic/gin"
)

var blockList = []string{}
var registeredRoutes = []gin.RouteInfo{}

func InitRegisteredRoutesList(router *gin.Engine) {
	registeredRoutes = router.Routes()
}

func ExtractAuthKey(token string) string {
	// array[0] = Auth-Key
	return strings.Split(token, " ")[1]
}

func GetAuthKeyFromRequest(r *http.Request) string {
	return ExtractAuthKey(r.Header.Get("Authorization"))
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Header["Authorization"] != nil {
			if GetAuthKeyFromRequest(c.Request) == config.AppConfig.AuthKey {
				c.Next()
			} else {
				util.SendUnauthorizedResponse(c, "You're Unauthorized due Wrong auth key provided.")
				c.Abort()
				return
			}
		} else {
			util.SendUnauthorizedResponse(c, "You're Unauthorized due to No token in the header.")
			c.Abort()
			return
		}

	}
}

func BlockInvalidPaths() gin.HandlerFunc {
	return func(c *gin.Context) {
		clientIP := c.ClientIP()
		path := c.Request.URL.Path

		if !PathExists(path) {
			if !IsIPToBlock(clientIP) {
				blockList = append([]string{clientIP}, blockList...)
			}
			c.JSON(http.StatusForbidden, gin.H{"error": "Forbidden"})
			c.Abort()
			return
		}

		c.Next()
	}
}

func IsIPToBlock(clientIP string) bool {
	for _, ip := range blockList {
		if ip == clientIP {
			return true
		}
	}
	return false
}

func PathExists(path string) bool {
	for _, route := range registeredRoutes {
		if path == route.Path {
			return true
		}
	}

	return false
}
