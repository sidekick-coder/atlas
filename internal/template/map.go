package template

import "github.com/sidekick-coder/atlas/internal/utils"

func EvaluateMap(payload map[string]any, context map[string]any) (map[string]any, error) {
	result := make(map[string]any)

	for key, value := range utils.FlattenMap(payload, "") {
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
