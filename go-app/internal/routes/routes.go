package routes

import (
	"fmt"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/kowaslki2019/mailmanager/internal/handlers"
	"github.com/kowaslki2019/mailmanager/internal/middleware"
)

func SetupRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		// Format: [GIN] TimeStamp | StatusCode | Latency | ClientIP | Method Path
		return fmt.Sprintf("[GIN] %s | %3d | %13v | %15s | %-7s \"%s\"\n",
			param.TimeStamp.Format("2006/01/02 - 15:04:05"),
			param.StatusCode,
			param.Latency,
			param.ClientIP,
			param.Method,
			param.Path,
		)
	}))
	router.Use(gin.Recovery())

	router.Use(cors.New(cors.Config{
		AllowOrigins:  []string{"http://localhost:9091", "http://127.0.0.1:9091", "http://localhost:9090", "http://127.0.0.1:9090"},
		AllowMethods:  []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:  []string{"Origin", "Content-Length", "X-Requested-With", "Content-Type", "Authorization", "Content-Language", "Accept", "User-Agent", "Deployment-Mode"},
		ExposeHeaders: []string{"Content-Length"},
		MaxAge:        12 * time.Hour,
	}))

	api := router.Group("/api/v1")
	{
		api.GET("/health", handlers.Health)
		api.POST("/autoreply", handlers.HandleAutoRelply)
		api.POST("/rules", handlers.HandleRules)
		api.GET("/users/:username/mailboxes", handlers.GetMailboxes)
		api.DELETE("/rules/:rule_id/:username", handlers.DeleteRule)
	}

	middleware.InitRegisteredRoutesList(router)

	return router
}
