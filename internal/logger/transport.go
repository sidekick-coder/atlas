package logger

type Transport interface {
	Log(level int, msg string, args ...any)
}

