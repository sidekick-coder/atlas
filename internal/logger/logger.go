package logger

import "maps"

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
	transports []Transport
	metas      map[string]any
}

func Create() *Logger {
	return &Logger{
		metas: make(map[string]any),
	}
}

func (l *Logger) AddTransport(t Transport) {
	l.transports = append(l.transports, t)
}

func (l *Logger) Set(key string, value any) {
	l.metas[key] = value
}

func (l *Logger) Child() {
	c := Create()

	maps.Copy(c.metas, l.metas)
}

func (l *Logger) Log(level string, msg string, args ...any) {
	for _, transport := range l.transports {
		transport.Log(level, msg, args...)
	}
}


