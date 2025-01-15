package log

import (
	"log/slog"
	"os"
	"sync"

	"warhoop/app/config"
)

type Logger struct {
	*slog.Logger
	Uptrace *UptraceLogger
}

var (
	logger *Logger
	once   sync.Once
)

var cfg = config.Get()

// Get initializes logger with log level from config. Once.
func Get() *Logger {
	once.Do(func() {
		// Get config
		level := slog.LevelInfo

		if cfg != nil {
			switch cfg.Service.LogLevel {
			case "debug":
				level = slog.LevelDebug
			case "info":
				level = slog.LevelInfo
			case "warn":
				level = slog.LevelWarn
			case "err":
				level = slog.LevelError
			case "fatal":
				level = slog.LevelError
			case "panic":
				level = slog.LevelError
			default:
				level = slog.LevelInfo
			}
		}

		consoleOpts := &slog.HandlerOptions{
			AddSource: true,
			Level:     level,
		}
		consoleHandler := slog.NewTextHandler(os.Stdout, consoleOpts)

		var uptraceLogger *UptraceLogger
		if cfg.Uptrace.Enable {
			uptraceLogger = NewUptraceLogger()
		}

		logger = &Logger{
			Logger:  slog.New(consoleHandler),
			Uptrace: uptraceLogger,
		}
	})

	return logger
}

func (l *Logger) Debug(msg string, fields ...Field) {
	l.Logger.Debug(msg, fieldsToAny(fields)...)
	if l.Uptrace != nil {
		l.Uptrace.Debug(msg, fields)
	}
}

func (l *Logger) Info(msg string, fields ...Field) {
	l.Logger.Info(msg, fieldsToAny(fields)...)
	if l.Uptrace != nil {
		l.Uptrace.Info(msg, fields)
	}
}

func (l *Logger) Warn(msg string, fields ...Field) {
	l.Logger.Warn(msg, fieldsToAny(fields)...)
	if l.Uptrace != nil {
		l.Uptrace.Warn(msg, fields)
	}
}

func (l *Logger) Error(msg string, fields ...Field) {
	l.Logger.Error(msg, fieldsToAny(fields)...)
	if l.Uptrace != nil {
		l.Uptrace.Error(msg, fields)
	}
}

func fieldsToAny(fields []Field) []any {
	result := make([]any, len(fields))
	for i, field := range fields {
		result[i] = field
	}
	return result
}
