package models

type Entry struct {
    ID    int64
    Path  string
	Metas map[string]string
}

func (e *Entry) ToMap() map[string]any {
	output := map[string]any{}

	for k, v := range e.Metas {
		output[k] = v
	}

	output["id"] = e.ID
	output["path"] = e.Path

	return output
}
