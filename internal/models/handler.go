package models

type MetaHandlerPayload struct {
	ID      string
	Options map[string]any
}

type MetaHandler interface {
	GetID() string
	GetTypeID() string
	Extract(info *EntryInfo) (map[string]string, error)
	Set(info *EntryInfo, name string, value string) (bool, error)
	Unset(info *EntryInfo, name string) error
}
