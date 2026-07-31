package logger

import (
	"fmt"
	"log/slog"
	"os"
)

type FileTransport struct {
	logger *slog.Logger
}

func CreateFileTransport(filename string) (*FileTransport, error) {
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)

	if err != nil {
		return nil, fmt.Errorf("failed to open log file: %w", err)
	}

	logger := slog.New(slog.NewJSONHandler(file, nil))

	transport := &FileTransport{
		logger: logger,
	}

	return transport, nil
}

func (s *FileTransport) Log(level int, msg string, args ...any) {

	if level == Levels.Info {
		s.logger.Info(msg, args...)
		return
	}

	if level == Levels.Error {
		s.logger.Error(msg, args...)
		return
	}

	if level == Levels.Debug {
		s.logger.Debug(msg, args...)
		return
	}

	if level == Levels.Warn {
		s.logger.Warn(msg, args...)
		return
	}
}

func (s *FileTransport) GetLogger() *slog.Logger {
	return s.logger
}
