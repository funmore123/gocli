package auth

import (
	"fmt"

	"github.com/example/gocli/internal/config"
	"github.com/example/gocli/internal/output"
	"github.com/spf13/cobra"
)

func NewCmdAuth() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "auth",
		Short: "Authentication management",
	}
	cmd.AddCommand(newLoginCmd())
	cmd.AddCommand(newStatusCmd())
	cmd.AddCommand(newLogoutCmd())
	return cmd
}

func newLoginCmd() *cobra.Command {
	var apiKey, apiSecret string
	cmd := &cobra.Command{
		Use:   "login",
		Short: "Login with API key and secret",
		RunE: func(cmd *cobra.Command, args []string) error {
			cfg := &config.Config{APIKey: apiKey, APISecret: apiSecret}
			if err := config.Save(cfg); err != nil {
				return err
			}
			fmt.Printf("Login successful. Config saved to %s\n", config.Path())
			return nil
		},
	}
	cmd.Flags().StringVar(&apiKey, "api-key", "", "API key (required)")
	cmd.Flags().StringVar(&apiSecret, "api-secret", "", "API secret (required)")
	_ = cmd.MarkFlagRequired("api-key")
	_ = cmd.MarkFlagRequired("api-secret")
	return cmd
}

func newStatusCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "status",
		Short: "Show current auth status",
		RunE: func(cmd *cobra.Command, args []string) error {
			cfg, err := config.Load()
			if err != nil {
				return err
			}
			return output.JSON(map[string]interface{}{
				"api_key":     cfg.APIKey,
				"api_secret":  mask(cfg.APISecret),
				"config_path": config.Path(),
			})
		},
	}
}

func newLogoutCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "logout",
		Short: "Remove stored credentials",
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := config.Remove(); err != nil {
				return err
			}
			fmt.Println("Logged out.")
			return nil
		},
	}
}

func mask(s string) string {
	if len(s) <= 4 {
		return "****"
	}
	return s[:4] + "****"
}
