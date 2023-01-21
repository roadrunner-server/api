package logger

import (
	"go.uber.org/zap"
)

// Named interface provides methods for named logger
type Named interface {
	// NamedLogger returns a named logger with the specified name
	NamedLogger(name string) *zap.Logger
}

// Log interface provides methods for logging
type Log interface {
	// Debug logs a message with debug level
	Debug(msg string, fields ...zap.Field)
	// Warn logs a message with warning level
	Warn(msg string, fields ...zap.Field)
	// Error logs a message with error level
	Error(msg string, fields ...zap.Field)
	// Info logs a message with info level
	Info(msg string, fields ...zap.Field)
	// DPanic logs a message with panic level and calls panic()
	DPanic(msg string, fields ...zap.Field)
	// Panic logs a message with panic level
	Panic(msg string, fields ...zap.Field)
	// Fatal logs a message with fatal level and calls os.Exit(1)
	Fatal(msg string, fields ...zap.Field)
}
