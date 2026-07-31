package logger

type Log struct {
	Time    string
	Level   string
	Msg     string
	Options map[string]any
}

type ListOptions struct {
	Limit  int
	Offset int
}

type Transport interface {
	Log(level string, msg string, args ...any)
	List(options ...ListOptions) ([]Log, error)
}
