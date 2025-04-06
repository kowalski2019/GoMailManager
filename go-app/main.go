package main

import (
	"fmt"

	"github.com/kowaslki2019/mailmanager/config"
	"github.com/kowaslki2019/mailmanager/internal/logger"
	"github.com/kowaslki2019/mailmanager/internal/routes"
	"github.com/kowaslki2019/mailmanager/cmd/tasks"

	"github.com/robfig/cron/v3"

	"go.uber.org/zap/zapcore"
)

func main() {
	// Initialize configuration
	logger.InitLogger(zapcore.DebugLevel)
	config.LoadConfig()

	c := cron.New()
	c.AddFunc(config.AppConfig.Cron_CheckAutoRepliesInterval, tasks.CheckAutoReplies)
	c.AddFunc(config.AppConfig.Cron_CheckEmailsRulesInterval, tasks.CheckEmailsRules)
	//c.Start()

	// Initialize router
	router := routes.SetupRouter()
	logger.Debug("Routes successfully set")

	// Start the server
	host := fmt.Sprintf("127.0.0.1:%d", config.AppConfig.HTTPPort)
	router.Run(host)

}
