package demo

import (
	"context"
	"fmt"

	"github.com/example/gocli/internal/skill"
)

// Skills 返回 demo 领域的所有 Skill
func Skills() []skill.Skill {
	return []skill.Skill{
		{
			Service:     "demo",
			Command:     "+greet",
			Description: "Greet someone (demo skill)",
			Flags: []skill.Flag{
				{Name: "name", Desc: "Name to greet", Default: "World"},
			},
			Execute: func(ctx context.Context, rt *skill.Runtime) error {
				name := rt.Str("name")
				return rt.Out(map[string]interface{}{
					"message": fmt.Sprintf("Hello, %s!", name),
					"api_key": rt.Config.APIKey,
				})
			},
		},
	}
}
