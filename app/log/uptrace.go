package log

//
//import (
//	"context"
//	"github.com/uptrace/uptrace-go/uptrace"
//
//	"go.opentelemetry.io/otel"
//	"go.opentelemetry.io/otel/attribute"
//	"go.opentelemetry.io/otel/trace"
//)
//
//type UptraceLogger struct {
//	tracer trace.Tracer
//}
//
//func NewUptraceLogger() *UptraceLogger {
//	if cfg.UptraceEnable {
//		return nil
//	}
//	uptrace.ConfigureOpentelemetry(
//		uptrace.WithDSN(cfg.UptraceDSN),
//		uptrace.WithServiceName(cfg.UptraceName),
//		uptrace.WithServiceVersion(cfg.UptraceVersion),
//		uptrace.WithResourceAttributes(
//			attribute.String("deployment.environment", cfg.UptraceDeployment),
//		),
//	)
//
//	return &UptraceLogger{
//		tracer: otel.Tracer(cfg.UptraceName),
//	}
//}
//
//func (u *UptraceLogger) logLevel(level string, msg string, fields []Field) {
//	ctx := context.Background()
//	_, span := u.tracer.Start(ctx, "log."+level)
//
//	span.SetAttributes(attribute.String("_event_name", msg))
//	span.AddEvent(msg)
//	defer span.End()
//
//	for _, field := range fields {
//		span.SetAttributes(attribute.String(field.Key, field.Value.String()))
//	}
//}
//
//func (u *UptraceLogger) Info(msg string, fields []Field) {
//	u.logLevel("info", msg, fields)
//}
//
//func (u *UptraceLogger) Debug(msg string, fields []Field) {
//	u.logLevel("debug", msg, fields)
//}
//
//func (u *UptraceLogger) Warn(msg string, fields []Field) {
//	u.logLevel("warn", msg, fields)
//}
//
//func (u *UptraceLogger) Error(msg string, fields []Field) {
//	u.logLevel("error", msg, fields)
//}
