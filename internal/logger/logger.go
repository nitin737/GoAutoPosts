package logger

import (
	"log/slog"
	"os"
)

// Logger wraps slog for structured logging
type Logger struct {
	*slog.Logger
}

// NewLogger creates a new logger instance
func NewLogger() *Logger {
	handler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	})

	return &Logger{
		Logger: slog.New(handler),
	}
}

// NewDevelopmentLogger creates a logger for development with text output
func NewDevelopmentLogger() *Logger {
	handler := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	})

	return &Logger{
		Logger: slog.New(handler),
	}
}
