package metadata

import (
	"strings"
	"github.com/sidekick-coder/atlas/internal/drive/v2"
)

func Extract(info *drive.EntryInfo, handlers []Handler) (map[string]string, error) {
	metadata := make(map[string]string)

	ids := []string{}

	for _, handler := range handlers {
		data, err := handler.Extract(info)

		ids = append(ids, handler.ID())

		if err != nil {
			return nil, err
		}

		for key, value := range data {
			metadata[key] = value
		}
	}

	metadata["basename"] = info.BaseName
	metadata["type"] = info.Type
	metadata["ext"] = info.Ext
	metadata["handlers"] = strings.Join(ids, ",")

	return metadata, nil
}
