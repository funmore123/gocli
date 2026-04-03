package ark

import (
	"context"
	"fmt"

	"github.com/example/gocli/internal/shortcut"
	"github.com/volcengine/volcengine-go-sdk/volcengine/universal"
)

var listModelsShortcut = shortcut.Shortcut{
	Service:     "ark",
	Command:     "+list-models",
	Description: "List available foundation models on Volcengine ARK",
	Flags: []shortcut.Flag{
		{Name: "name", Desc: "Filter by model name (substring match)"},
	},
	Execute: func(ctx context.Context, rt *shortcut.Runtime) error {
		client, err := newUniversalClient(rt.Config)
		if err != nil {
			return fmt.Errorf("failed to create ARK client: %w", err)
		}

		input := map[string]interface{}{
			"PageNumber": 1,
			"PageSize":   100,
			"SortOrder":  "Desc",
			"SortBy":     "UpdateTime",
		}

		nameFilter := rt.Str("name")
		if nameFilter != "" {
			input["Filter"] = map[string]interface{}{
				"FoundationModelName": nameFilter,
			}
		} else {
			input["Filter"] = map[string]interface{}{}
		}

		resp, err := client.DoCall(universal.RequestUniversal{
			ServiceName: serviceName,
			Action:      "ListFoundationModels",
			Version:     apiVersion,
			HttpMethod:  universal.POST,
			ContentType: universal.ApplicationJSON,
		}, &input)
		if err != nil {
			return fmt.Errorf("ListFoundationModels failed: %w", err)
		}

		return rt.Out(resp)
	},
}
