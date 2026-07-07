package json

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/sidekick-coder/atlas/internal/models"
	"github.com/sidekick-coder/atlas/internal/utils"
)

type Handler struct {}

func UnmarshalFromBytes(content []byte) (map[string]any, error) {
	result := map[string]any{}

	err := json.Unmarshal(content, &result)


	if err != nil {
		return nil, err
	}

	flat := utils.FlattenMap(result, "")

	return flat, nil
}

func MarshalToBytes(metas map[string]any) ([]byte, error) {
	data, err := json.MarshalIndent(metas, "", "  ")

	if err != nil {
		return nil, err
	}

	return data, nil
}

func (m Handler) ID() string {
	return "json"
}

func (m Handler) Extract(info *models.EntryInfo) (map[string]string, error) {
	contents, err := os.ReadFile(filepath.Join(info.AbsolutePath))

	if err != nil {
		return nil, err
	}

	flat, err := UnmarshalFromBytes(contents)

	if err != nil {
		return nil, err
	}

	result := utils.StringifyMap(flat)

	return result, nil
}

func (m Handler) Set(info *models.EntryInfo, name string, value string) (bool, error) {
	contents, err := os.ReadFile(filepath.Join(info.AbsolutePath))

	if err != nil {
		return false, err
	}

	metas, err := UnmarshalFromBytes(contents)

	if err != nil {
		return false, err
	}

	metas[name] = value

	data := utils.Unflatten(metas)

	newContents, err := json.MarshalIndent(data, "", "  ")

	if err != nil {
		return false, err
	}

	err = os.WriteFile(filepath.Join(info.AbsolutePath), []byte(newContents), 0644)

	if err != nil {
		return false, err
	}

	return true, nil
}

func (m Handler) Unset(info *models.EntryInfo, name string) error {
	contents, err := os.ReadFile(filepath.Join(info.AbsolutePath))

	if err != nil {
		return err
	}

	metas, err := UnmarshalFromBytes(contents)

	if err != nil {
		return err
	}

	delete(metas, name)


	newContents, err := MarshalToBytes(metas)

	if err != nil {
		return err
	}

	err = os.WriteFile(filepath.Join(info.AbsolutePath), []byte(newContents), 0644)

	if err != nil {
		return err
	}

	return nil
}
