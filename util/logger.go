package util

import (
	"context"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger interface {
	Debug(ctx context.Context, msg string, fields map[string]interface{})
	Info(ctx context.Context, msg string, fields map[string]interface{})
	Warn(ctx context.Context, msg string, fields map[string]interface{})
	Error(ctx context.Context, msg string, err error, fields map[string]interface{})
	Fatal(ctx context.Context, msg string, err error, fields map[string]interface{})

	WithDuration(ctx context.Context, operation string, fields map[string]interface{}) func(err *error)
	Sync() error
}

type zapLogger struct {
	log *zap.Logger
}

func NewZapLogger() Logger {
	z, _ := zap.NewProduction()
	return &zapLogger{log: z}
}

func (l *zapLogger) Debug(ctx context.Context, msg string, fields map[string]interface{}) {
	l.log.Debug(msg, convertFields(fields)...)
}

func (l *zapLogger) Info(ctx context.Context, msg string, fields map[string]interface{}) {
	l.log.Info(msg, convertFields(fields)...)
}

func (l *zapLogger) Warn(ctx context.Context, msg string, fields map[string]interface{}) {
	l.log.Warn(msg, convertFields(fields)...)
}

func (l *zapLogger) Error(ctx context.Context, msg string, err error, fields map[string]interface{}) {
	fs := convertFields(fields)
	if err != nil {
		fs = append(fs, zap.Error(err))
	}
	l.log.Error(msg, fs...)
}

func (l *zapLogger) Fatal(ctx context.Context, msg string, err error, fields map[string]interface{}) {
	fs := convertFields(fields)
	if err != nil {
		fs = append(fs, zap.Error(err))
	}
	l.log.Fatal(msg, fs...)
}

func (l *zapLogger) WithDuration(ctx context.Context, operation string, fields map[string]interface{}) func(err *error) {
	start := time.Now()
	return func(err *error) {
		fields["duration"] = time.Since(start).String()
		if err != nil {
			l.Error(ctx, operation, *err, fields)
		} else {
			l.Info(ctx, operation, fields)
		}
	}
}

func (l *zapLogger) Sync() error {
	return l.log.Sync()
}

func convertFields(fields map[string]interface{}) []zapcore.Field {
	zf := make([]zapcore.Field, 0, len(fields))
	for k, v := range fields {
		zf = append(zf, zap.Any(k, v))
	}
	return zf
}
