package template

import (
	"fmt"
	"strings"

	"github.com/sidekick-coder/atlas/internal/utils"
)

type EvaluateMapOptions struct {
	Phase     string
	Functions map[string]any
}

func NewEvaluateMapOptions() EvaluateMapOptions {
	return EvaluateMapOptions{
		Phase:     "",
		Functions: make(map[string]any),
	}
}

func EvaluateMap(payload map[string]any, context map[string]any, args ...EvaluateMapOptions) (map[string]any, error) {
	result := make(map[string]any)

	options := NewEvaluateMapOptions()

	if len(args) > 0 {
		options = args[0]
	}

	if options.Functions == nil {
		options.Functions = make(map[string]any)
	}

	flattend := utils.FlattenMap(payload, "")

	phaseKeys := make(map[string]string)

	for key, value := range flattend {
		// _evaluate is a special key that can be used to specify the phase for a specific key
		if strings.Contains(key, "_evaluate") {
			pk := strings.ReplaceAll(key, "_evaluate", "")

			// fix possible double dots in the key due to flattening
			pk = strings.ReplaceAll(pk, "..", ".")

			// remove leading dot if present
			pk = strings.TrimPrefix(pk, ".")
			pk = strings.TrimSuffix(pk, ".")

			phaseKeys[pk] = fmt.Sprintf("%v", value)
		}
	}


	for key, value := range flattend {
		keyPhase := phaseKeys[key]

		if keyPhase != "" && keyPhase != options.Phase {
			result[key] = value
			continue
		}

		vs, ok := value.(string)

		if !ok {
			result[key] = value
			continue
		}

		ev, err := Eval(vs, context)

		if err != nil {
			result[key] = value
			return nil, err
		}

		result[key] = ev
	}

	unflattened := utils.Unflatten(result)

	return unflattened, nil
}
