package log

import (
	"log/slog"
	"os"
	"sync"

	"grimoire/app/config"
)

type Logger struct {
	*slog.Logger
}

var (
	logger *Logger
	once   sync.Once
)

func Get() *Logger {
	once.Do(func() {
		// init log config
		cfg := config.Get()
		level := slog.LevelInfo

		if cfg != nil {
			switch cfg.Service.LogLevel {
			case "debug":
				level = slog.LevelDebug
			case "info":
				level = slog.LevelInfo
			case "warn", "warning":
				level = slog.LevelWarn
			case "err", "error":
				level = slog.LevelError
			case "fatal":
				level = slog.LevelError
			case "panic":
				level = slog.LevelError
			default:
				level = slog.LevelInfo
			}
		}

		opts := &slog.HandlerOptions{
			AddSource: true,
			Level:     level,
		}

		logger = &Logger{slog.New(slog.NewTextHandler(os.Stdout, opts))}
	})
	return logger
}
