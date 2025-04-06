package tasks

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"time"

	"github.com/kowaslki2019/mailmanager/config"
	"github.com/kowaslki2019/mailmanager/internal/logger"
)

func CheckAutoReplies() {
	startTime := time.Now()
	logger.Debug("Checking auto replies at:", startTime)

	logPath := config.AppConfig.Log_Path_Lin
	if runtime.GOOS == "windows" {
		logPath = config.AppConfig.Log_Path_Win
	}
	// Open log file for writing
	logFile, err := os.OpenFile(filepath.Join(logPath, "CheckAutoReplies_Task.log"), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		logger.Error("Error opening log file:", err)
		return
	}
	defer logFile.Close()

	// Do auto replies check
	// auto_reply_<email>_<date>.txt

	endTime := time.Now()
	logger.Debug("Check auto replies finished at:", endTime)
	logFile.WriteString(fmt.Sprintf("%s Task completed. Duration: %v\n\n", endTime.Format("2006-01-02 15:04:05"), endTime.Sub(startTime)))
}

func CheckEmailsRules() {
	startTime := time.Now()
	logger.Debug("Checking emails rules at:", startTime)

	logPath := config.AppConfig.Log_Path_Lin
	if runtime.GOOS == "windows" {
		logPath = config.AppConfig.Log_Path_Win
	}
	// Open log file for writing
	logFile, err := os.OpenFile(filepath.Join(logPath, "CheckEmailsRules_Task.log"), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		logger.Error("Error opening log file:", err)
		return
	}
	defer logFile.Close()

	// Do emails rules check

	endTime := time.Now()
	logger.Debug("Check emails rules finished at:", endTime)
	logFile.WriteString(fmt.Sprintf("%s Task completed. Duration: %v\n\n", endTime.Format("2006-01-02 15:04:05"), endTime.Sub(startTime)))
}
