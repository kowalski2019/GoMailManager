package config

import (
	"log"
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	JWTSecret string
	HTTPPort  int
	AuthKey   string
	Cron_CheckAutoRepliesInterval string
	Cron_CheckEmailsRulesInterval string
	Log_Path string
	Log_Path_Lin string
	Log_Path_Win string
}

var AppConfig Config

func LoadConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(GetConfigFilePath())

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	AppConfig = Config{
		JWTSecret: viper.GetString("jwt.secret"),
		HTTPPort:  viper.GetInt("http.port"),
		AuthKey:   viper.GetString("auth.key"),
		Cron_CheckAutoRepliesInterval: viper.GetString("cron.check_auto_replies_interval"),
		Cron_CheckEmailsRulesInterval: viper.GetString("cron.check_emails_rules_interval"),
		Log_Path: viper.GetString("log.path"),
		Log_Path_Lin: viper.GetString("log.path_lin"),
		Log_Path_Win: viper.GetString("log.path_win"),
	}
}

func GetConfigFilePath() string {
	if os.Getenv("GO_MAIL_MANAGER_APP_MOD") == "PROD" || os.Getenv("GO_MAIL_MANAGER_APP_MOD") == "DEV" {
		return "/opt/mailmanager/config"
	}
	return "."
}
