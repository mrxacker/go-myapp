package zap

import (
	"fmt"

	log "github.com/mrxacker/go-myapp/pkg/logger"
	"go.uber.org/zap"
)

type ZapLogger struct {
	L *zap.Logger
}

func NewLogger() (log.Logger, func(), error) {
	cfg := zap.Config{
		Level:            zap.NewAtomicLevelAt(zap.InfoLevel),
		Development:      false,
		Encoding:         "json",                            // or "console"
		OutputPaths:      []string{"stdout", "./app.log"},   // write to console + file
		ErrorOutputPaths: []string{"stderr", "./error.log"}, // errors to separate file
		EncoderConfig:    zap.NewProductionEncoderConfig(),
	}

	z, err := cfg.Build()
	if err != nil {
		return nil, nil, err
	}

	l := &ZapLogger{
		L: z,
	}

	cleanup := func() {
		_ = l.L.Sync()
	}

	return l, cleanup, nil
}

func (l *ZapLogger) Trace(msg string, fields ...log.Field) {
	if ce := l.L.Check(zap.DebugLevel, msg); ce != nil {
		ce.Write(zapifyFields(fields...)...)
	}
}

// Tracef logs at Trace log level using fmt formatter
func (l *ZapLogger) Tracef(msg string, args ...interface{}) {
	if ce := l.L.Check(zap.DebugLevel, ""); ce != nil {
		ce.Message = fmt.Sprintf(msg, args...)
		ce.Write()
	}
}

// Debug logs at Debug log level using fields
func (l *ZapLogger) Debug(msg string, fields ...log.Field) {
	if ce := l.L.Check(zap.DebugLevel, msg); ce != nil {
		ce.Write(zapifyFields(fields...)...)
	}
}

// Debugf logs at Debug log level using fmt formatter
func (l *ZapLogger) Debugf(msg string, args ...interface{}) {
	if ce := l.L.Check(zap.DebugLevel, ""); ce != nil {
		ce.Message = fmt.Sprintf(msg, args...)
		ce.Write()
	}
}

// Info logs at Info log level using fields
func (l *ZapLogger) Info(msg string, fields ...log.Field) {
	if ce := l.L.Check(zap.InfoLevel, msg); ce != nil {
		ce.Write(zapifyFields(fields...)...)
	}
}

// Infof logs at Info log level using fmt formatter
func (l *ZapLogger) Infof(msg string, args ...interface{}) {
	if ce := l.L.Check(zap.InfoLevel, ""); ce != nil {
		ce.Message = fmt.Sprintf(msg, args...)
		ce.Write()
	}
}

// Warn logs at Warn log level using fields
func (l *ZapLogger) Warn(msg string, fields ...log.Field) {
	if ce := l.L.Check(zap.WarnLevel, msg); ce != nil {
		ce.Write(zapifyFields(fields...)...)
	}
}

// Warnf logs at Warn log level using fmt formatter
func (l *ZapLogger) Warnf(msg string, args ...interface{}) {
	if ce := l.L.Check(zap.WarnLevel, ""); ce != nil {
		ce.Message = fmt.Sprintf(msg, args...)
		ce.Write()
	}
}

// Error logs at Error log level using fields
func (l *ZapLogger) Error(msg string, fields ...log.Field) {
	if ce := l.L.Check(zap.ErrorLevel, msg); ce != nil {
		ce.Write(zapifyFields(fields...)...)
	}
}

// Errorf logs at Error log level using fmt formatter
func (l *ZapLogger) Errorf(msg string, args ...interface{}) {
	if ce := l.L.Check(zap.ErrorLevel, ""); ce != nil {
		ce.Message = fmt.Sprintf(msg, args...)
		ce.Write()
	}
}

// Fatal logs at Fatal log level using fields
func (l *ZapLogger) Fatal(msg string, fields ...log.Field) {
	if ce := l.L.Check(zap.FatalLevel, msg); ce != nil {
		ce.Write(zapifyFields(fields...)...)
	}
}

// Fatalf logs at Fatal log level using fmt formatter
func (l *ZapLogger) Fatalf(msg string, args ...interface{}) {
	if ce := l.L.Check(zap.FatalLevel, ""); ce != nil {
		ce.Message = fmt.Sprintf(msg, args...)
		ce.Write()
	}
}
