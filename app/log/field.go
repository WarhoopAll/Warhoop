package log

import (
	"log/slog"
)

type Field = slog.Attr

func Int(key string, value int) Field {
	return slog.Int(key, value)
}

func Int32(key string, value int32) Field {
	return slog.Int(key, int(value))
}

func Int64(key string, value int64) Field {
	return slog.Int64(key, value)
}

func Uint(key string, value uint) Field {
	return slog.Uint64(key, uint64(value))
}

func Uint64(key string, value uint64) Field {
	return slog.Uint64(key, value)
}

func Float32(key string, value float32) Field {
	return slog.Float64(key, float64(value))
}

func Float64(key string, value float64) Field {
	return slog.Float64(key, value)
}

func ByteString(key string, value []byte) Field {
	return slog.String(key, string(value))
}

func String(key string, value string) Field {
	return slog.String(key, value)
}

func Any(key string, value interface{}) Field {
	return slog.Any(key, value)
}

func Object(key string, value interface{}) Field {
	return slog.Any(key, value)
}

func Array(key string, value interface{}) Field {
	return slog.Any(key, value)
}
