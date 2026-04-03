package cmd

import (
	"github.com/example/gocli/cmd/auth"
	"github.com/example/gocli/internal/shortcut"
	"github.com/example/gocli/shortcuts"
	"github.com/spf13/cobra"
)

func Execute() int {
	rootCmd := &cobra.Command{
		Use:          "gocli",
		Short:        "A Go CLI template with Shortcut + Skill support",
		SilenceUsage: true,
	}

	// 固定命令
	rootCmd.AddCommand(auth.NewCmdAuth())

	// Shortcut 声明式命令
	shortcut.RegisterAll(rootCmd, shortcuts.All())

	if err := rootCmd.Execute(); err != nil {
		return 1
	}
	return 0
}
