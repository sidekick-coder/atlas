package logger

import "slices"

var logger *Logger = Create()
var transports []Transport

func AddTransport(t ...Transport) {
	transports = slices.Concat(transports, t)
}

func Child(args ...any) *Logger {
	return logger.Child(args...)
}

func Info(msg string, args ...any) {
	logger.Log(Levels.Info, msg, args...)
}

func Error(msg string, args ...any) {
	logger.Log(Levels.Error, msg, args...)
}

func Debug(msg string, args ...any) {
	logger.Log(Levels.Debug, msg, args...)
}

func Warn(msg string, args ...any) {
	logger.Log(Levels.Warn, msg, args...)
}

func List(options ...ListOptions) ([]Log, error) {
	if len(transports) == 0 {
		return []Log{}, nil
	}

	t := transports[0]

	return t.List(options...)
}
