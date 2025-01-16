package log

import (
	"context"
	"github.com/uptrace/uptrace-go/uptrace"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

type UptraceLogger struct {
	tracer trace.Tracer
}

func NewUptraceLogger() *UptraceLogger {
	if cfg.Uptrace.Enable {
		return nil
	}
	uptrace.ConfigureOpentelemetry(
		uptrace.WithDSN(cfg.Uptrace.DSN),
		uptrace.WithServiceName(cfg.Uptrace.Name),
		uptrace.WithServiceVersion(cfg.Uptrace.Version),
		uptrace.WithResourceAttributes(
			attribute.String("deployment.environment", cfg.Uptrace.Deployment),
		),
	)

	return &UptraceLogger{
		tracer: otel.Tracer(cfg.Uptrace.Name),
	}
}

func (u *UptraceLogger) logLevel(level string, msg string, fields []Field) {
	ctx := context.Background()
	_, span := u.tracer.Start(ctx, "log."+level)

	span.SetAttributes(attribute.String("_event_name", msg))
	span.AddEvent(msg)
	defer span.End()

	for _, field := range fields {
		span.SetAttributes(attribute.String(field.Key, field.Value.String()))
	}
}

func (u *UptraceLogger) Info(msg string, fields []Field) {
	u.logLevel("info", msg, fields)
}

func (u *UptraceLogger) Debug(msg string, fields []Field) {
	u.logLevel("debug", msg, fields)
}

func (u *UptraceLogger) Warn(msg string, fields []Field) {
	u.logLevel("warn", msg, fields)
}

func (u *UptraceLogger) Error(msg string, fields []Field) {
	u.logLevel("error", msg, fields)
}
