package toast

import (
	"fmt"

	"github.com/sidekick-coder/atlas/tui/messages"
)

type Handler struct {
	ID      string
	Options map[string]string
}

func Create() *Handler {
	return &Handler{
		ID:      "toast",
		Options: map[string]string{},
	}
}

func (h *Handler) Execute(ctx map[string]any) (map[string]any, error) {
	result := make(map[string]any)

	message, ok := ctx["message"].(string)

	if !ok || message == "" {
		return result, fmt.Errorf("invalid or missing 'message' in context")
	}

	msg := messages.ToastSuccessMessage(message)

	result["tea_message"] = msg

	return result, nil
}
