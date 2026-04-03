package demo

import (
	"context"
	"fmt"

	"github.com/example/gocli/internal/shortcut"
)

// Shortcuts 返回 demo 领域的所有 Shortcut
func Shortcuts() []shortcut.Shortcut {
	return []shortcut.Shortcut{
		{
			Service:     "demo",
			Command:     "+greet",
			Description: "Greet someone (demo shortcut)",
			Flags: []shortcut.Flag{
				{Name: "name", Desc: "Name to greet", Default: "World"},
			},
			Execute: func(ctx context.Context, rt *shortcut.Runtime) error {
				name := rt.Str("name")
				return rt.Out(map[string]interface{}{
					"message": fmt.Sprintf("Hello, %s!", name),
					"api_key": rt.Config.APIKey,
				})
			},
		},
	}
}
