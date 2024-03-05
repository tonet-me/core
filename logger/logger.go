package logger

import (
	"io"
	"log/slog"
	"os"
	"sync"

	"gopkg.in/natefinch/lumberjack.v2"
)

var (
	sLogger *slog.Logger
	once    sync.Once
)

const (
	defaultFilePath        = "logs/logs.json"
	defaultUseLocalTime    = false
	defaultFileMaxSizeInMB = 10
	defaultFileAgeInDays   = 15
)

func GetLogger() *slog.Logger {
	once.Do(func() {
		fileWriter := &lumberjack.Logger{
			Filename:  defaultFilePath,
			LocalTime: defaultUseLocalTime,
			MaxSize:   defaultFileMaxSizeInMB,
			MaxAge:    defaultFileAgeInDays,
		}
		sLogger = slog.New(slog.NewJSONHandler(io.MultiWriter(fileWriter, os.Stdout), &slog.HandlerOptions{}))
	})
	return sLogger
}
