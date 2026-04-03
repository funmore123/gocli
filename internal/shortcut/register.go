package shortcut

import (
	"context"

	"github.com/spf13/cobra"
)

// RegisterAll 将所有 Shortcut 挂载到 cobra 命令树上
func RegisterAll(root *cobra.Command, shortcuts []Shortcut) {
	for _, s := range shortcuts {
		parent := root
		if s.Service != "" {
			parent = findOrCreateServiceCmd(root, s.Service)
		}
		parent.AddCommand(buildCmd(s))
	}
}

// findOrCreateServiceCmd 查找已有的服务命令，没有则创建
func findOrCreateServiceCmd(root *cobra.Command, service string) *cobra.Command {
	for _, c := range root.Commands() {
		if c.Name() == service {
			return c
		}
	}
	c := &cobra.Command{
		Use:   service,
		Short: service + " service commands",
	}
	root.AddCommand(c)
	return c
}

func buildCmd(s Shortcut) *cobra.Command {
	cmd := &cobra.Command{
		Use:   s.Command,
		Short: s.Description,
		RunE: func(cmd *cobra.Command, args []string) error {
			rt, err := newRuntime(cmd)
			if err != nil {
				return err
			}
			return s.Execute(context.Background(), rt)
		},
	}
	for _, f := range s.Flags {
		cmd.Flags().String(f.Name, f.Default, f.Desc)
		if f.Required {
			_ = cmd.MarkFlagRequired(f.Name)
		}
	}
	return cmd
}
