package sync

import (
	"strings"

	"github.com/sidekick-coder/atlas/internal/metadata"
	"github.com/sidekick-coder/atlas/internal/models"
)

func (s *Sync) OneByInfo(info *models.EntryInfo) error {
	smtmt := []string{}
	params := []any{}

	m, err := metadata.Create(info)

	if err != nil {
		return err
	}

	data, err := m.ExtractMap()

	if err != nil {
		return err
	}

	entry, err := s.entryRepo.Upsert(info.Path)

	if err != nil {
		return err
	}

	is, ip := s.entryMetaRepo.InsertManySql(entry.ID, data)

	smtmt = append(smtmt, is)
	params = append(params, ip...)

	finalSmtmt := strings.Join(smtmt, ";\n")

	_, err = s.Database.Exec(finalSmtmt, params...)

	println(finalSmtmt)

	if err != nil {
		return err
	}

	return nil
}

func (s *Sync) One(filepath string) error {
	info, err := s.drive.Get(filepath)

	if err != nil {
		return err
	}

	return s.OneByInfo(info)
}
