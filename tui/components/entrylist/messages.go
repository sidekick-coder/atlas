package entrylist

import "github.com/sidekick-coder/atlas/internal/models"

type ChangedMsg struct {
	Entry models.Entry
	Exists bool
}
