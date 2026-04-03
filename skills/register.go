package skills

import (
	"github.com/example/gocli/internal/skill"
	"github.com/example/gocli/skills/demo"
)

// All 汇总所有领域的 Skill
func All() []skill.Skill {
	return demo.Skills()
}
