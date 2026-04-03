package ark

import "github.com/example/gocli/internal/shortcut"

// Shortcuts 返回 ARK 领域的所有 Shortcut
func Shortcuts() []shortcut.Shortcut {
	return []shortcut.Shortcut{
		listModelsShortcut,
	}
}
