package logger

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log/slog"
	"os"

	"github.com/sidekick-coder/atlas/internal/utils/maputil"
)

type FileTransport struct {
	filename string
	logger   *slog.Logger
}

func CreateFileTransport(filename string) (*FileTransport, error) {
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)

	if err != nil {
		return nil, fmt.Errorf("failed to open log file: %w", err)
	}

	opts := &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}

	logger := slog.New(slog.NewJSONHandler(file, opts))

	transport := &FileTransport{
		filename: filename,
		logger:   logger,
	}

	return transport, nil
}

func (s *FileTransport) Log(level string, msg string, args ...any) {

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

func (s *FileTransport) List(options ...ListOptions) ([]Log, error) {
	file, err := os.Open(s.filename)

	if err != nil {
		return nil, fmt.Errorf("failed to open log file for reading: %w", err)
	}

	defer file.Close()

	var logs []Log

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Bytes()

		data := map[string]any{}

		err := json.Unmarshal(line, &data)

		if err != nil {
			return nil, fmt.Errorf("failed to unmarshal log entry: %w", err)
		}

		log := Log{
			Time:  fmt.Sprintf("%v", data["time"]),
			Level: fmt.Sprintf("%v", data["level"]),
			Msg:   fmt.Sprintf("%v", data["msg"]),
			Options:  maputil.Except(data, "time", "level", "msg"),
		}

		logs = append(logs, log)
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading log file: %w", err)
	}

	return logs, nil

}
