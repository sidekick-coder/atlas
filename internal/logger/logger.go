package logger

import "maps"

// levels
type LoggerLevel struct {
	Info  int
	Error int
	Debug int
	Warn  int
}

var Levels = LoggerLevel{
	Info:  1,
	Error: 2,
	Debug: 3,
	Warn:  4,
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

func (l *Logger) Log(level int, msg string, args ...any) {
	for _, transport := range l.transports {
		transport.Log(level, msg, args...)
	}
}
