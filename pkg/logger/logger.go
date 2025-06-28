package logger

import (
	"log/slog"
	"os"

	"github.com/tdutanton/Rogue_Game_go/internal/services/config"
)

// SetSettings - set settings to logger
func SetSettings(loggerInfo config.LoggerInfo) {
	file, err := os.OpenFile(loggerInfo.OutputFile, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0o644)
	if err != nil {
		panic("file open error")
	}

	var handler slog.Handler
	switch loggerInfo.Level {
	case "debug":
		handler = slog.NewJSONHandler(file, &slog.HandlerOptions{
			Level: slog.LevelDebug,
		})
	case "prod":
		handler = slog.NewJSONHandler(file, &slog.HandlerOptions{
			Level: slog.LevelError,
		})
	}

	logger := slog.New(handler)
	slog.SetDefault(logger)
}
