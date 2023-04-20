package logger

import (
	"armut-notification-api/internal/util/env"
	"go.uber.org/zap"
	"runtime"
)

type ILogger interface {
	Info(message string, fields ...zap.Field)
	Error(message string, fields ...zap.Field)
	Panic(message string, fields ...zap.Field)
	With(fields ...zap.Field) ILogger
}

type Logger struct {
	environment env.IEnvironment
	logger      *zap.Logger
}

func New(environment env.IEnvironment) ILogger {
	zapLogger, err := zap.NewProduction()
	if err != nil {
		panic("Panicked while creating zap logger.")
	}

	hostname, err := environment.GetHostname()
	if err != nil {
		panic("Panicked while getting hostname.")
	}
	zapLogger = zapLogger.With(
		zap.String("os", runtime.GOOS),
		zap.String("arch", runtime.GOARCH),
		zap.String("version", runtime.Version()),
		zap.String("hostName", hostname),
		zap.String("environment", environment.Get(env.AppEnvironment)),
	)

	return &Logger{
		environment: environment,
		logger:      zapLogger,
	}
}

func (l *Logger) With(fields ...zap.Field) ILogger {
	return &Logger{
		environment: l.environment,
		logger:      l.logger.With(fields...),
	}
}
func (l *Logger) Info(message string, fields ...zap.Field) {
	l.logger.Info(message, fields...)
}

func (l *Logger) Error(message string, fields ...zap.Field) {
	l.logger.Error(message, fields...)
}

func (l *Logger) Panic(message string, fields ...zap.Field) {
	l.logger.Panic(message, fields...)
}
