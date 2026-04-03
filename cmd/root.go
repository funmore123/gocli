package cmd

import (
	"github.com/example/gocli/cmd/auth"
	"github.com/example/gocli/internal/skill"
	"github.com/example/gocli/skills"
	"github.com/spf13/cobra"
)

func Execute() int {
	rootCmd := &cobra.Command{
		Use:          "gocli",
		Short:        "A Go CLI template with Skill support",
		SilenceUsage: true,
	}

	// 固定命令
	rootCmd.AddCommand(auth.NewCmdAuth())

	// Skill 声明式命令
	skill.RegisterAll(rootCmd, skills.All())

	if err := rootCmd.Execute(); err != nil {
		return 1
	}
	return 0
}
