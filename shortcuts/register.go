package shortcuts

import (
	"github.com/example/gocli/internal/shortcut"
	"github.com/example/gocli/shortcuts/ark"
	"github.com/example/gocli/shortcuts/demo"
)

// All 汇总所有领域的 Shortcut
func All() []shortcut.Shortcut {
	var all []shortcut.Shortcut
	all = append(all, demo.Shortcuts()...)
	all = append(all, ark.Shortcuts()...)
	return all
}
