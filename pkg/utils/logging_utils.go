package utils

import (
	"context"
	"log/slog"
	"os"
)

type LoggerType struct {
	*slog.Logger
}

var Logger *LoggerType

const (
	LevelFatal = slog.Level(12)
)

// Initialize logger
func InitLogger(logLevel string) {
	var level slog.Level
	switch logLevel {
	case "DEBUG":
		level = slog.LevelDebug
	case "INFO":
		level = slog.LevelInfo
	case "WARN":
		level = slog.LevelWarn
	case "ERROR":
		level = slog.LevelError
	default:
		level = slog.LevelInfo
	}

	var LevelNames = map[slog.Leveler]string{
		LevelFatal: "FATAL",
	}

	opts := &slog.HandlerOptions{
		Level: level,
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			// Replace error values with slog.Any(err)
			switch v := a.Value.Any().(type) {
			case error:
				a = slog.Any("error", v)
			}

			// Rename custom log levels
			if a.Key == slog.LevelKey {
				level := a.Value.Any().(slog.Level)
				levelLabel, exists := LevelNames[level]
				if !exists {
					levelLabel = level.String()
				}
				a.Value = slog.StringValue(levelLabel)
			}

			return a
		},
	}

	handler := slog.NewTextHandler(os.Stderr, opts)
	Logger = &LoggerType{slog.New(handler)}
}

func (l *LoggerType) Info(msg string, args ...any) {
	l.Log(context.Background(), slog.LevelInfo, msg, args...)
}

func (l *LoggerType) Debug(msg string, args ...any) {
	l.Log(context.Background(), slog.LevelDebug, msg, args...)
}

func (l *LoggerType) Error(msg string, args ...any) {
	l.Log(context.Background(), slog.LevelError, msg, args...)
}

func (l *LoggerType) Warn(msg string, args ...any) {
	l.Log(context.Background(), slog.LevelWarn, msg, args...)
}

func (l *LoggerType) Fatal(msg string, args ...any) {
	l.Log(context.Background(), LevelFatal, msg, args...)
	os.Exit(1)
}
