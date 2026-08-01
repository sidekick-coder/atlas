package logger

import (
	"maps"
)

// levels
type LoggerLevel struct {
	Info  string
	Error string
	Debug string
	Warn  string
}

var Levels = LoggerLevel{
	Info:  "INFO",
	Error: "ERROR",
	Debug: "DEBUG",
	Warn:  "WARN",
}

type Logger struct {
	metas      map[string]any
}

func Create() *Logger {
	return &Logger{
		metas: make(map[string]any),
	}
}

func (l *Logger) Set(key string, value any) {
	l.metas[key] = value
}

func (l *Logger) Child(args ...any) *Logger {
	c := Create()

	for i := 0; i < len(args); i += 2 {
		if i+1 < len(args) {
			key, ok := args[i].(string)

			if !ok {
				continue
			}

			c.Set(key, args[i+1])
		}
	}

	maps.Copy(c.metas, l.metas)

	return c
}

func (l *Logger) Log(level string, msg string, args ...any) {
	for _, transport := range transports {
		transport.Log(level, msg, args...)
	}
}

func (l *Logger) Info(msg string, args ...any) {
	l.Log(Levels.Info, msg, args...)
}

func (l *Logger) Error(msg string, args ...any) {
	l.Log(Levels.Error, msg, args...)
}

func (l *Logger) Debug(msg string, args ...any) {
	l.Log(Levels.Debug, msg, args...)
}

func (l *Logger) Warn(msg string, args ...any) {
	l.Log(Levels.Warn, msg, args...)
}


