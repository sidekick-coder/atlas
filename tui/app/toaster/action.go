package toaster

import (
	"fmt"

	"github.com/sidekick-coder/atlas/tui/action"
	"github.com/sidekick-coder/atlas/tui/messages"
)

func HandleAction(ctx map[string]any) (map[string]any, error) {
	result := make(map[string]any)

	message, ok := ctx["message"].(string)

	if !ok || message == "" {
		return result, fmt.Errorf("invalid or missing 'message' in context")
	}

	msg := messages.ToastSuccessMessage(message)

	result["tea_message"] = msg

	return result, nil
}

func (c * Component) InitAction() {
	action.AddDefinition("toast", HandleAction)
}

