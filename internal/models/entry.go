package models

type Entry struct {
    ID    int64
    Path  string
	Metas map[string]string
}
