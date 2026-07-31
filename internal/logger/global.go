package logger

var logger *Logger 

func SetLogger(l *Logger) {
	logger = l
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
