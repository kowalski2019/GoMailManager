package logger

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	logger *zap.Logger
	once   sync.Once
)

func LoggerMiddleware(logger *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		duration := time.Since(start)

		// Get the caller information
		pc, file, line, ok := runtime.Caller(1)
		var functionName, fileName string
		if ok {
			functionName = runtime.FuncForPC(pc).Name()
			fileName = filepath.Base(file) // Get just the file name, not the full path
		}

		logger.Info("Request",
			zap.String("method", c.Request.Method),
			zap.String("url", c.Request.URL.String()),
			zap.Int("status", c.Writer.Status()),
			zap.Duration("duration", duration),
			zap.String("file", fileName),
			zap.String("function", functionName),
			zap.Int("line", line),
		)
	}
}

// Initialize the logger
func InitLogger(level zapcore.Level) {
	once.Do(func() {
		encoderConfig := zapcore.EncoderConfig{
			TimeKey:        "time",
			LevelKey:       "level",
			NameKey:        "logger",
			CallerKey:      "caller",
			MessageKey:     "msg",
			StacktraceKey:  "stacktrace",
			LineEnding:     zapcore.DefaultLineEnding,
			EncodeLevel:    zapcore.CapitalLevelEncoder,
			EncodeTime:     zapcore.ISO8601TimeEncoder,
			EncodeDuration: zapcore.StringDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
		}

		core := zapcore.NewCore(
			zapcore.NewConsoleEncoder(encoderConfig),
			zapcore.AddSync(os.Stdout),
			level,
		)

		logger = zap.New(core)
	})
}
// GetLogger returns the initialized logger
func GetLogger() *zap.Logger {
    if logger == nil {
        InitLogger(zapcore.InfoLevel) // Default to Info level if not initialized
    }
    return logger
}

func mlog(level zapcore.Level, message string) {
    if logger == nil {
        InitLogger(zapcore.InfoLevel) // Default to Info level if not initialized
    }

    // Get caller information
    _, file, line, ok := runtime.Caller(2) // Use 2 to skip this function and the calling function
    if !ok {
        file = "unknown"
        line = 0
    }

    // Create a formatted log message
    logMessage := fmt.Sprintf("%s:%d - %s",
        file, 
        line, 
        message,
    )

    // Log the message
    switch level {
    case zapcore.DebugLevel:
        logger.Debug(logMessage)
    case zapcore.InfoLevel:
        logger.Info(logMessage)
    case zapcore.WarnLevel:
        logger.Warn(logMessage)
    case zapcore.ErrorLevel:
        logger.Error(logMessage)
    }
}

// Debug logs a message at DebugLevel
func Debug(args ...interface{}) {
    mlog(zapcore.DebugLevel, fmt.Sprintln(args...))
}

// Info logs a message at InfoLevel
func Info(args ...interface{}) {
    mlog(zapcore.InfoLevel, fmt.Sprintln(args...))
}

// Warn logs a message at WarnLevel
func Warn(args ...interface{}) {
    mlog(zapcore.WarnLevel, fmt.Sprintln(args...))
}

// Error logs a message at ErrorLevel
func Error(args ...interface{}) {
    mlog(zapcore.ErrorLevel, fmt.Sprintln(args...))
}

// Debugf logs a formatted message at DebugLevel
func Debugf(format string, args ...interface{}) {
    Debug(fmt.Sprintf(format, args...))
}

// Infof logs a formatted message at InfoLevel
func Infof(format string, args ...interface{}) {
    Info(fmt.Sprintf(format, args...))
}

// Warnf logs a formatted message at WarnLevel
func Warnf(format string, args ...interface{}) {
    Warn(fmt.Sprintf(format, args...))
}

// Errorf logs a formatted message at ErrorLevel
func Errorf(format string, args ...interface{}) {
    Error(fmt.Sprintf(format, args...))
}